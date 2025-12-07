(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def diagram (->> (str/split-lines (slurp "input.txt"))
                  (map vec)))

(defn find-start-index [diagram]
  (first (first (filter #(= \S (second %)) (map-indexed vector (first diagram))))))

(defn count-splits [diagram]
  (reduce (fn [[total beams] line]
            [(+ total (count (filter #(= (nth line %) \^) beams)))
             (distinct (mapcat (fn [index] (if (= (nth line index) \^) [(dec index) (inc index)] [index])) beams))])
          [0 [(find-start-index diagram)]]
          (rest diagram)))

(defn get-timelines [diagram]
  (reduce (fn [timelines line]
            (vec (map-indexed (fn [i t]
                                (if (= \^ t)
                                  (+ (nth timelines (dec i)) (nth timelines (inc i)))
                                  (nth timelines i))) line)))
          (vec (repeat (count (first diagram)) 1))
          (reverse diagram)))

(println (format "The beam is split %d times." (first (count-splits diagram))))
(println (format "The particle ends up in %d timelines." (nth (get-timelines diagram) (find-start-index diagram))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

