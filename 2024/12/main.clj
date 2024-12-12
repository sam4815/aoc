(ns main
  (:require [clojure.string :as str]
            [clojure.set :as set]))

(def start-time (System/currentTimeMillis))

(def farm (->> (str/split (slurp "input.txt") #"\n")
               (mapv (comp vec seq))))

(defn find-adjacent [[x y]]
  (mapv (fn [[i j]] [(+ x i) (+ y j)]) [[-1 0] [1 0] [0 -1] [0 1]]))

(defn find-adjacent-unvisited [[x y] visited]
  (->> (find-adjacent [x y])
       (filter (fn [adj] (and (not (get-in visited adj))
                              (= (get-in farm adj) (get-in farm [x y])))))))

(defn count-perimeter [[x y]]
  (->> (find-adjacent [x y])
       (mapv (fn [[i j]] (get-in farm [i j])))
       (filter (fn [plant] (not= plant (get-in farm [x y]))))
       count))

(defn find-gaps [line]
  (reduce (fn [[gaps [i j]] [x y]]
            (if (= (inc j) y)
              [gaps [x y]]
              [{:open (conj (get gaps :open) y) :close (conj (get gaps :close) j)} [x y]]))
          [{:open #{(second (first line))} :close #{(second (last line))}} (first line)]
          (rest line)))

(defn count-change-in-gaps [area]
  (loop [rows (partition-by first area) num-sides 0 gaps-above {:open #{} :close #{}}]
    (if (empty? rows) num-sides
      (let [row (first rows) [gaps-below _] (find-gaps row)]
        (recur (rest rows)
               (+ num-sides
                  (count (set/difference (get gaps-below :open) (get gaps-above :open)))
                  (count (set/difference (get gaps-below :close) (get gaps-above :close))))
               gaps-below)))))

(defn count-sides [area]
  (+ (count-change-in-gaps (sort area))
     (count-change-in-gaps (sort (mapv #(into [] (reverse %)) area)))))

(defn get-region [position init-visited regions]
  (loop [queue [position] area [] visited init-visited]
    (if (empty? queue)
      [(conj regions {:area (count area)
                      :perimeter (reduce + (map count-perimeter area))
                      :sides (count-sides area)}) visited]
      (let [curr (first queue)]
        (if (get-in visited curr)
          (recur (rest queue) area visited)
          (recur (concat (find-adjacent-unvisited curr visited) (rest queue))
                 (conj area curr)
                 (assoc-in visited curr true)))))))

(defn get-regions [farm]
  (->> (for [x (range (count farm)) y (range (count (first farm)))] [x y])
       (reduce (fn [[regions visited] position]
                 (if (get-in visited position)
                   [regions visited]
                   (get-region position visited regions)))
               [[] {}])))

(defn find-cost [regions property]
  (reduce + (map #(* (get % :area) (get % property)) regions)))

(let [[regions _] (get-regions farm)]
  (println (format "The total price of fencing is %d." (find-cost regions :perimeter)))
  (println (format "The new price of fencing is %d." (find-cost regions :sides))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

