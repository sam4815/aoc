(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def machines (->> (str/split (slurp "input.txt") #"\n\n")
                   (mapv #(mapv parse-long (re-seq #"\d+" %)))))

(def fixed-machines (mapv (fn [[ax ay bx by tx ty]]
                            [ax ay bx by (+ tx 10000000000000) (+ ty 10000000000000)]) machines))

(defn gcd [a b]
  (if (zero? b) a (gcd b (mod a b))))

(defn lcm [a b]
  (/ (* a b) (gcd a b)))

(defn lcm-factors [a b]
  [(/ (lcm a b) a) (/ (lcm a b) b)])

(defn solve [[ax ay bx by tx ty]]
  (let [[x-factor y-factor] (lcm-factors ax ay)
        b-coefficient (- (* bx x-factor) (* by y-factor))
        rhs (- (* tx x-factor) (* ty y-factor))
        b (/ rhs b-coefficient)]
    (when (integer? b) [(/ (- tx (* b bx)) ax) b])))

(defn cost [machine]
  (let [presses (solve machine)]
    (if (some? presses) (+ (* 3 (first presses)) (second presses)) 0)))

(println (format "%d tokens are required to win all prizes." (reduce + (map cost machines))))
(println (format "Once fixed, %d tokens are required to win all prizes." (reduce + (map cost fixed-machines))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

