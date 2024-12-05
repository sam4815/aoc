(ns main
  (:require [clojure.string :as str]
            [clojure.set :as set]))

(def start-time (System/currentTimeMillis))

(def rules (->> (str/split (slurp "input.txt") #"\n\n")
                first
                (re-seq #"\d+")
                (map parse-long)
                (partition 2)))

(def after (reduce (fn [a [b c]] (assoc a b (conj (get a b []) c))) {} rules))
(def before (reduce (fn [a [b c]] (assoc a c (conj (get a c []) b))) {} rules))

(def updates (as-> (str/split (slurp "input.txt") #"\n\n") v
               (last v)
               (str/split v #"\n")
               (map #(mapv parse-long (re-seq #"\d+" %)) v)))

(defn is-ordered [pages-update]
  (loop [i 1]
    (if (= i (count pages-update))
      true
      (if (and
            (empty? (set/intersection (set (get after (nth pages-update i) [])) (set (subvec pages-update 0 i))))
            (empty? (set/intersection (set (get before (nth pages-update i) [])) (set (subvec pages-update (inc i))))))
        (recur (inc i))
        false))))

(defn sum-middle [pages-update]
  (->> (map-indexed (fn [i page]
                      (let [removed (concat (subvec pages-update 0 i) (subvec pages-update (inc i)))
                            a (set/intersection (set (get after page [])) (set removed))
                            b (set/intersection (set (get before page [])) (set removed))]
                        (if (= (count a) (count b)) page nil))) pages-update)
       (filter some?)
       (reduce +)))


(def grouped (group-by is-ordered updates))

(println (format "The sum of the correctly-ordered middle pages is %d." (reduce + (map sum-middle (get grouped true)))))
(println (format "The sum of the incorrectly-ordered middle pages is %d." (reduce + (map sum-middle (get grouped false)))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

