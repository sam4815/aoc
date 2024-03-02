(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def num-elves (parse-long (str/trim (slurp "input.txt"))))

(defn find-next [index elves]
  (loop [n 1]
    (if (> n (dec (count elves))) nil
      (if (some? (nth elves (mod (+ index n) (count elves))))
        (mod (+ index n) (count elves))
        (recur (inc n))))))

(defn remove-left [size]
  (loop [n 0 elves (mapv #(inc %) (range size)) left-n (find-next n elves)]
    (if (nil? left-n) (first (filter some? elves))
      (let [next-elves (assoc elves left-n nil) next-n (find-next left-n next-elves)]
        (recur next-n next-elves (find-next next-n next-elves))))))

(defn remove-opposite [size]
  (loop [n 0 elves (mapv #(inc %) (range size)) opposite-n (quot size 2) num-elves size]
    (if (nil? n) (first (filter some? elves))
      (let [next-elves (assoc elves opposite-n nil)
            next-n (find-next n next-elves)
            next-opposite (find-next opposite-n next-elves)]
        (recur next-n next-elves (if (even? num-elves) next-opposite (find-next next-opposite next-elves)) (dec num-elves))))))

(println (format "Under the first set of rules, the Elf that gets all the presents is Elf %d." (remove-left num-elves)))
(println (format "Under the second set of rules, the Elf that gets all the presents is Elf %d." (remove-opposite num-elves)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

