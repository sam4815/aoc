(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def chunks (str/split (slurp "input.txt") #"\n\n"))
(def towels (str/split (first chunks) #", "))
(def designs (str/split (second chunks) #"\n"))

(defn find-matches [substring design]
  (->> (filter (fn [towel] (str/starts-with? design (str substring towel))) towels)
       (map (fn [towel] {:curr (str substring towel) :prev substring}))))

(defn count-permutations [design]
  (loop [queue (map (fn [t] {:curr t :prev ""}) towels) visited {"" 1}]
    (if (empty? queue) (get visited design 0)
      (let [{:keys [curr prev]} (first queue)]
        (if (get visited curr)
          (recur (rest queue) (assoc visited curr (+ (get visited curr) (get visited prev))))
          (recur (sort-by (comp count :curr) (concat (find-matches curr design) (rest queue)))
                 (assoc visited curr (get visited prev))))))))

(let [num-permutations (map count-permutations designs)]
  (println (format "%d designs are possible." (count (filter (complement zero?) num-permutations))))
  (println (format "The number of ways to make the designs is %d." (reduce + num-permutations))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

