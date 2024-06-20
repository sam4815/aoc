(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(defn look-and-say [string]
  (->> (re-seq #"(\d)\1*" string)
       (map (fn [[match number]] (str (count match) number)))
       (apply str)))

(defn look-and-say-nth [n]
  (nth (iterate look-and-say (slurp "input.txt")) n))

(println (format "Playing look-and-say 40 times yields a result of length %d." (count (look-and-say-nth 40))))
(println (format "Playing look-and-say 50 times yields a result of length %d." (count (look-and-say-nth 50))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

