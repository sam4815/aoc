(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def firewall (map (fn [line] (map #(parse-long %) (re-seq #"\d+" line))) (str/split (slurp "input.txt") #"\n")))

(defn find-caught-layers [firewall packet-delay]
  (filter #(= 0 (mod (+ (first %) packet-delay) (* (dec (second %)) 2))) firewall))

(def severity (reduce + (map #(* (first %) (second %)) (find-caught-layers firewall 0))))

(defn find-delay [firewall]
  (loop [packet-delay 1]
    (if (= (count (find-caught-layers firewall packet-delay)) 0)
      packet-delay
      (recur (inc packet-delay)))))

(println (format "The severity of the immediate trip is %d." severity))
(println (format "To pass through undetected, the trip should be delayed by %d picoseconds." (find-delay firewall)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

