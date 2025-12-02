(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def ranges (->> (str/split (slurp "input.txt") #",")
                 (map #(re-seq #"\d+" %))))

(defn find-invalid [factor [start end]]
  (let [length (count start)
        parsed-start (parse-long start)
        parsed-end (parse-long end)
        initial (if (zero? (rem length factor))
                  (subs start 0 (quot length factor))
                  (str "1" (apply str (repeat (quot length factor) "0"))))]
    (loop [candidate initial invalids []]
      (let [invalid (parse-long (apply str (repeat factor candidate)))]
        (if (> invalid parsed-end)
          invalids
          (recur
            (str (inc (parse-long candidate)))
            (if (< invalid parsed-start) invalids (conj invalids invalid))))))))


(defn find-all-invalid [[start end]]
  (let [factors (range 2 (inc (count end)))]
    (distinct (mapcat #(find-invalid % [start end]) factors))))

(defn sum-invalids [invalids]
  (reduce + (map #(reduce + %) invalids)))

(def invalid-doubles (map (partial find-invalid 2) ranges))

(def invalid-tuples (map find-all-invalid ranges))

(println (format "The sum of the invalid IDs is %d." (sum-invalids invalid-doubles)))
(println (format "The sum of the invalid IDs using the new rules is %d." (sum-invalids invalid-tuples)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

