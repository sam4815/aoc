(ns main
  (:require [clojure.math :as math]))

(def start-time (System/currentTimeMillis))

(def coordinates (partition 3 (map parse-long (re-seq #"\d+" (slurp "input.txt")))))

(defn find-distance [[x1 y1 z1] [x2 y2 z2]]
  (math/sqrt (+ (math/pow (- x1 x2) 2)
                (math/pow (- y1 y2) 2)
                (math/pow (- z1 z2) 2))))

(defn find-combos [coordinates]
  (->> (for [[i a] (map-indexed vector coordinates) b (subvec (vec coordinates) (inc i))]
         [(find-distance a b) a b])
       (sort-by first)))

(defn find-groups [n combos]
  (loop [i 0 sets []]
    (let [[_ a b] (nth combos i)
          a-set (some #(when (contains? % a) %) sets)
          b-set (some #(when (contains? % b) %) sets)]
      (if (or (= i n) (= (count coordinates) (reduce + (map count sets)))) [(nth combos (dec i)) sets]
        (recur (inc i)
               (if (and (nil? a-set) (nil? b-set))
                 (conj sets #{a b})
                 (if (and (nil? a-set) (some? b-set))
                   (conj (filter #(not= % b-set) sets) (conj b-set a))
                   (if (and (some? a-set) (nil? b-set))
                     (conj (filter #(not= % a-set) sets) (conj a-set b))
                     (if (not= a-set b-set)
                       (conj (filter #(and (not= % a-set) (not= % b-set)) sets) (into a-set b-set))
                       sets)))))))))

(defn find-group-product [groups]
  (reduce * (take 3 (sort > (map count groups)))))

(defn find-box-product [[_ a b]]
  (* (first a) (first b)))

(let [combos (find-combos coordinates)
      [_ groups] (find-groups 1000 combos)
      [combo _] (find-groups Integer/MAX_VALUE combos)]
  (println (format "The product of the circuits is %d." (find-group-product groups)))
  (println (format "The product of the last two boxes is %d." (find-box-product combo)))
  (println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000)))))

