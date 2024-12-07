(ns main
  (:require [clojure.string :as str]
            [clojure.set :as set]))

(def start-time (System/currentTimeMillis))

(def guard-map (mapv (comp vec seq) (str/split (slurp "input.txt") #"\n")))

(defn find-symbol [sym guard-map]
  (vec (filter some? (for [x (range (count (first guard-map))) y (range (count guard-map))]
                       (when (= sym (get-in guard-map [x y])) [x y])))))

(defn off-grid [[x y] guard-map]
  (or (< x 0) (< y 0) (>= x (count (first guard-map))) (>= y (count guard-map))))

(defn find-right [direction]
  (case direction \^ \> \> \v \v \< \< \^))

(defn find-next-position [[x y] direction]
  (case direction \^ [(dec x) y] \v [(inc x) y] \> [x (inc y)] \< [x (dec y)]))

(defn find-next-obstacle [[x y] direction obstacles]
  (let [obstacle (case direction
                   \^ (last (sort-by first (filter (fn [[i j]] (and (< i x) (= j y))) obstacles)))
                   \v (first (sort-by first (filter (fn [[i j]] (and (> i x) (= j y))) obstacles)))
                   \> (first (sort-by last (filter (fn [[i j]] (and (= i x) (> j y))) obstacles)))
                   \< (last (sort-by last (filter (fn [[i j]] (and (= i x) (< j y))) obstacles))))]
    (if (nil? obstacle) []
      (let [[i j] (identity obstacle)]
        (case direction
          \^ [\> [(inc i) j]]
          \v [\< [(dec i) j]]
          \> [\v [i (dec j)]]
          \< [\^ [i (inc j)]])))))

(defn find-loop [route init-direction obstacles loops]
  (let [proposed (last obstacles)]
    (if (some #(= % proposed) (drop-last route)) loops
      (loop [position (last route) direction init-direction steps 0]
        (let [[next-dir next-pos] (find-next-obstacle position direction obstacles)]
          (if (nil? next-pos) loops
            (if (> steps 200) (conj loops proposed)
              (recur next-pos next-dir (inc steps)))))))))

(defn track-guard [init-guard-map]
  (let [obstacles (find-symbol \# guard-map)]
    (loop [route [(first (find-symbol \^ guard-map))] direction \^ loops []]
      (let [curr-pos (last route)
            next-pos (find-next-position curr-pos direction)]
        (if (off-grid next-pos guard-map)
          [route loops]
          (if (= \# (get-in guard-map next-pos))
            (recur route (find-right direction) loops)
            (recur (conj route next-pos)
                   direction
                   (find-loop route (find-right direction) (conj obstacles next-pos) loops))))))))

(let [[route loops] (track-guard guard-map)]
  (println (format "The guard will visit %d distinct positions." (count (distinct route))))
  (println (format "An obstacle that would cause a loop could be placed in %d positions." (count (distinct loops)))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

