(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def stone-map (->> (re-seq #"\d+" (slurp "input.txt"))
                    (map #(vec (list (parse-long %) 1)))
                    (into {})))

(defn split [stone]
  (let [string (str stone) midpoint (/ (count string) 2)]
    (mapv parse-long [(subs string 0 midpoint) (subs string midpoint)])))

(defn transform [stone]
  (if (zero? stone) [1]
    (if (even? (count (str stone))) (split stone)
      [(* stone 2024)])))

(defn blink [init-stone-map]
  (reduce (fn [stone-map [number freq]]
            (reduce (fn [stone-map transformed]
                      (assoc stone-map transformed (+ freq (get stone-map transformed 0))))
                    stone-map
                    (transform number)))
          {}
          init-stone-map))

(defn blink-n [stones n]
  (nth (iterate blink stones) n))

(defn count-stones [stone-map]
  (reduce + (vals stone-map)))

(println (format "After blinking 25 times, there are %d stones." (count-stones (blink-n stone-map 25))))
(println (format "After blinking 75 times, there are %d stones." (count-stones (blink-n stone-map 75))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

