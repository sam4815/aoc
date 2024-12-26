(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def schematics (->> (str/split (slurp "input.txt") #"\n\n")
                     (mapv (fn [schematic] (mapv (comp vec seq) (str/split schematic #"\n"))))))

(defn transpose [matrix]
  (apply mapv vector matrix))

(defn find-heights [schematic]
  (->> (transpose schematic)
       (map (fn [line] (count (filter #{\#} line))))))

(defn find-combos [schematics]
  (let [grouped (group-by (comp first first) schematics)]
    (->> (for [t-key (map find-heights (get grouped \#)) t-lock (map find-heights (get grouped \.))]
           (map + t-key t-lock))
         (filter (fn [pair] (every? (fn [num-filled] (<= num-filled 7)) pair)))
         count)))

(println (format "There are %d unique pairs." (find-combos schematics)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

