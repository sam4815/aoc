(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def eqs (->> (str/split (slurp "input.txt") #"\n")
              (map #(map parse-long (re-seq #"\d+" %)))
              (map (fn [[x & xs]] [x xs]))))

(defn is-valid [operators [target numbers]]
  (loop [queue [[(first numbers) (rest numbers)]]]
    (when (not-empty queue)
      (let [[total remaining] (first queue)]
        (or (and (= total target) (empty? remaining))
            (if (or (> total target) (empty? remaining)) (recur (rest queue))
              (recur (concat (map (fn [op] [(op total (first remaining)) (rest remaining)]) operators)
                             (rest queue)))))))))

(defn sum-eqs [eqs]
  (reduce + (map first eqs)))

(defn || [a b]
  (parse-long (str a b)))

(println (format "The initial calibration result is %d." (sum-eqs (filter (partial is-valid [* +]) eqs))))
(println (format "The final calibration result is %d." (sum-eqs (filter (partial is-valid [* + ||]) eqs))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

