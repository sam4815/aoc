(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def instructions (map #(seq %) (str/split (str/trim (slurp "input.txt")) #"\n")))

(def imagined-keypad [[1 2 3] [4 5 6] [7 8 9]])
(def real-keypad [[0 0 1 0 0] [0 2 3 4 0] [5 6 7 8 9] [0 \A \B \C 0] [0 0 \D 0 0]])

(defn next-position [[x y] direction]
  (case direction
    \U `(~x ~(dec y)) \D `(~x ~(inc y))
    \L `(~(dec x) ~y) \R `(~(inc x) ~y)))

(defn validate-position [[x y] keypad] (not= 0 (nth (nth keypad y []) x 0)))

(defn follow-instructions [instructions keypad]
  (loop [position '(1 1) index 0]
    (if (>= index (count instructions)) (nth (nth keypad (second position)) (first position))
      (let [next-position (next-position position (nth instructions index))]
        (recur (if (validate-position next-position keypad) next-position position) (inc index))))))

(def imagined-bathroom-code (str/join (map #(follow-instructions % imagined-keypad) instructions)))
(def real-bathroom-code (str/join (map #(follow-instructions % real-keypad) instructions)))

(println (format "The imagined bathroom code is %s." imagined-bathroom-code))
(println (format "The real bathroom code is %s." real-bathroom-code))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

