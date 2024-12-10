(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def topology (->> (str/split (slurp "input.txt") #"\n")
                   (mapv #(mapv parse-long (re-seq #"\d" %)))))

(defn find-zeros [topology]
  (filter some? (for [x (range (count topology)) y (range (count (first topology)))]
                  (when (= (get-in topology [x y]) 0) [x y]))))

(defn find-next [[x y]]
  (->> (map (fn [[i j]] [(+ x i) (+ y j)]) [[-1 0] [1 0] [0 -1] [0 1]])
       (filter (fn [[i j]] (= (inc (get-in topology [x y])) (get-in topology [i j]))))))

(defn find-nines [zero]
  (loop [queue [zero] nines []]
    (if (empty? queue) nines
      (let [position (first queue)]
        (if (= (get-in topology position) 9)
          (recur (rest queue) (conj nines position))
          (recur (concat (find-next position) (rest queue)) nines))))))

(defn score-trailheads [zeros]
  (reduce + (map (comp count distinct find-nines) zeros)))

(defn rate-trailheads [zeros]
  (reduce + (map (comp count find-nines) zeros)))

(println (format "The sum of the trailhead scores is %d." (score-trailheads (find-zeros topology))))
(println (format "The sum of the trailhead ratings is %d." (rate-trailheads (find-zeros topology))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

