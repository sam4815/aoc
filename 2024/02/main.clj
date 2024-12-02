(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def reports (->> (str/split (slurp "input.txt") #"\n")
                  (map #(mapv parse-long (re-seq #"\d+" %)))))

(defn is-safe [report]
  (let [differences (map - (rest report) report)]
    (or (every? #{1 2 3} differences)
        (every? #{-1 -2 -3} differences))))

(defn has-safe-permutation [report]
  (some is-safe
        (map-indexed (fn [i _]
                       (concat (subvec report 0 i)
                               (subvec report (inc i))))
                     report)))

(println (format "Normally, %d reports are safe." (count (filter is-safe reports))))
(println (format "Using the dampener, %d reports are safe." (count (filter has-safe-permutation reports))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

