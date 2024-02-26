(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def instructions (map (fn [line] (do {:op (re-find #"rect|row|column" line)
                                       :nums (map #(parse-long %) (re-seq #"\d+" line))}))
                       (str/split (slurp "input.txt") #"\n")))


(defn apply-instruction [pixels {:keys [op nums]}]
  (let [a (first nums) b (second nums)]
    (case op
      "rect" (map-indexed (fn [y row] (map-indexed (fn [x cell] (if (and (< x a) (< y b)) \# cell)) row)) pixels)
      "row" (map-indexed (fn [y row] (if (= y a) (concat (take-last b row) (take (- (count row) b) (drop-last b row))) row)) pixels)
      "column" (map-indexed (fn [y row] (map-indexed (fn [x cell] (if (= x a) (nth (nth pixels (mod (- y b) (count pixels))) a) cell)) row)) pixels))))

(def display (reduce apply-instruction (repeat 6 (repeat 50 \.)) instructions))

(def pixel-count (reduce + (map #(if (= \# %) 1 0) (apply concat display))))

(println (format "The number of pixels that should be lit is %d." pixel-count))
(println (format "The message that appears on the display is \n%s." (str/join "\n" (map str/join display))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

