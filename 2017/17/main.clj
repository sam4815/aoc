(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def step-size (parse-long (str/trim (slurp "input.txt"))))

(defn step [[state position] value step-size]
  (let [next-position (inc (mod (+ position step-size) (count state)))
        [before after] (split-at next-position state)]
    (vector (concat before [value] after) next-position)))

(defn step-pos [position state-size step-size]
  (inc (mod (+ position step-size) state-size)))
    
(defn spinlock [step-size num-steps]
  (reduce #(step %1 (inc %2) step-size) '([0] 0) (range num-steps)))

(defn find-element-after [state element]
  (nth state (inc (.indexOf state element))))

(defn find-zero-after [step-size num-steps]
  (loop [after-zero 0 n 1 pos 0]
    (if (> n num-steps) after-zero
      (let [next-pos (step-pos pos n step-size)]
        (if (= next-pos 1) (recur n (inc n) next-pos) (recur after-zero (inc n) next-pos))))))

(def after-last-written (find-element-after (first (spinlock step-size 2017)) 2017))
(def zero-after-billions (find-zero-after step-size 50000000))

(println (format "After 2017 iterations, the value after 2017 is %d." after-last-written))
(println (format "After 50000000 iterations, the value after 0 is %d." zero-after-billions))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

