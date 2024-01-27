(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def digits (->> (str/split (slurp "input.txt") #"\n")
                 (map (fn [line] (map #(parse-long %) (str/split line #"\s"))))))

(defn get-divisors [num nums]
  (->> nums
       (map #(if (= (mod num %) 0) (/ num %) 0))
       (filter #(> % 1))))

(defn find-division [nums]
  (->> nums
       (map #(get-divisors % nums))
       (remove empty?)
       first
       first))

(def checksum (reduce + (map #(- (apply max %) (apply min %)) digits)))
(def divisible-sum (reduce + (map #(find-division %) digits)))

(println (format "The checksum of the spreadsheet is %s." checksum))
(println (format "The sum of the evenly divisible values is %s." divisible-sum))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

