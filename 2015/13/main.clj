(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def preferences (->> (str/split (slurp "input.txt") #"\n")
                      (map #(concat (re-seq #"[A-Z][a-z]+" %)
                                    (re-seq #"gain|lose" %)
                                    (map parse-long (re-seq #"\d+" %))))
                      (reduce (fn [all [a b effect magnitude]]
                                (assoc-in all [a b] (if (= effect "gain") magnitude (- magnitude)))) {})))

(def preferences-with-you (reduce (fn [all a] (assoc-in (assoc-in all ["You" a] 0) [a "You"] 0))
                                  preferences
                                  (keys preferences)))

(defn permutations [collection]
  (if (= (count collection) 1)
    (list collection)
    (for [head collection
          tail (permutations (disj (set collection) head))]
      (cons head tail))))

(defn find-happiness [arrangement preferences]
  (first
    (reduce (fn [[happiness a] b] [(+ happiness (get-in preferences [a b]) (get-in preferences [b a])) b])
            [0 (last arrangement)]
            arrangement)))

(defn find-optimal-happiness [preferences]
  (->> (permutations (set (keys preferences)))
       (map #(find-happiness % preferences))
       (apply max)))

(println (format "The happiness score of the optimal seating arrangement is %d." (find-optimal-happiness preferences)))
(println (format "Including you, the happiness score is %d." (find-optimal-happiness preferences-with-you)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

