(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def digits (mapv parse-long (str/split (str/trim (slurp "input.txt")) #"")))

(defn sum-matches [collection step]
  (reduce (fn [sum [idx ch]]
            (if (= ch (get collection (mod (+ idx step) (count collection))))
              (+ sum ch)
              sum))
          0
          (map-indexed vector collection)))

(def initial-captcha (sum-matches digits 1))
(def final-captcha (sum-matches digits (/ (count digits) 2)))

(println (format "The result of the initial captcha is %s." initial-captcha))
(println (format "The result of the final captcha is %s." final-captcha))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))
