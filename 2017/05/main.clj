(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def jumps (vec (map #(parse-long %) (str/split (slurp "input.txt") #"\n"))))

(defn regular-jump [position jumps]
  (list
    (+ position (nth jumps position))
    (assoc jumps position (inc (nth jumps position)))))

(defn strange-jump [position jumps]
  (def offset (nth jumps position))
  (list
    (+ position offset)
    (assoc jumps position (if (>= offset 3) (dec offset) (inc offset)))))

(defn count-escape-steps [position jumps steps jump-func]
  (loop [current-position position
         current-jumps jumps
         steps 0]
    (let [[next-position next-jumps] (jump-func current-position current-jumps)]
      (if (>= next-position (count current-jumps))
        (+ steps 1)
        (recur next-position next-jumps (inc steps))))))

(println (format "With a regular jump, it takes %d steps to reach the exit." (count-escape-steps 0 jumps 0 regular-jump)))
(println (format "With a strange jump, it takes %d steps to reach the exit." (count-escape-steps 0 jumps 0 strange-jump)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

