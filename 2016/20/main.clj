(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def ranges (map (fn [line] (map #(parse-long %) (re-seq #"\d+" line)))
                        (str/split (slurp "input.txt") #"\n")))

(defn find-next-unblocked [start ranges]
  (loop [x start]
    (let [enclosing-range (first (filter (fn [[low high]] (and (>= x low) (<= x high))) ranges))]
      (if (nil? enclosing-range) x
        (recur (inc (last enclosing-range)))))))

(defn find-next-blocked [start ranges]
  (first (first (sort-by first (filter (fn [[low]] (> low start)) ranges)))))

(defn count-unblocked [ranges]
  (loop [x (find-next-unblocked 0 ranges) sum 0]
    (let [next-blocked (find-next-blocked x ranges)
          next-unblocked (find-next-unblocked next-blocked ranges)
          diff (- next-blocked x)]
      (if (>= next-unblocked 4294967295) (+ diff sum) (recur next-unblocked (+ diff sum))))))

(println (format "The lowest IP address that is not blocked is %d." (find-next-unblocked 0 ranges)))
(println (format "The number of IP addresses that are allowed is %d." (count-unblocked ranges)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

