(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def triangle-rows (map (fn [line] (map #(parse-long %) (str/split (str/trim line) #"\s+")))
                        (str/split (slurp "input.txt") #"\n")))

(def triangle-columns (apply concat (map (fn [rows] (map-indexed (fn [i _] (map #(nth % i) rows)) rows))
                                         (partition 3 triangle-rows))))

(defn is-valid-triangle [triangle] (let [[a b c] (sort triangle)] (> (+ a b) c)))

(println (format "Read as rows, there are %d valid triangles." (count (filter is-valid-triangle triangle-rows))))
(println (format "Read as columns, there are %d valid triangles." (count (filter is-valid-triangle triangle-columns))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

