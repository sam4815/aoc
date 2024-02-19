(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def tubes (map #(seq %) (str/split (slurp "input.txt") #"\n")))

(defn find-start [tubes]
  `(~(first (filter some? (map-indexed #(if (= %2 \|) %1 nil) (first tubes)))) 0))

(defn opposite [direction] (get { :north :south :south :north :east :west :west :east } direction))

(defn get-cell [[x y] tubes] (nth (nth tubes y) x))

(defn is-valid-position [[x y]] (and (>= x 0) (>= y 0) (< x (count (first tubes))) (< y (count tubes))))

(defn add-to-path [cell path] (if (Character/isLetter cell) (concat path [cell]) path)) 

(defn find-valid-steps [[x y] direction tubes]
  (->> [[:north `(~x ~(dec y))] [:east `(~(inc x) ~y)] [:south `(~x ~(inc y))] [:west `(~(dec x) ~y)]]
       (filter (fn [[dir pos]] (and (is-valid-position pos)
                                    (not= dir (opposite direction))
                                    (not= \space (get-cell pos tubes)))))))

(defn find-next-step [pos dir tubes]
  (let [valid-steps (find-valid-steps pos dir tubes)]
    (if (> (count valid-steps) 1) (first (filter #(= dir (first %)) valid-steps)) (first valid-steps))))

(defn follow-path [tubes]
  (loop [position (find-start tubes) direction :south path [] num-steps 1]
    (let [[next-dir next-pos] (find-next-step position direction tubes)]
      (if (nil? next-pos) `(~path ~num-steps)
        (recur next-pos next-dir (add-to-path (get-cell next-pos tubes) path) (inc num-steps))))))

(let [[path num-steps] (follow-path tubes)]
  (println (format "Following the path, the packet will see the letters %s." (str/join path)))
  (println (format "The packet needs to take %d steps." num-steps)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

