(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def position (map parse-long (re-seq #"\d+" (slurp "input.txt"))))

(defn find-n [[row column]]
  (reduce + (concat '(1) (range 1 row) (range (inc row) (+ row column)))))

(defn nth-code [n]
  (mod (* 20151125 (.modPow (biginteger 252533) (biginteger (dec n)) (biginteger 33554393))) 33554393))

(println (format "The code for the machine is %s." (nth-code (find-n position))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

