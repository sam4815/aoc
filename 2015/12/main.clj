(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def json (str/trim (slurp "input.txt")))

(defn sum-object [string start-index omit-red]
  (loop [n start-index sum 0 has-red false]
    (case (get string n)
      \{ (let [[object-sum last-n] (sum-object string (inc n) omit-red)]
           (recur last-n (+ sum object-sum) has-red))
      (\0 \1 \2 \3 \4 \5 \6 \7 \8 \9 \-) (let [match (first (re-seq #"-?\d+" (subs string n)))]
                                           (recur (+ n (count match)) (+ sum (parse-long match)) has-red))
      \} [(if (and omit-red has-red) 0 sum) (inc n)]
      (recur (inc n) sum (or has-red (str/starts-with? (subs string n) ":\"red"))))))

(println (format "The sum of all numbers in the document is %d." (first (sum-object json 1 false))))
(println (format "Omitting red objects, the sum of all numbers in the document is %d." (first (sum-object json 1 true))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

