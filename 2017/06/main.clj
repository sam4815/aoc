(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def banks (mapv #(parse-long %) (str/split (slurp "input.txt") #"\s")))

(defn find-bank-index [banks]
  (let [max-bank (apply max banks)]
    (first (keep-indexed #(when (= %2 max-bank) (vector %1 %2)) banks))))

(defn redistribute [banks]
  (let [[bank-index bank-value] (find-bank-index banks)
        remainder (rem bank-value (count banks))
        full-cycles (quot bank-value (count banks))]
    (vec (keep-indexed
           #(if (<= (mod (- %1 bank-index) (count banks)) remainder) (+ %2 (inc full-cycles)) (+ %2 full-cycles))
           (assoc banks bank-index -1)))))

(defn find-first-repeat [banks]
  (loop [current-banks banks current-steps 1 visited (set (vector (hash banks)))]
    (let [next-banks (redistribute current-banks) next-hash (hash next-banks)]
      (if (contains? visited next-hash)
        (vector current-steps next-banks)
        (recur next-banks (inc current-steps) (conj visited next-hash))))))

(defn count-repeat [banks]
  (loop [current-banks banks current-steps 1]
    (let [next-banks (redistribute current-banks)]
      (if (= (hash banks) (hash next-banks))
        current-steps
        (recur next-banks (inc current-steps))))))

(let [[num-steps repeating-banks] (find-first-repeat banks)]
  (println (format "%d redistribution cycles are required to produce a duplicate configuration." num-steps))
  (println (format "There are %d cycles in the infinite loop." (count-repeat repeating-banks))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

