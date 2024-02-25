(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def messages (str/split (slurp "input.txt") #"\n"))
(def message-columns (map (fn [i] (map #(nth % i) messages)) (range (count (first messages)))))

(def corrected-message (str/join (map (fn [message] (key (apply max-key val (frequencies message)))) message-columns)))
(def original-message (str/join (map (fn [message] (key (apply min-key val (frequencies message)))) message-columns)))

(println (format "The error-corrected message is %s." corrected-message))
(println (format "The original message is %s." original-message))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

