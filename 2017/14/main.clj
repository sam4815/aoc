(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def hash-inputs (map-indexed #(str %2 "-" %1) (take 128 (repeat (str/trim (slurp "input.txt"))))))

(defn reverse-range [start-xs start-index end-index]
  (loop [xs start-xs i start-index j end-index]
    (if (or (= i j) (and (not= i start-index) (= i (mod (inc j) (count xs)))))
      xs
      (let [i-element (get xs i) j-element (get xs j)]
        (recur (assoc (assoc xs j i-element) i j-element) (mod (inc i) (count xs)) (mod (dec j) (count xs)))))))

(defn apply-length [[xs position] [skip-size length]]
  (list
    (reverse-range xs position (mod (+ position (max (dec length) 0)) (count xs)))
    (mod (+ position skip-size length) (count xs))))

(defn apply-lengths [xs lengths]
  (reduce apply-length (list xs 0) (map-indexed #(list %1 %2) lengths)))

(defn apply-knot-hash [string]
  (let [ascii-lengths (conj (mapv int string) 17 31 73 47 23)
        sparse-hash (apply-lengths (vec (range 256)) (vec (flatten (replicate 64 ascii-lengths))))
        dense-hash (map (fn [xs] (reduce #(bit-xor %1 %2) xs)) (partition 16 (first sparse-hash)))]
    (str/join "" (map #(format "%08d" (parse-long (Long/toString % 2))) dense-hash))))

(defn find-adjacent [[x y] hashes]
  (filter (fn [[i j]] (and (>= i 0) (>= j 0) (< i (count hashes)) (< j (count hashes))))
          (list `(~x ~(+ y 1)) `(~x ~(- y 1)) `(~(+ x 1) ~y) `(~(- x 1) ~y))))

(defn find-region [[init-x init-y] hashes]
  (loop [to-explore (list `(~init-x ~init-y)) explored {}]
    (if (= (count to-explore) 0) explored
      (let [squares (filter (fn [[x y]] (= (nth (nth hashes x) y) \1)) to-explore)
            square-hashes (map #(hash %) squares)]
        (recur (set (filter #(not (contains? explored (hash %))) (apply concat (map #(find-adjacent % hashes) squares))))
               (reduce #(assoc %1 %2 true) explored square-hashes))))))

(def hashed-inputs (map #(apply-knot-hash %) hash-inputs))
(def square-count (reduce + (map #(get (frequencies %) \1) hashed-inputs)))
(def regions (flatten (map (fn [x] (map #(find-region `(~x ~%) hashed-inputs) (range 128))) (range 128))))

(println (format "The number of squares used is %d." square-count))
(println (format "The number of distinct regions is %d." (- (count (set regions)) 1)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

