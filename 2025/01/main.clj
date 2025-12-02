(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def rotations (->> (str/split (slurp "input.txt") #"\n")
                (map #(identity [(re-find #"[LR]" %) (parse-long (re-find #"\d+" %))]))))

(defn count-rotations [rotations]
  (reduce (fn [[zero-count turn-count current] [direction magnitude]]
            (let [turns (quot magnitude 100)
                  remainder (rem magnitude 100)
                  result (case direction "L" (- current remainder) (+ current remainder))
                  on-zero (= (mod result 100) 0)
                  passed-zero (and (not= current 0) (or (>= result 100) (<= result 0)))]
              [(+ zero-count (if on-zero 1 0)) (+ turn-count turns (if passed-zero 1 0)) (mod result 100)]))
          [0 0 50]
          rotations))

(def counts (count-rotations rotations))

(println (format "The password is %d." (first counts)))
(println (format "The actual password is %d." (second counts)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

