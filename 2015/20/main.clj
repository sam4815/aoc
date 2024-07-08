(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def required-presents (parse-long (str/trim (slurp "input.txt"))))

(defn visit-houses [target house-limit multiplier]
  (loop [step 1 house-multiple 1 houses {}]
    (let [house-number (* step house-multiple)
          num-presents (+ (get houses house-number 0) (* multiplier step))]
      (if (>= num-presents target) house-number
        (if (or (> house-multiple house-limit) (> house-number 1000000))
          (recur (inc step) 1 houses)
          (recur step (inc house-multiple) (assoc houses house-number num-presents)))))))

(println (format "When the elves visit infinite houses, the lowest house number is %d." (visit-houses required-presents 5000 10)))
(println (format "When the elves visit 50 houses, the lowest house number is %d." (visit-houses required-presents 50 11)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

