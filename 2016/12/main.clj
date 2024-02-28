(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def instructions
  (->> (str/split (slurp "input.txt") #"\n")
       (map #(str/split % #"\s"))
       (map #(list (first %) (or (parse-long (second %)) (second %)) (or (parse-long (last %)) (last %))))))

(defn update-registers [[opcode x y] registers]
  (case opcode
    "cpy" (assoc registers y (if (int? x) x (get registers x 0)))
    "inc" (assoc registers x (inc (get registers x 0)))
    "dec" (assoc registers x (dec (get registers x 0)))))

(defn process-instruction [[opcode x y] {:keys [reg pc]}]
    (case opcode
      "jnz" (if (not= (if (int? x) x (get reg x 0)) 0) (list reg (+ pc y)) (list reg (inc pc)))
      (list (update-registers (list opcode x y) reg) (inc pc))))

(defn process-instructions [instructions init-program]
  (loop [program init-program n 0]
    (if (>= (get program :pc) (count instructions)) program
      (let [[next-reg next-pc] (process-instruction (nth instructions (get program :pc)) program)]
        (recur { :reg next-reg :pc next-pc } (inc n))))))

(def program (process-instructions instructions { :reg {} :pc 0 }))
(def initialized-program (process-instructions instructions { :reg { "c" 1 } :pc 0 }))

(println (format "After executing the assembunny code, the value in register a is %d." (get-in program [:reg "a"])))
(println (format "With register c initialized, the value in register a is %d." (get-in initialized-program [:reg "a"])))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

