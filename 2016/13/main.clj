(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def favourite-number (parse-long (str/trim (slurp "input.txt"))))

(defn find-structure [x y]
  (let [quadratic (+ (* x x) (* 3 x) (* 2 x y) y (* y y))
        binary (Integer/toString (+ quadratic favourite-number) 2)]
    (if (= 0 (mod (get (frequencies binary) \1) 2)) \. \#)))

(defn find-adjacent-paths [grid [x y] steps]
  (->> `((~(dec x) ~y) (~(inc x) ~y) (~x ~(dec y)) (~x ~(inc y)))
       (filter (fn [[x y]] (= \. (get (get grid y []) x \#))))
       (map (fn [position] {:steps (inc steps) :position position}))))

(defn mark-steps [init-grid]
  (loop [queue [{:steps 0 :position '(1 1)}] grid init-grid]
    (if (empty? queue) grid
      (let [{:keys [steps position]} (first queue) [x y] position
            marked-grid (update grid y (fn [row] (update row x (fn [cell] steps))))]
        (recur (sort-by :steps (concat (find-adjacent-paths marked-grid position steps) (rest queue))) marked-grid)))))

(def grid (mapv (fn [y] (mapv (fn [x] (find-structure x y)) (range 100))) (range 100)))
(def marked-grid (mark-steps grid))

(def min-steps (nth (nth marked-grid 39) 31))
(def num-positions-under-50 (reduce + (map #(if (and (int? %) (<= % 50)) 1 0) (flatten marked-grid))))

(println (format "The fewest number of steps required to reach 31,39 is %d." min-steps))
(println (format "There are %d locations that can be reached within 50 steps." num-positions-under-50))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

