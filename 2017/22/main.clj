(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def init-states
  (->> (str/split (slurp "input.txt") #"\n")
       (map-indexed (fn [y row] (map-indexed #(vector (list %1 y) (if (= %2 \#) :infected :clean)) (seq row))))
       (apply concat)
       (into {})))

(def start-pos (let [mid-point (/ (dec (count (str/split (slurp "input.txt") #"\n"))) 2)] `(~mid-point ~mid-point)))

(defn turn-left [dir] (get { :up :left :left :down :down :right :right :up } dir))
(defn turn-right [dir] (get { :up :right :right :down :down :left :left :up } dir))
(defn reverse-dir [dir] (get { :up :down :down :up :left :right :right :left } dir))

(defn find-next-pos [[x y] dir]
  (case dir :up `(~x ~(dec y)) :down `(~x ~(inc y)) :left `(~(dec x) ~y) :right `(~(inc x) ~y)))

(defn find-next-dir [cell dir]
  (case cell :clean (turn-left dir) :weakened dir :infected (turn-right dir) :flagged (reverse-dir dir)))

(defn virus-transform [cell] (get { :clean :weakened :weakened :infected :infected :flagged :flagged :clean } cell))

(defn burst-v1 [n]
  (loop [states init-states pos start-pos dir :up num-infected 0 i 0]
    (if (= n i) num-infected
      (if (= (get states pos) :infected)
        (recur (dissoc states pos) (find-next-pos pos (turn-right dir)) (turn-right dir) num-infected (inc i))
        (recur (assoc states pos :infected) (find-next-pos pos (turn-left dir)) (turn-left dir) (inc num-infected) (inc i))))))

(defn burst-v2 [n]
  (loop [states init-states  pos start-pos dir :up num-infected 0 i 0]
    (if (= n i) num-infected
      (let [cell (or (get states pos) :clean)
            next-dir (find-next-dir cell dir)
            next-pos (find-next-pos pos next-dir)
            next-num-infected (if (= cell :weakened) (inc num-infected) num-infected)]
        (recur (assoc states pos (virus-transform cell)) next-pos next-dir next-num-infected (inc i))))))

(println (format "After 10000 bursts, there have been %d infections." (burst-v1 10000)))
(println (format "After 10000000 bursts, there have been %d infections." (burst-v2 10000000)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

