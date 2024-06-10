(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def dimensions (->> (str/split (slurp "input.txt") #"\n")
                     (map #(str/split % #"x"))
                     (map (fn [nums] (map #(parse-long %) nums)))))

(defn calculate-paper-required [[l w h]]
  (let [[x y] (sort `(~l ~w ~h))]
    (+ (* 2 l w)
       (* 2 w h)
       (* 2 h l)
       (* x y))))

(defn calculate-ribbon-required [[l w h]]
  (let [[x y] (sort `(~l ~w ~h))]
    (+ (* l w h)
       (* 2 (+ x y)))))

(def paper-sq-ft (reduce + (map calculate-paper-required dimensions)))
(def ribbon-sq-ft (reduce + (map calculate-ribbon-required dimensions)))

(println (format "The elves need %d square feet of wrapping paper." paper-sq-ft))
(println (format "The elves need %d square feet of ribbon." ribbon-sq-ft))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

