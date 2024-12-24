(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def connections-list (->> (str/split (slurp "input.txt") #"\n")
                           (mapv #(str/split % #"-"))))

(def connections-map (reduce (fn [a [b c]]
                               (assoc (assoc a c (assoc (get a c {}) b true)) b (assoc (get a b {}) c true)))
                             {} connections-list))

(defn find-triplets [connections]
  (loop [n 0 connected []]
    (if (= n (count connections)) (vec (distinct connected))
      (let [[a b] (get connections n)]
        (recur (inc n) (concat connected (->> (keys (get connections-map a))
                                              (filter (fn [c] (get (get connections-map b) c)))
                                              (map (fn [c] (sort [a b c]))))))))))

(defn grow [[a b c]]
  (loop [queue (keys (get connections-map a)) connected [a b c]]
    (if (empty? queue) (sort connected)
      (let [d (first queue)]
        (if (every? (fn [existing] (get (get connections-map existing) d)) connected)
          (recur (rest queue) (conj connected d))
          (recur (rest queue) connected))))))

(defn find-max [triplets]
  (loop [n 0 max-group []]
    (if (= n (count triplets)) max-group
      (let [grown (grow (get triplets n))]
        (recur (inc n) (if (> (count grown) (count max-group)) grown max-group))))))

(def triplets (find-triplets connections-list))
(def t-triplets (filter (fn [three] (some (fn [c] (= (first c) \t)) three)) triplets))

(println (format "%d sets of 3-connected computers contain a t." (count t-triplets)))
(println (format "The password to get into the LAN party is %s." (str/join "," (find-max triplets))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

