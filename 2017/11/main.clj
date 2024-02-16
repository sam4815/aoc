(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def steps (str/split (str/trim (slurp "input.txt")) #","))

(defn get-distance [[x y]] (+ x (/ (- y x) 2)))

(defn move [[x y] step]
  (case step
    "n" `(~x ~(+ y 2)) "ne" `(~(+ x 1) ~(+ y 1)) "nw" `(~(- x 1) ~(+ y 1))
    "s" `(~x ~(- y 2)) "sw" `(~(- x 1) ~(- y 1)) "se" `(~(+ x 1) ~(- y 1))))

(defn move-add-max [[x y max-distance] step]
  (let [[next-x next-y] (move `(~x ~y) step)]
    `(~next-x ~next-y ~(max max-distance (get-distance `(~next-x ~next-y))))))

(defn get-position [steps] (reduce move-add-max '(0, 0, 0) steps))

(let [[x y max-distance] (get-position steps)]
  (println (format "The minimum number of steps to the child process is %d." (get-distance `(~x ~y))))
  (println (format "The furthest number of steps the child process reached was %d steps." max-distance)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

