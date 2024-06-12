(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def instructions (->> (str/split (slurp "input.txt") #"\n")
                       (map (fn [line] [(re-find #"^[a-z]+(?:\s[a-z]+)?" line)
                                        (map #(parse-long %) (re-seq #"\d+" line))]))))

(defn get-lit-fn [string]
  (case string
    "turn on" (fn [cell] true)
    "turn off" (fn [cell] false)
    "toggle" (fn [cell] (not cell))))

(defn get-brightness-fn [string]
  (case string
    "turn on" (fn [cell] (inc cell))
    "turn off" (fn [cell] (max (dec cell) 0))
    "toggle" (fn [cell] (+ cell 2))))

(defn update-lights [grid [update-fn [x1 y1 x2 y2]]]
  (vec (concat
         (subvec grid 0 x1)
         (mapv (fn [row]
                 (vec (concat (subvec row 0 y1)
                              (mapv update-fn (subvec row y1 (inc y2)))
                              (subvec row (inc y2) (count row)))))
               (subvec grid x1 (inc x2)))
         (subvec grid (inc x2) (count grid)))))

(def num-lit (->> (map (fn [[phrase nums]] [(get-lit-fn phrase) nums]) instructions)
                  (reduce update-lights (vec (repeat 1000 (vec (repeat 1000 false)))))
                  (map #(count (filter true? %)))
                  (reduce +)))

(def brightness (->> (map (fn [[phrase nums]] [(get-brightness-fn phrase) nums]) instructions)
                     (reduce update-lights (vec (repeat 1000 (vec (repeat 1000 0)))))
                     (map #(reduce + %))
                     (reduce +)))

(println (format "After following the instructions, %d lights are lit." num-lit))
(println (format "The total brightness is %d." brightness))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

