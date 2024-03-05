(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def grid (mapv #(vec %) (str/split (slurp "input.txt") #"\n")))

(defn find-numbers [grid]
  (->> (map-indexed (fn [y row] (map-indexed (fn [x cell] (if (int? (parse-long (str cell))) [cell [x y]] nil)) row)) grid)
       (apply concat)
       (filter some?)
       (into {})))

(defn find-moves [[x y] steps grid]
  (->> [[(inc x) y] [(dec x) y] [x (inc y)] [x (dec y)]]
       (filter (fn [[x y]] (not= \# (nth (nth grid y) x))))
       (map (fn [position] {:position position :steps (inc steps)}))))

(defn find-all-steps [start grid]
  (loop [queue [{:position start :steps 0}] visited {}]
    (if (empty? queue) visited
      (let [{:keys [position steps]} (first queue)]
        (if (get visited position)
          (recur (rest queue) visited)            
          (recur (sort-by :steps (concat (find-moves position steps grid) (rest queue)))
                 (assoc visited position steps)))))))

(defn find-paths [position visited steps paths]
  (->> (filter (fn [[number]] (not (some #{number} visited))) (get paths position))
       (map (fn [[number distance]] {:position number
                                     :visited (vec (concat visited [number]))
                                     :steps (+ steps distance)}))))

(defn find-shortest-path [distances go-back-to-zero]
  (loop [queue [{:position \0 :visited [\0] :steps 0}] min-steps 100000]
    (if (empty? queue) min-steps
      (let [{:keys [position visited steps]} (first queue) next-paths (find-paths position visited steps distances)]
        (if (empty? next-paths)
          (let [path-length (if go-back-to-zero (+ steps (get (get distances position) \0)) steps)]
            (recur (rest queue) (if (< path-length min-steps) path-length min-steps)))
          (recur (concat next-paths (rest queue)) min-steps))))))

(def distances (let [numbers (find-numbers grid)]
                 (into {} (map (fn [[number position]]
                                 (let [visited (find-all-steps position grid)]
                                   [number (into {} (map (fn [[other-number other-position]]
                                                           [other-number (get visited other-position)])) numbers)])) numbers))))

(def shortest-path (find-shortest-path distances false))
(def shortest-zero-path (find-shortest-path distances true))

(println (format "The shortest path that goes through all numbers takes %d steps." shortest-path))
(println (format "The shortest path that goes through all numbers and then back to zero takes %d steps." shortest-zero-path))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

