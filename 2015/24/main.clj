(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def weights (reverse (map parse-long (re-seq #"\d+" (slurp "input.txt")))))

(defn compare-groups [a b]
  (if (= (count a) (count b))
    (if (< (reduce * a) (reduce * b)) a b)
    (if (< (count a) (count b)) a b)))

(defn find-min-group [weights target]
  (loop [queue [{:index 0 :group []}] min-group weights]
    (if (zero? (count queue)) min-group
      (let [{:keys [index group]} (first queue) group-sum (reduce + group)]
        (if (= group-sum target)
          (recur (rest queue) (compare-groups group min-group))
          (if (or (> group-sum target)
                  (> (count group) (count min-group))
                  (>= index (count weights)))
            (recur (rest queue) min-group)
            (recur (concat [{:index (inc index) :group (cons (nth weights index) group)}
                            {:index (inc index) :group group}]
                           (rest queue))
                   min-group)))))))
  
(def quantum-third (reduce * (find-min-group weights (/ (reduce + weights) 3))))
(def quantum-quarter (reduce * (find-min-group weights (/ (reduce + weights) 4))))

(println (format "With three groups, the quantum entanglement of the first group is %d." quantum-third))
(println (format "With four groups, the quantum entanglement of the first group is %d." quantum-quarter))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

