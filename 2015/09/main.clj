(ns main
  (:require [clojure.string :as str]
            [clojure.set :as set]))

(def start-time (System/currentTimeMillis))

(defn add-connection [connections [start end distance]]
  (-> connections
      (assoc-in [start end] distance)
      (assoc-in [end start] distance)))

(def connections (->> (str/split (slurp "input.txt") #"\n")
                      (map #(concat (re-seq #"[A-Z]\w+" %) (map parse-long (re-seq #"\d+" %))))
                      (reduce add-connection {})))

(defn find-best-distance [connections best-fn]
  (loop [queue (map (fn [k] {:location k :distance 0 :visited [k]}) (keys connections)) best-distance nil]
    (let [{:keys [location distance visited]} (first queue)]
      (if (empty? queue) best-distance

        (if (= (count visited) (count connections))
          (recur (rest queue) (if (nil? best-distance) distance (best-fn distance best-distance)))

          (let [routes (map (fn [k] {:location k
                                     :distance (+ distance (get-in connections [location k]))
                                     :visited (conj visited k)})
                            (set/difference (set (keys connections)) (set visited)))]
            (recur (concat routes (rest queue)) best-distance)))))))

(println (format "The distance of the shortest route is %d." (find-best-distance connections min)))
(println (format "The distance of the longest route is %d." (find-best-distance connections max)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

