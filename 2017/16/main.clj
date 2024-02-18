(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def moves (str/split (slurp "input.txt") #","))

(defn spin [x programs]
  (concat (take-last x programs) (drop-last x programs)))

(defn exchange [[a-index b-index] programs]
  (let [a (nth programs a-index) b (nth programs b-index)]
    (map-indexed #(if (= %1 a-index) b (if (= %1 b-index) a %2)) programs)))

(defn partner [a b programs]
  (map #(if (= % a) b (if (= % b) a %)) programs))

(defn perform-move [dance-move programs]
     (case (first dance-move)
       \s (spin (parse-long (re-find #"\d+" dance-move)) programs)
       \x (exchange (map #(parse-long %) (re-seq #"\d+" dance-move)) programs)
       \p (partner (second dance-move) (last dance-move) programs)))

(defn perform-dance [moves programs]
  (reduce #(perform-move %2 %1) programs moves))

(defn count-steps-to-recur [moves programs]
  (loop [next-positions (perform-dance moves programs) n 1]
    (if (= programs next-positions) n (recur (perform-dance moves next-positions) (inc n)))))

(def init-programs (seq "abcdefghijklmnop"))
(def single-dance (str/join (perform-dance moves init-programs)))

(def period (count-steps-to-recur moves init-programs))
(def remainder (rem 1000000000 period))
(def billion-dance (str/join (reduce (fn [programs _] (perform-dance moves programs)) init-programs (range remainder))))

(println (format "After performing the dance once, the configuration of the programs is %s." single-dance))
(println (format "After performing the dance a billion times, their configuration is %s." billion-dance))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

