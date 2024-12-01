(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def lists (->> (re-seq #"\d+" (slurp "input.txt"))
                (map parse-long)
                (partition 2)
                (apply mapv vector)
                (map sort)))

(defn count-distance [[list-a list-b]]
  (reduce + (map (comp abs -) list-a list-b)))

(defn count-similarity [[list-a list-b]]
  (let [list-b-freq (frequencies list-b)]
    (reduce + (map #(* (get list-b-freq % 0) %) list-a))))

(println (format "The total distance between the lists is %d." (count-distance lists)))
(println (format "The similarity score of the lists is %d." (count-similarity lists)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

