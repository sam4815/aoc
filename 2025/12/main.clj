(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def chunks (str/split (slurp "input.txt") #"\n\n"))

(def shapes (->> (drop-last chunks)
                 (map #(rest (str/split-lines %)))
                 (map (fn [chunk] (for [i (range 3) j (range 3)] (if (= \# (nth (nth chunk i) j)) [i j] nil))))
                 (map (partial remove nil?))))

(def regions (->> (last chunks)
                  str/split-lines
                  (map (fn [region] (map parse-long (re-seq #"\d+" region))))))

(defn could-fit [[x y & indices]]
  (let [area (* x y)
        space-required (map-indexed (fn [i n] (* (count (nth shapes i)) n)) indices)]
    (>= area (reduce + space-required))))

(def valid-regions (filter could-fit regions))

(println (format "The number of regions that can fit all the presents is %d." (count valid-regions)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

