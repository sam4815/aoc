(ns main
  (:require [clojure.string :as str]
            [clojure.set :as set]))

(def start-time (System/currentTimeMillis))

(def memory-size 70)

(def corrupted-bytes (->> (re-seq #"\d+" (slurp "input.txt"))
                          (map parse-long)
                          (partition 2)))

(defn off-grid [[x y]]
  (or (< x 0) (< y 0) (> x memory-size) (> y memory-size)))

(defn find-next-positions [[x y] cost obstacles]
  (->> [[(dec x) y] [(inc x) y] [x (inc y)] [x (dec y)]]
       (filter (fn [pos] (and (not (off-grid pos)) (not (get obstacles pos)))))
       (map (fn [pos] {:pos pos :cost (inc cost)}))))

(defn find-min-path [corrupted-bytes]
  (let [obstacles (into {} (map (fn [pos] [pos true]) corrupted-bytes))]
    (loop [queue [{:pos [0 0] :cost 0}] visited {}]
      (if (empty? queue) nil
        (let [{:keys [pos cost]} (first queue)]
          (if (= pos [memory-size memory-size]) cost
            (if (get visited pos)
              (recur (rest queue) visited)
              (recur (sort-by :cost (concat (find-next-positions pos cost obstacles) (rest queue)))
                     (assoc visited pos true)))))))))

(defn find-blocked [corrupted-bytes]
  (loop [n 0 jump (count corrupted-bytes)]
    (let [min-path (find-min-path (take n corrupted-bytes))]
      (if (= jump 0)
        (nth corrupted-bytes (if (some? min-path) n (dec n)))
        (recur (if (some? min-path) (+ n jump) (- n jump)) (quot jump 2))))))

(println (format "The shortest path to the exit takes %d steps." (find-min-path (take 1024 corrupted-bytes))))
(println (format "The first byte to block the exit is at %s." (str/join "," (find-blocked corrupted-bytes))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

