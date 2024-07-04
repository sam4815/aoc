(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def container-sizes (map parse-long (str/split (slurp "input.txt") #"\n")))

(defn find-combos [sizes target]
  (let [smaller (filter (partial > target) sizes)]
    (concat (map list (filter (partial = target) sizes))
            (for [i (range (count smaller))
                  tail (find-combos (drop (inc i) smaller) (- target (nth smaller i)))]
              (cons (nth smaller i) tail)))))

(def container-combos (find-combos container-sizes 150))

(def combo-sizes (map count container-combos))
(def min-combos (filter (partial = (apply min combo-sizes)) combo-sizes))

(println (format "There are %d container combinations." (count container-combos)))
(println (format "There are %d ways of filling the minimum number of containers." (count min-combos)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

