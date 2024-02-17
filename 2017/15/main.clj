(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def generator-starts (map #(parse-long %) (re-seq #"\d+" (slurp "input.txt"))))

(defn next-value [value factor multiple]
  (rem (* value factor) 2147483647))

(defn next-rec-value [value factor multiple]
  (loop [next-val (next-value value factor multiple)]
    (if (= (mod next-val multiple) 0)
      next-val
      (recur (next-value next-val factor multiple)))))

(defn compare-bits [value-one value-two]
  (= (bit-and 65535 value-one) (bit-and 65535 value-two)))

(defn next-with-total [[value-one value-two total] func]
  (let [next-one (func value-one 16807 4) next-two (func value-two 48271 8)]
    (if (compare-bits next-one next-two)
      (list next-one next-two (inc total))
      (list next-one next-two total))))

(def forty-count (reduce (fn [acc _] (next-with-total acc next-value)) `(~@generator-starts 0) (range 40000000)))
(def five-count (reduce (fn [acc _] (next-with-total acc next-rec-value)) `(~@generator-starts 0) (range 5000000)))

(println (format "After 40 million pairs, the final count is %d." (last forty-count)))
(println (format "After 5 million pairs, the final count is %d." (last five-count)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

