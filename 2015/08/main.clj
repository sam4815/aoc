(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def strings (str/split (slurp "input.txt") #"\n"))

(defn memory-count [string]
  (-> string
      (subs 1 (dec (count string)))
      (str/replace #"\\\\" "-")
      (str/replace #"\\\"" "-")
      (str/replace #"[\\]x[0-9a-f]{2}" "-")
      count))

(defn encoded-count [string]
  (-> string
      (str/replace #"[\\]" "$0$0")
      (str/replace #"\"" "\\\\$0")
      count
      (+ 2)))

(def memory-difference (- (reduce + (map count strings)) (reduce + (map memory-count strings))))
(def encoded-difference (- (reduce + (map encoded-count strings)) (reduce + (map count strings))))

(println (format "The memory difference in characters is %d." memory-difference))
(println (format "The encoding difference in characters is %d." encoded-difference))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

