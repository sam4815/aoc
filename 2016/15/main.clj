(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def discs (map (fn [line] (map #(parse-long %) (re-seq #"\d+" line)))(str/split (slurp "input.txt") #"\n")))
(def discs-plus (concat discs `((~(inc (count discs)) 11 0 0))))

(defn find-time [discs]
  (loop [seconds 0]
    (if (every? (fn [[target modulo _ offset]] (= 0 (mod (+ offset target seconds) modulo))) discs) seconds
      (recur (inc seconds)))))

(println (format "The first time the button can be pressed is at %d seconds." (find-time discs)))
(println (format "The second time the button can be pressed is at %d seconds." (find-time discs-plus)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

