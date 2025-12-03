(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def banks (->> (str/split (slurp "input.txt") #"\n")
                (mapv #(mapv parse-long (str/split % #"")))))

(defn parse-longs [nums]
  (parse-long (apply str nums)))

(defn find-max-joltage [init-target init-bank]
  (loop [target init-target bank init-bank digits []]
    (if (zero? target) digits 
      (let [[max-index max-value] (apply max-key second (reverse (map-indexed vector (drop-last (dec target) bank))))]
        (recur (dec target) (subvec bank (inc max-index)) (conj digits max-value))))))

(def two-joltage (reduce + (map parse-longs (map (partial find-max-joltage 2) banks))))
(def twelve-joltage (reduce + (map parse-longs (map (partial find-max-joltage 12) banks))))

(println (format "The output joltage is %d." two-joltage))
(println (format "The new output joltage is %d." twelve-joltage))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

