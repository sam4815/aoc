(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def initial-state (str/trim (slurp "input.txt")))

(defn invert [state] (map #(if (= \1 %) \0 \1) state))

(defn apply-dragon-curve [initial-state target-length]
  (loop [state initial-state]
    (if (>= (count state) target-length) (subs state 0 target-length)
      (recur (str state "0" (apply str (invert (reverse state))))))))

(defn reduce-pairs [state]
  (map (fn [[a b]] (if (= a b) \1 \0)) (partition 2 state)))

(defn calculate-checksum [state]
  (loop [checksum (reduce-pairs state)]
    (if (odd? (count checksum)) checksum (recur (reduce-pairs checksum)))))

(def first-checksum (str/join (calculate-checksum (apply-dragon-curve initial-state 272))))
(def second-checksum (str/join (calculate-checksum (apply-dragon-curve initial-state 35651584))))

(println (format "The checksum for the first disk is %s." first-checksum))
(println (format "The checksum for the second disk is %s." second-checksum))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

