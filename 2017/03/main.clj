(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def square-num (parse-long (str/trim (slurp "input.txt"))))

(defn find-ring-diameter [target diameter]
  (if (and (< (* diameter diameter) target) (>= (* (+ diameter 2) (+ diameter 2)) target))
    (+ diameter 2)
    (find-ring-diameter target (+ diameter 2))))

(defn find-nearest-corner [target diameter corner]
  (if (and (>= corner target) (< (- corner (- diameter 1)) target))
    corner
    (find-nearest-corner target diameter (- corner (- diameter 1)))))

(defn find-midpoint-distance [target diameter corner]
  (def midpoint (- corner (/ (- diameter 1) 2)))
  (abs (- midpoint target)))

(defn find-center-distance [target]
  (def diameter (find-ring-diameter target 1))
  (def corner (find-nearest-corner target diameter (* diameter diameter)))
  (def midpoint-distance (find-midpoint-distance target diameter corner))
  (+ midpoint-distance (/ (- diameter 1) 2)))

(def num-steps (find-center-distance square-num))

(defn get-square [x y squares]
  (or ((or (squares x) {}) y) 0))

(defn flatten-squares [squares]
  (sort (flatten (map vals (vals squares)))))

(defn get-adjacent-coordinates [x y]
  (map #(list (+ (first %) x) (+ (second %) y))
       '((-1 -1) (-1 0) (-1 1) (0 -1) (0 1) (1 -1) (1 0) (1 1))))

(defn get-adjacent-squares [x y squares]
  (map #(get-square (first %) (second %) squares) (get-adjacent-coordinates x y)))

(defn add-square [x y squares]
  (merge squares { x (merge (squares x) { y (reduce + (get-adjacent-squares x y squares)) }) }))

(defn get-ring [diameter]
  (concat
    (map #(list diameter %) (range (+ 1 (- 0 diameter)) (+ diameter 1)))
    (map #(list % diameter) (range (- diameter 1) (- 0 (+ diameter 1)) -1))
    (map #(list (- 0 diameter) %) (range (- diameter 1) (- 0 (+ diameter 1)) -1))
    (map #(list % (- 0 diameter)) (range (- 0 (- diameter 1)) (+ diameter 1)))))

(defn find-squares [diameter]
  (if (= diameter 0)
    {0 { 0 1 }}
    (reduce #(add-square (first %2) (second %2) %1) (find-squares (- diameter 1)) (get-ring diameter))))

(defn first-exceeds-target [target diameter]
  (def squares (flatten-squares (find-squares diameter)))
  (if (some #(> % target) squares)
    (first (filter #(> % target) squares))
    (first-exceeds-target target (+ diameter 1))))

(def first-exceeds (first-exceeds-target square-num 1))

(println (format "The number of steps required to reach the access port is %d." num-steps))
(println (format "The first value that is greater than the input is %d." first-exceeds))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

