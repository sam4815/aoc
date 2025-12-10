(ns main
  (:require [clojure.math :as math]))

(def start-time (System/currentTimeMillis))

(def coordinates (partition 2 (map parse-long (re-seq #"\d+" (slurp "input.txt")))))

(defn get-lines [coordinates]
  (first (reduce (fn [[all [x1 y1]] [x2 y2]]
                   [(conj all [(vec (sort [x1 x2])) (vec (sort [y1 y2]))]) [x2 y2]])
                 [[] (vec (last coordinates))]
                 coordinates)))

(defn find-distance [[x1 y1] [x2 y2]]
  (* (inc (abs (- x1 x2)))
     (inc (abs (- y1 y2)))))

(defn find-combos [coordinates]
  (->> (for [[i a] (map-indexed vector coordinates) b (subvec (vec coordinates) (inc i))]
         [(find-distance a b) a b])
       (sort-by first >)))

(defn forms-contiguous-area [lines [size [x1 y1] [x2 y2]]]
  (let [min-x (min x1 x2) max-x (max x1 x2) min-y (min y1 y2) max-y (max y1 y2)]
    (and (empty? (filter (fn [[x3 y3]] (and (< x3 max-x) (> x3 min-x) (< y3 max-y) (> y3 min-y))) coordinates))
         (empty? (filter (fn [[[x3 x4] [y3 y4]]] (if (= x3 x4)
                                                   (or
                                                     (and (> x3 min-x) (< x3 max-x) (> min-y y3) (< min-y y4))
                                                     (and (> x3 min-x) (< x3 max-x) (> max-y y3) (< max-y y4)))
                                                   (or
                                                     (and (> y3 min-y) (< y3 max-y) (> min-x x3) (< min-x x4))
                                                     (and (> y3 min-y) (< y3 max-y) (> max-x x3) (< max-x x4))))) lines)))))

(defn find-contiguous [combos lines]
  (first (filter (partial forms-contiguous-area lines) combos)))

(let [combos (find-combos coordinates) lines (get-lines coordinates) contiguous (find-contiguous combos lines)]
  (println (format "The largest area is %d." (first (first combos))))
  (println (format "The largest contiguous area is %d." (first contiguous)))
  (println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000)))))
