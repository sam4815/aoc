(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def first-row (vec (str/trim (slurp "input.txt"))))

(defn find-next-row [row]
  (mapv (fn [i] (if (= (nth row (dec i) \.) (nth row (inc i) \.)) \. \^)) (range (count row))))

(defn count-safe-row [row] (get (frequencies row) \.))

(defn count-safe-room [first-row target]
  (loop [row (find-next-row first-row) n 1 num-safe (count-safe-row first-row)]
    (if (= n target) num-safe (recur (find-next-row row) (inc n) (+ num-safe (count-safe-row row))))))

(println (format "In the first 40 rows, there are %d safe tiles." (count-safe-room first-row 40)))
(println (format "In the first 400000 rows, there are %d safe tiles." (count-safe-room first-row 400000)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

