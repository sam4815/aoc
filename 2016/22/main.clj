(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(defn parse-node [string]
  (let [[x y size used avail] (map #(parse-long %) (re-seq #"\d+" string))]
    {:position [x y] :size size :used used :avail avail}))

(defn is-viable-pair [a b]
  (and (not (zero? (get a :used)))
       (not= (get a :position) (get b :position))
       (<= (get a :used) (get b :avail))))

(def nodes
  (let [parsed-nodes (->> (str/split (slurp "input.txt") #"\n")
                          (map parse-node)
                          (filter #(some? (get % :size))))]
    (->> parsed-nodes
         (map (fn [node] [(get node :position) (assoc node :viable (filter #(is-viable-pair node %) parsed-nodes))]))
         (filter (fn [[node pairs]] (not-empty pairs)))
         (into {}))))

(defn find-target [nodes] [(first (first (sort-by first > (keys nodes)))) 0])

(defn find-open [nodes] (get (first (get (val (first nodes)) :viable)) :position))

(defn get-adjacent [[x y] nodes]
  (filter (fn [position] (some? (get nodes position))) [[(inc x) y] [(dec x) y] [x (inc y)] [x (dec y)]]))

(defn find-moves [open target steps nodes]
  (let [open-size (get (get nodes open) :size)]
    (->> (get-adjacent open nodes)
         (filter (fn [position] (<= (get (get nodes position) :used) open-size)))
         (map (fn [position] (if (= position target)
                               {:target open :open position :steps (inc steps)}
                               {:target target :open position :steps (inc steps)}))))))

(defn find-shortest-path [nodes]
  (loop [queue [{:target (find-target nodes) :open (find-open nodes) :steps 0}] visited {}]
    (let [{:keys [open target steps]} (first queue)]
      (if (= target [0 0]) steps
        (if (or (some? (get visited (concat open target))) (> (second target) 2))
          (recur (rest queue) visited)            
          (recur (sort-by :steps (concat (find-moves open target steps nodes) (rest queue)))
                 (assoc visited (concat open target) steps)))))))

(println (format "The number of viable pairs is %d." (reduce + (map #(count (get % :viable)) (vals nodes)))))
(println (format "The fewest number of steps required to reach the goal data is %d." (find-shortest-path nodes)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

