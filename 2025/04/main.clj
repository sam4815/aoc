(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def grid (->> (str/split (slurp "input.txt") #"\n")
               (mapv #(str/split % #""))))

(def adjacent [[-1 -1] [-1 0] [-1 1] [0 -1] [0 1] [1 -1] [1 0] [1 1]])

(defn count-removed [removed]
  (reduce + (map count (vals removed))))

(defn get-cell [grid removed [x y]]
  (let [is-removed (get (get removed x {}) y false)]
    (if is-removed "." (nth (nth grid y []) x "."))))

(defn find-coordinates [grid position]
  (let [width (count (first grid))]
    [(rem position width) (quot position width)]))

(defn is-accessible? [grid removed coordinates]
  (let [neighbours (map (partial get-cell grid removed) (map #(map + % coordinates) adjacent))]
    (> 4 (count (filter #{"@"} neighbours)))))

(defn find-accessible [grid removed]
  (loop [accessible {} position 0]
    (if (= position (* (count grid) (count (first grid))))
      accessible
      (let [coordinates (find-coordinates grid position)
            is-accessible (and (= "@" (get-cell grid removed coordinates)) (is-accessible? grid removed coordinates))]
        (recur (if is-accessible (assoc-in accessible coordinates true) accessible) (inc position))))))

(defn find-all-accessible [grid]
  (loop [all-removed {}]
    (let [removed (find-accessible grid all-removed)]
      (if (zero? (count-removed removed)) all-removed (recur (merge-with into all-removed removed))))))

(println (format "Initally, %d rolls of paper can be accessed by a forklift." (count-removed (find-accessible grid {}))))
(println (format "In total, %d rolls of paper can be accessed by a forklift." (count-removed (find-all-accessible grid))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

