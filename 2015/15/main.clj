(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def cookies (->> (str/split (slurp "input.txt") #"\n")
                  (map #(map parse-long (re-seq #"-?\d+" %)))))

(defn sum-combos [n target]
  (if (= n 1)
    (list (list target))
    (for [head (range (inc target))
          tail (sum-combos (dec n) (- target head))]
      (cons head tail))))

(defn score-cookie [[a b c d] [a-num b-num c-num d-num]]
  [(* (max 0 (+ (* (nth a 0) a-num) (* (nth b 0) b-num) (* (nth c 0) c-num) (* (nth d 0) d-num)))
      (max 0 (+ (* (nth a 1) a-num) (* (nth b 1) b-num) (* (nth c 1) c-num) (* (nth d 1) d-num)))
      (max 0 (+ (* (nth a 2) a-num) (* (nth b 2) b-num) (* (nth c 2) c-num) (* (nth d 2) d-num)))
      (max 0 (+ (* (nth a 3) a-num) (* (nth b 3) b-num) (* (nth c 3) c-num) (* (nth d 3) d-num))))
   (+ (* (nth a 4) a-num) (* (nth b 4) b-num) (* (nth c 4) c-num) (* (nth d 4) d-num))])

(def scored-cookies (map (partial score-cookie cookies) (sum-combos 4 100)))
(def best-cookie (apply max (map first scored-cookies)))
(def best-500-calorie-cookie (apply max (map first (filter #(= 500 (second %)) scored-cookies))))

(println (format "The score of the best cookie is %d." best-cookie))
(println (format "The score of the best 500-calorie cookie is %d." best-500-calorie-cookie))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

