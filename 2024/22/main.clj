(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def secrets (map parse-long (re-seq #"\d+" (slurp "input.txt"))))

(defn prune [a] (mod a 16777216))
(defn step-1 [x] (->> (* x 64) (bit-xor x) prune))
(defn step-2 [x] (->> (quot x 32) (bit-xor x) prune))
(defn step-3 [x] (->> (* x 2048) (bit-xor x) prune))

(defn next-secret [x] (->> x step-1 step-2 step-3))

(defn get-nth-secrets [n x]
  (reduce (fn [[secrets secret] _] [(conj secrets secret) (next-secret secret)])
          [[x] (next-secret x)]
          (range n)))

(defn get-sequence [n secrets]
  (mapv #(get (get secrets (+ n %)) :diff) (range 4)))

(defn find-sequence [secrets changes]
  (loop [n 0]
    (if (> n (- (count secrets) 4)) 0
      (if (and (= (get changes 0) (get (get secrets (+ n 0)) :diff))
               (= (get changes 1) (get (get secrets (+ n 1)) :diff))
               (= (get changes 2) (get (get secrets (+ n 2)) :diff))
               (= (get changes 3) (get (get secrets (+ n 3)) :diff)))
        (get (get secrets (+ n 3)) :price)
        (recur (inc n))))))

(defn count-changes [all-secrets changes]
  (reduce + (map #(find-sequence % changes) all-secrets)))

(def nth-secrets (mapv (fn [secret] (get-nth-secrets 2000 secret)) secrets))

(def diffs (map (fn [arr] (vec (map-indexed (fn [i secret]
                                              {:price (mod secret 10)
                                               :diff (when (> i 0) (- (mod secret 10) (mod (get (first arr) (dec i)) 10)))})
                                            (first arr)))) nth-secrets))

(defn find-most-bananas [diffs]
  (loop [n 0 most-bananas 0]
    (if (> n 1996) most-bananas
      (recur (inc n) (max most-bananas (count-changes diffs (get-sequence n (first diffs))))))))

(println (format "The sum of the 2000th numbers is %d." (reduce + (map #(last (first %)) nth-secrets))))
(println (format "The most bananas you can get is %d." (find-most-bananas diffs)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

