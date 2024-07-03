(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def reindeer (->> (str/split (slurp "input.txt") #"\n")
                   (map #(identity [(first (re-seq #"[A-Z][a-z]+" %)) (map parse-long (re-seq #"\d+" %))]))
                   (into {})))

(def race-duration 2503)

(defn calculate-distance [total-duration [speed fly-duration rest-duration]]
  (let [elapsed-periods (quot total-duration (+ fly-duration rest-duration))
        remaining-time (mod total-duration (+ fly-duration rest-duration))]
    (if (>= remaining-time fly-duration)
      (* speed fly-duration (inc elapsed-periods))
      (+ (* speed fly-duration elapsed-periods) (* speed remaining-time)))))

(defn find-winners [reindeer seconds]
  (let [distances (update-vals reindeer (partial calculate-distance seconds))
        max-distance (val (apply max-key val distances))]
    (filter (fn [[k v]] (= max-distance v)) distances)))

(def winning-distance (val (first (find-winners reindeer race-duration))))

(def winning-points (->> (range 1 (inc race-duration))
                         (map (partial find-winners reindeer))
                         (reduce (fn [points winners]
                                   (reduce (fn [points winner] (update points (key winner) inc))
                                           points
                                           winners))
                                 (into {} (map #(identity [% 0]) (keys reindeer))))
                         (apply max-key val)
                         val))

(println (format "The winning reindeer has travelled %dkm." winning-distance))
(println (format "The winning reindeer has %d points." winning-points))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

