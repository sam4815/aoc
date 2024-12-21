(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def maze (mapv (comp vec seq) (str/split (slurp "input.txt") #"\n")))

(defn find-symbol [sym maze]
  (first (filter some? (for [x (range (count (first maze))) y (range (count maze))]
                         (when (= sym (get-in maze [x y])) [x y])))))

(defn get-manhattan [[x y] [i j]]
  (+ (abs (- x i)) (abs (- y j))))

(defn find-empty-moves [[x y] cost]
  (->> [[0 1] [1 0] [-1 0] [0 -1]]
       (map (fn [[i j]] [(+ i x) (+ j y)]))
       (filter (fn [pos] (not= \# (get-in maze pos))))
       (map (fn [pos] {:pos pos :cost (inc cost)}))))

(defn find-within-cost [[x y] cost]
  (->> (for [i (range (- cost) (inc cost)) j (range (- cost) (inc cost))] [(+ i x) (+ j y)])
       (filter (fn [pos] (<= (get-manhattan pos [x y]) cost)))))

(defn traverse [maze start]
  (loop [queue [{:pos start :cost 0}] visited {}]
    (let [{:keys [pos cost]} (first queue)]
      (if (empty? queue) visited
        (if (get visited pos)
          (recur (rest queue) visited)
          (recur (concat (rest queue) (find-empty-moves pos cost))
                 (assoc visited pos cost)))))))

(defn find-cheats [max-cheat-cost steps]
  (->> (for [start (keys steps) end (find-within-cost start max-cheat-cost)] [start end])
       (map (fn [[start end]] [(- (get steps end 0) (get steps start 0)) (get-manhattan start end)]))
       (filter (fn [[saving cost]] (>= (- saving cost) 100)))))

(defn count-cheats [maze max-cheat-cost]
  (->> (find-symbol \S maze)
       (traverse maze)
       (find-cheats max-cheat-cost)
       count))

(println (format "There are %d 2-move cheats that save 100 picoseconds." (count-cheats maze 2)))
(println (format "There are %d 20-move cheats that save 100 picoseconds." (count-cheats maze 20)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

