(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def particles (map #(map parse-long (re-seq #"[\d|-]+" %)) (str/split (slurp "input.txt") #"\n")))

(defn tick [[px py pz vx vy vz ax ay az]]
  (let [vvx (+ vx ax) vvy (+ vy ay) vvz (+ vz az)]
    (list (+ px vvx) (+ py vvy) (+ pz vvz) vvx vvy vvz ax ay az)))

(defn remove-duplicates [particles]
  (let [position-frequency (frequencies (map #(hash (take 3 %)) particles))]
    (filter #(= (get position-frequency (hash (take 3 %))) 1) particles)))

(defn get-manhattan-distance [[px py pz]]
  (+ (abs px) (abs py) (abs pz)))

(defn simulate [n init-particles]
  (reduce (fn [particles _] (map tick particles)) init-particles (range n)))

(defn simulate-with-collisions [n init-particles]
  (reduce (fn [particles _] (map tick (remove-duplicates particles))) init-particles (range n)))

(defn find-closest-to-origin [particles]
  (apply min-key #(get-manhattan-distance (second %)) (map-indexed #(list %1 %2) particles)))

(def closest-to-origin (find-closest-to-origin (simulate 500 particles)))
(def num-particles (count (simulate-with-collisions 500 particles)))

(println (format "In the long term, the particle closest to the origin is %s." (first closest-to-origin)))
(println (format "After all collisions are resolved, there are %d particles." num-particles))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

