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
  (every? true? (map-indexed
                  (fn [i page]
                    (and (empty? (set/intersection (set (get after page [])) (set (subvec pages-update 0 i))))
                         (empty? (set/intersection (set (get before page [])) (set (subvec pages-update (inc i)))))))
                  pages-update)))

(defn find-middle [pages-update]
  (first (keep-indexed
           (fn [i page]
             (let [removed (concat (subvec pages-update 0 i) (subvec pages-update (inc i)))
                   a (set/intersection (set (get after page [])) (set removed))
                   b (set/intersection (set (get before page [])) (set removed))]
               (when (= (count a) (count b)) page)))
           pages-update)))

(def grouped (group-by is-ordered updates))

(println (format "The sum of the correctly-ordered middle pages is %d." (reduce + (map find-middle (get grouped true)))))
(println (format "The sum of the incorrectly-ordered middle pages is %d." (reduce + (map find-middle (get grouped false)))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

