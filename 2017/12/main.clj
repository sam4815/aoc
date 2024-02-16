(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def program-connections (into {} (mapv #(let [nums (re-seq #"\d+" %)] [(first nums) (rest nums)]) (str/split (slurp "input.txt") #"\n"))))

(defn find-connected [group connections]
  (loop [members (set ()) connected (get connections group)]
    (if (= (count connected) 0)
      members
      (recur (set (concat members connected))
             (filter #(not (contains? members %)) (flatten (map #(get connections %) connected)))))))

(def zero-group-size (count (set (find-connected "0" program-connections))))
(def group-count (count (set (map #(hash (find-connected % program-connections)) (keys program-connections)))))

(println (format "The number of programs in the group that contains 0 is %d." zero-group-size))
(println (format "There are %d groups in total." group-count))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

