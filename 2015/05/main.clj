(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def strings (str/split (slurp "input.txt") #"\n"))

(defn is-nice-old [string]
  (and (>= (count (re-seq #"[aeiou]" string)) 3)
       (re-seq #"([a-z])\1" string)
       (not (re-seq #"ab|cd|pq|xy" string))))

(defn is-nice-new [string]
  (and (re-seq #"([a-z][a-z]).*\1" string)
       (re-seq #"([a-z]).\1" string)))

(def num-nice-old (count (filter is-nice-old strings)))
(def num-nice-new (count (filter is-nice-new strings)))

(println (format "Under the old rules, the number of nice strings is %d." num-nice-old))
(println (format "Under the new rules, the number of nice strings is %d." num-nice-new))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

