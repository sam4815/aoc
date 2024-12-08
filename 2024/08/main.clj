(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def city-map (mapv (comp vec seq) (str/split (slurp "input.txt") #"\n")))

(defn find-symbols [city-map]
  (vec (filter some? (for [x (range (count (first city-map))) y (range (count city-map))]
                       (when (not= \. (get-in city-map [x y])) [x y (get-in city-map [x y])])))))

(defn within-city [city-map [x y]]
  (and (>= x 0) (>= y 0) (< x (count (first city-map))) (< y (count city-map))))

(defn find-antinodes [group search-space]
  (apply concat
         (for [[x y sym] group [i j _] (remove #{[x y sym]} group)]
           (let [x-diff (- i x) y-diff (- j y)]
             (map (fn [i] [(- x (* i x-diff)) (- y (* i y-diff))]) search-space)))))

(defn count-antinodes [city-map search-space]
  (let [symbols (group-by last (find-symbols city-map))]
    (->> (vals symbols)
         (map (fn [group] (find-antinodes group search-space)))
         (apply concat)
         distinct
         (filter (partial within-city city-map))
         count)))

(println (format "Under the initial model, there are %d unique antinode locations." (count-antinodes city-map [1])))
(println (format "In reality, there are %d unique antinode locations." (count-antinodes city-map (range 100))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

