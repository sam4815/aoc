(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def instructions (slurp "input.txt"))
(def conditional-instructions (str/replace instructions #"don't\(\)(.|\n)*?do\(\)" ""))

(defn sum-muls [instructions]
  (->> (re-seq #"mul\((\d+),(\d+)\)" instructions)
       (map #(reduce * (map parse-long (rest %))))
       (reduce +)))

(println (format "The sum of the mul instructions is %d." (sum-muls instructions)))
(println (format "With conditionals, the sum of the muls is %d." (sum-muls conditional-instructions)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

