(ns main
  (:require [clojure.string :as str]
            [clojure.math :as math]))

(def start-time (System/currentTimeMillis))

(def enhancements
  (->> (str/split (slurp "input.txt") #"\n")
       (map #(str/split % #" => "))
       (map (fn [rule] (map #(str/split % #"/") rule)))
       (map (fn [rule] (map (fn [pattern] (mapv #(vec (seq %)) pattern)) rule)))))

(defn flip-horizontal [pixels] (reverse pixels))

(defn flip-vertical [pixels] (map #(reverse %) pixels))

(defn rotate-90 [pixels]
  (flip-horizontal (map-indexed (fn [index row] (map #(nth % index) pixels)) pixels)))

(defn get-permutations [pixels]
  [pixels (rotate-90 pixels) (rotate-90 (rotate-90 pixels)) (rotate-90 (rotate-90 (rotate-90 pixels)))
   (flip-horizontal pixels) (rotate-90 (flip-horizontal pixels))
   (flip-vertical pixels) (rotate-90 (flip-vertical pixels))])

(defn find-enhancement [pixels enhancements]
  (let [permutations (get-permutations pixels)]
    (second (first (filter (fn [[pattern]] (some #(= % pattern) permutations)) enhancements)))))

(defn chunk-by [pixels n]
  (let [num-chunks (/ (count pixels) n)]
    (apply concat
           (mapv (fn [x] (mapv
                           (fn [y] (mapv
                                     (fn [row] (subvec row (* y n) (* (inc y) n)))
                                     (subvec pixels (* x n) (* (inc x) n))))
                           (range num-chunks)))
                 (range num-chunks)))))

(defn join [pixels]
  (let [size (int (math/sqrt (count pixels)))]
    (vec (apply concat
                (mapv (fn [squares]
                        (mapv (fn [x] (vec (apply concat
                                                  (mapv #(nth % x) squares))))
                              (range (count (first squares)))))
                      (mapv vec (partition size pixels)))))))

(defn break-up [pixels] (if (= (mod (count pixels) 2) 0) (chunk-by pixels 2) (chunk-by pixels 3)))

(defn enhance [pixels] (join (map #(find-enhancement % enhancements) (break-up pixels))))

(defn enhance-n [init-pixels n]
  (loop [pixels init-pixels i 0]
    (if (= n i) pixels (recur (enhance pixels) (inc i)))))

(defn count-pixels-on [pixels]
  (reduce + (map (fn [row] (reduce #(if (= \# %2) (inc %1) %1) 0 row)) pixels)))

(def initial-pattern [[\. \# \.] [\. \. \#] [\# \# \#]])

(def five-pixels-count (count-pixels-on (enhance-n initial-pattern 5)))
(def eighteen-pixels-count (count-pixels-on (enhance-n initial-pattern 18)))

(println (format "After 5 iterations, there are %d pixels on." five-pixels-count))
(println (format "After 18 iterations, there are %d pixels on." eighteen-pixels-count))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

