(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def lengths (mapv #(parse-long %) (str/split (str/trim (slurp "input.txt")) #",")))
(def ascii-lengths (conj (mapv int (str/trim (slurp "input.txt"))) 17 31 73 47 23))

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

(def initial-hash (apply-lengths (vec (range 256)) lengths))
(def initial-product (* (first (first initial-hash)) (second (first initial-hash))))

(def sparse-hash (apply-lengths (vec (range 256)) (vec (flatten (replicate 64 ascii-lengths)))))
(def dense-hash (map (fn [xs] (reduce #(bit-xor %1 %2) xs)) (partition 16 (first sparse-hash))))
(def knot-hash (str/join "" (map #(format "%02x" %) dense-hash)))

(println (format "The product of the first two numbers in the processed list is %d." initial-product))
(println (format "The Knot Hash of the input is %s." knot-hash))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

