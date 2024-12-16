(ns main
  (:require [clojure.string :as str]
            [clojure.set :as set]))

(def start-time (System/currentTimeMillis))

(def maze (mapv (comp vec seq) (str/split (slurp "input.txt") #"\n")))

(defn find-symbol [sym guard-map]
  (vec (filter some? (for [x (range (count (first guard-map))) y (range (count guard-map))]
                       (when (= sym (get-in guard-map [x y])) [x y])))))

(defn rotate-cw [direction]
  (case direction \^ \> \> \v \v \< \< \^))

(defn rotate-ccw [direction]
  (case direction \^ \< \< \v \v \> \> \^))

(defn find-next-position [[x y dir]]
  (case dir \^ [(dec x) y dir] \v [(inc x) y dir] \> [x (inc y) dir] \< [x (dec y) dir]))

(defn save-path [visited path]
  (let [{:keys [score tiles]} (get visited (last path))]
    (assoc visited (last path) {:score score :tiles (concat path tiles)})))

(defn find-ends [visited maze]
  (let [ends (map (fn [dir] (conj (first (find-symbol \E maze)) dir)) [\^ \> \v \<])
        end-scores (map (fn [end] (get (get visited end {}) :score Integer/MAX_VALUE)) ends)]
    (filter (fn [end] (= (get (get visited end {}) :score) (apply min end-scores))) ends)))

(defn find-distinct [ends visited]
  (let [end-paths (apply concat (map (fn [end] (get (get visited end) :tiles)) ends))
        all-paths (apply concat (map (fn [tile] (get (get visited tile) :tiles)) end-paths))
        all-all-paths (apply concat (map (fn [tile] (get (get visited tile) :tiles)) (distinct all-paths)))
        all-tiles (map (fn [[x y _]] [x y]) (distinct all-all-paths))]
    (count (distinct all-tiles))))

(defn find-moves [path score maze]
  (let [[x y dir] (last path)
        move-forward (find-next-position (last path))
        rotations [{:path (conj path [x y (rotate-cw dir)]) :score (+ score 1000)}
                   {:path (conj path [x y (rotate-ccw dir)]) :score (+ score 1000)}]]
    (if (= (get-in maze (take 2 move-forward)) \#) rotations
      (conj rotations {:path (conj path move-forward) :score (+ score 1)}))))

(defn traverse [maze]
  (let [end (first (find-symbol \E maze))]
    (loop [queue [{:path [(conj (first (find-symbol \S maze)) \>)] :score 0}] visited {}]
      (if (empty? queue) visited
        (let [{:keys [path score]} (first queue) [x y _] (last path)]
            (if (<= (get (get visited (last path) {}) :score Integer/MAX_VALUE) score)
              (recur (rest queue)
                     (if (= (get (get visited (last path)) :score) score) (save-path visited path) visited))
              (recur (sort-by :score (concat (find-moves path score maze) (rest queue)))
                     (assoc visited (last path) {:score score :tiles path}))))))))

(let [visited (traverse maze)
      ends (find-ends visited maze)]
  (println (format "The lowest possible score is %d." (get (get visited (first ends)) :score)))
  (println (format "%d tiles are part of at least one of the best paths." (find-distinct ends visited))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

