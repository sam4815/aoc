(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def word-search (->> (str/split (slurp "input.txt") #"\n")
                      (mapv #(str/split % #""))))

(def xmas [(map (fn [move] (map #(map (partial * %) move) (range 4)))
                [[0 1] [1 1] [1 0] [1 -1] [0 -1] [-1 -1] [-1 0] [-1 1]])
           (fn [words] (count (filter #{"XMAS"} words)))])

(def x-mas [[[[1 1] [0 0] [-1 -1]] [[-1 1] [0 0] [1 -1]]]
            (fn [words] (quot (count (filter #(or (= % "MAS") (= % "SAM")) words)) 2))])

(defn find-coordinates [grid position]
  (let [width (count (first grid))]
    [(rem position width) (quot position width)]))

(defn get-letter [grid [x y]]
  (nth (nth grid y []) x ""))

(defn get-words-from-positions [grid positions]
  (map (fn [position] (str/join (map (partial get-letter grid) position))) positions))

(defn count-from-position [grid position [moves count-fn]]
  (let [[x y] (find-coordinates grid position)
        positions (map (fn [move] (map #(map + % [x y]) move)) moves)
        words (get-words-from-positions grid positions)]
    (count-fn words)))

(defn find-word [grid rubric]
  (loop [times-found 0 position 0]
    (if (= position (* (count grid) (count (first grid))))
      times-found
      (recur (+ times-found (count-from-position grid position rubric))
             (inc position)))))

(println (format "XMAS appears %d times." (find-word word-search xmas)))
(println (format "X-MAS appears %d times." (find-word word-search x-mas)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

