(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def stream (str/trim (slurp "input.txt")))

(defn score-trash [stream start-index]
  (loop [curr-index start-index trash-score 0]
    (case (get stream curr-index)
      \> (list (inc curr-index) trash-score)
      \! (recur (+ curr-index 2) trash-score)
      (recur (inc curr-index) (inc trash-score)))))

(defn score [stream start-index depth]
  (loop [curr-index start-index group-score depth trash-score 0]
    (case (get stream curr-index)
      \{ (let [[next-index subscore subtrash] (score stream (inc curr-index) (inc depth))]
           (recur next-index (+ group-score subscore) (+ trash-score subtrash)))
      \} (list (inc curr-index) group-score trash-score)
      \< (let [[next-index subscore] (score-trash stream (inc curr-index))]
           (recur next-index group-score (+ trash-score subscore)))

      (if (< curr-index (count stream))
        (recur (inc curr-index) group-score trash-score)))))

(let [[_ group-score trash-score] (score stream 1 1)]
  (println (format "The group score of the stream is %d." group-score))
  (println (format "The trash score of the stream is %d." trash-score)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

