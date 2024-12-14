(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def width 101)
(def height 103)

(def robots (->> (str/split (slurp "input.txt") #"\n")
                 (mapv #(mapv parse-long (re-seq #"-?\d+" %)))))

(defn elapse [n [x y dx dy]]
  [(mod (+ x (* n dx)) width) (mod (+ y (* n dy)) height)])

(defn get-quadrant [[x y]]
  (let [x-mid (/ (dec width) 2) y-mid (/ (dec height) 2)]
    (cond (and (< x x-mid) (< y y-mid)) 1
          (and (< x x-mid) (> y y-mid)) 2
          (and (> x x-mid) (< y y-mid)) 3
          (and (> x x-mid) (> y y-mid)) 4)))

(defn get-safety-factor [robots n]
  (->> (map (partial elapse n) robots)
       (map get-quadrant)
       (filter some?)
       frequencies
       vals
       (reduce *)))

(defn stringify [robots]
  (->> (reduce (fn [grid [x y]] (assoc-in grid [y x] \█))
               (vec (repeat height (vec (repeat width \ )))) robots)
       (map str/join)
       (str/join "\n")))

(defn find-least-safety [robots]
  (apply min-key (partial get-safety-factor robots) (range (* height width))))

(def timestamp (find-least-safety robots))

(println (format "After 100 seconds, the safety factor is %d." (get-safety-factor robots 100)))
(println (format "The Easter egg is displayed after %d seconds:\n%s" timestamp (stringify (map (partial elapse timestamp) robots))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

