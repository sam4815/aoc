(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def parts (str/split (slurp "input.txt") #"\n\n"))

(def ranges (partition 2 (map parse-long (re-seq #"\d+" (first parts)))))
(def ingredients (map parse-long (re-seq #"\d+" (second parts))))

(defn is-fresh [ranges ingredient]
  (some (fn [[a b]] (and (>= ingredient a) (<= ingredient b))) ranges))

(defn count-ranges [ranges]
  (reduce
    (fn [[total current] [a b]]
      (if (< current a) [(+ total (- b a) 1) b]
        (if (> current b) [total current] [(+ total (- b current)) b])))
    [0 0]
    (sort-by first ranges)))

(println (format "There are %d fresh ingredients." (count (filter (partial is-fresh ranges) ingredients))))
(println (format "There are %d fresh ingredient IDs." (first (count-ranges ranges))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

