(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def instructions (map #(list (first %) (parse-long (subs % 1))) (str/split (str/trim (slurp "input.txt")) #", ")))

(defn get-direction [turn direction]
  (get-in { \L {:N :W :W :S :S :E :E :N} \R {:N :E :E :S :S :W :W :N} } [turn direction]))

(defn get-position [direction amplitude position]
  (case direction
    :N (assoc position :x (+ (get position :x) amplitude))
    :S (assoc position :x (- (get position :x) amplitude))
    :E (assoc position :y (+ (get position :y) amplitude))
    :W (assoc position :y (- (get position :y) amplitude))))

(defn follow-instructions [instructions stop-twice]
  (loop [direction :N position { :x 0 :y 0 } visited {} index 0]
    (if (or (>= index (count instructions)) (and stop-twice (get visited (hash position)))) position
      (let [[turn amplitude] (nth instructions index)
            next-direction (or (get-direction turn direction) direction)
            next-position (get-position next-direction amplitude position)]
        (recur next-direction next-position (assoc visited (hash position) true) (inc index))))))

(defn find-distance [{:keys [x y]}] (+ (abs x) (abs y)))

(def final-position (follow-instructions instructions false))

(def broken-instructions (apply concat
                                (map (fn [[direction amplitude]]
                                       (map (fn [i] (list (if (= i 0) direction) 1)) (range amplitude))) instructions)))

(def duplicate-position (follow-instructions broken-instructions true))

(println (format "Easter Bunny HQ is %d blocks away." (find-distance final-position)))
(println (format "The first location visited twice is %d blocks away." (find-distance duplicate-position)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

