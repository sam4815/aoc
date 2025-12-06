(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(defn transpose [matrix]
  (apply mapv vector matrix))

(def lines (str/split (slurp "input.txt") #"\n"))
(def equations (transpose (map #(str/split (str/trim %) #"\s+") lines)))

(defn read-number [numbers index]
  (->> (drop-last numbers) 
       (map (fn [line] (str/trim (str (nth line index \space)))))
       str/join
       parse-long))

(defn solve [equation]
  (reduce (case (last equation) "+" + "*" *) (map parse-long (drop-last equation))))

(defn solve-cephalopod [numbers]
  (reduce (fn [[total acc acc-sym] [index sym]]
            (let [number (read-number numbers index)]
              (case sym
                \space [total (if (nil? number) acc (acc-sym acc number)) acc-sym]
                \+ [(+ total acc) number +]
                \* [(+ total acc) number *])))
          [0 0 +]
          (map-indexed vector (str (last numbers) "+"))))

(println (format "The sum of the problems is %d." (reduce + (map solve equations))))
(println (format "The true sum of the problems is %d." (first (solve-cephalopod lines))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

