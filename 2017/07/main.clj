(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(defn parse-program [program-string]
  (let [program-names (re-seq #"[a-z]+" program-string)]
    { :weight (parse-long (first (re-seq #"\d+" program-string)))
     :name (first program-names)
     :balancing (rest program-names) }))

(def programs (into {} (map (juxt :name identity) (mapv #(parse-program %) (str/split (slurp "input.txt") #"\n")))))

(defn find-bottom [programs]
  (let [supporting (set (mapv #(get % :name) (vals programs)))
        supported (set (flatten (mapv #(get % :balancing) (vals programs))))]
    (first (filter #(not (contains? supported %)) supporting))))

(defn find-weight [program-name]
  (let [program (get programs program-name)]
    (+ (get program :weight) (reduce + 0 (map find-weight (get program :balancing))))))

(defn find-unbalanced [program]
  (let [balancing-weights (map #(hash-map :name % :weight (find-weight %)) (get program :balancing))
        unbalanced-weight (first (first (filter #(= (count %) 1) (vals (group-by :weight balancing-weights)))))
        unbalanced-program (get programs (get unbalanced-weight :name))]
    (if (some? unbalanced-program)
      (or (find-unbalanced unbalanced-program)
          (hash-map :program unbalanced-program :difference (reduce - (keys (frequencies (map :weight balancing-weights)))))))))

(def bottom (get programs (find-bottom programs)))
(def unbalanced (find-unbalanced bottom))
(def desired-weight (- (get-in unbalanced [:program :weight]) (get unbalanced :difference)))

(println (format "The name of the program at the bottom of the tower is %s." (get bottom :name)))
(println (format "To balance the tower, the program with the wrong weight needs to have weight %d." desired-weight))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

