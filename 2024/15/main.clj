(ns main
  (:require [clojure.string :as str]
            [clojure.set :as set]))

(def start-time (System/currentTimeMillis))

(def parts (str/split (slurp "input.txt") #"\n\n"))
(def warehouse (mapv (comp vec seq) (str/split (first parts) #"\n")))
(def instructions (re-seq #"\S" (second parts)))

(defn find-symbol [sym warehouse]
  (vec (filter some? (for [x (range (count (first warehouse))) y (range (count warehouse))]
                       (when (= sym (get-in warehouse [x y])) [x y])))))

(defn to-map [positions]
  (into {} (map (fn [pos] (vec (list pos true))) positions)))

(defn find-next-position [[x y] direction]
  (case direction "^" [(dec x) y] "v" [(inc x) y] ">" [x (inc y)] "<" [x (dec y)]))

(defn find-next-positions [positions direction]
  (conj (reduce (fn [a [x y]]
                  (concat a [(find-next-position [x y] direction) (find-next-position [x (inc y)] direction)]))
                  []
                  (rest positions))
                 (find-next-position (first positions) direction)))

(defn move-boxes [box-positions direction boxes]
  (reduce (fn [b pos] (assoc (dissoc b pos) (find-next-position pos direction) true))
          boxes
          (reverse box-positions)))

(defn find-boxes [positions boxes under-move]
  (->> (filter (fn [[x y]] (or (some #(= [x y] %) positions) (some #(= [x (inc y)] %) positions))) (keys boxes))
       (filter (fn [pos] (not (some #(= pos %) under-move))))))

(defn move [under-move direction boxes walls]
  (let [next-pos (find-next-position (last under-move) direction)]
    (if (get boxes next-pos)
      (move (conj under-move next-pos) direction boxes walls)
      (if (get walls next-pos)
        [(first under-move) boxes]
        [(find-next-position (first under-move) direction)
         (move-boxes (rest under-move) direction boxes)]))))

(defn move-dbl [under-move direction boxes walls]
  (let [next-positions (find-next-positions under-move direction)]
    (if (some (fn [pos] (get walls pos)) next-positions)
      [(first under-move) boxes]
      (let [new-boxes (find-boxes next-positions boxes under-move)]
        (if (not (empty? new-boxes))
          (move-dbl (concat under-move new-boxes) direction boxes walls)
          [(find-next-position (first under-move) direction)
           (move-boxes (rest under-move) direction boxes)])))))

(defn track-robot [warehouse instructions]
  (let [walls (to-map (find-symbol \# warehouse))]
    (loop [position (first (find-symbol \@ warehouse))
           moves instructions
           boxes (to-map (find-symbol \O warehouse))]
      (if (empty? moves) (keys boxes)
        (let [[next-pos next-boxes] (move [position] (first moves) boxes walls)]
          (recur next-pos (rest moves) next-boxes))))))

(defn track-dbl-robot [warehouse instructions]
  (let [walls (to-map (reduce (fn [a [x y]] (concat a [[x (* y 2)] [x (inc (* y 2))]])) [] (find-symbol \# warehouse)))
        [x y] (first (find-symbol \@ warehouse))]
    (loop [position [x (* y 2)]
           moves instructions
           boxes (to-map (map (fn [[x y]] [x (* y 2)]) (find-symbol \O warehouse)))]
      (if (empty? moves) (keys boxes)
        (let [[next-pos next-boxes] (move-dbl [position] (first moves) boxes walls)]
          (recur next-pos (rest moves) next-boxes))))))

(defn count-gps [box-positions]
  (reduce + (map (fn [[x y]] (+ (* 100 x) y)) box-positions)))

(println (format "In the first warehouse, the GPS coordinates sum to %d." (count-gps (track-robot warehouse instructions))))
(println (format "In the second warehouse, the GPS coordinates sum to %d." (count-gps (track-dbl-robot warehouse instructions))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

