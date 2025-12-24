(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def devices (->> (str/split-lines (slurp "input.txt"))
                  (mapv #(re-seq #"\w+" %))
                  (mapv #(do [(first %) (rest %)]))
                  (into {})))

(def inverted
  (->> (conj (keys devices) "out")
       (mapv #(do [% (keys (filter (fn [[k v]] (some #{%} v)) devices))]))
       (into {})))

(defn is-reachable [start end]
  (loop [queue [start] reachable {}]
    (if (empty? queue) (some #{end} (keys reachable))
      (let [node (first queue)]
        (if (contains? reachable node)
          (recur (rest queue) reachable)
          (recur (concat (get devices node []) (rest queue)) (assoc reachable node true)))))))

(defn count-paths [[start end]]
  (loop [queue (get inverted end) counts {end 1}]
    (if (empty? queue) (get counts start 0)
      (let [node (first (filter (fn [n] (every? #(or (get counts %) (not (is-reachable % end))) (get devices n))) queue))]
        (recur (distinct (concat (get inverted node) (filter #(not= node %) queue)))
               (assoc counts node (reduce + (map #(get counts % 0) (get devices node)))))))))

(def svr-routes [["svr" "fft" "dac" "out"] ["svr" "dac" "fft" "out"]])
(def num-svr-paths (reduce + (map (fn [route] (reduce * (map count-paths (partition 2 1 route)))) svr-routes))) 

(println (format "There are %d paths from you." (count-paths ["you" "out"])))
(println (format "There are %d paths from svr." num-svr-paths))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

