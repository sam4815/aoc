(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def configuration (->> (str/split (slurp "input.txt") #"\n")
                        (mapv str)
                        vec))

(defn find-neighbours [[x y] grid]
  (for [dx (range -1 2)
        dy (range -1 2)
        :when (or (not (zero? dx)) (not (zero? dy)))]
    (nth (nth grid (+ y dy) '()) (+ x dx) ".")))

(defn hardwire [lights]
  (reduce (fn [lights light] (assoc-in lights light \#))
          lights
          [[0 0] [0 99] [99 0] [99 99]]))

(defn next-cell [grid y x cell]
  (case [cell (count (filter (partial = \#) (find-neighbours [x y] grid)))]
    ([\# 2] [\# 3] [\. 3]) \# \.))

(defn next-grid [grid _]
  (vec (map-indexed (fn [y row] (vec (map-indexed (partial next-cell grid y) row))) grid)))

(defn count-lit [grid]
  (count (filter (partial = \#) (str/join "" grid))))

(def animated (reduce next-grid configuration (range 100)))
(def faulty-animated (reduce (comp hardwire next-grid) configuration (range 100)))

(println (format "After 100 steps, %d lights are on." (count-lit animated)))
(println (format "With the faulty lights, %d lights are on." (count-lit faulty-animated)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

