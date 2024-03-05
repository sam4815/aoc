(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def instructions
  (->> (str/split (slurp "input.txt") #"\n")
       (map #(str/split % #"\s"))
       (mapv #(list (first %) (or (parse-long (second %)) (second %)) (or (parse-long (last %)) (last %))))))

(defn update-registers [[opcode x y] registers]
  (case opcode
    "out" (assoc registers :out (concat (get registers :out []) [(if (int? x) x (get registers x 0))]))
    "cpy" (assoc registers y (if (int? x) x (get registers x 0)))
    "inc" (assoc registers x (inc (get registers x 0)))
    "dec" (assoc registers x (dec (get registers x 0)))))

(defn process-instruction [[opcode x y] {:keys [reg pc]}]
  (case opcode
    "jnz" (if (not= (if (int? x) x (get reg x 0)) 0) [reg (+ pc (if (int? y) y (get reg y 0)))] [reg (inc pc)])
    [(update-registers [opcode x y] reg) (inc pc)]))

(defn process-instructions [init-instructions init-program]
  (loop [{:keys [reg pc]} init-program instructions init-instructions n 0]
    (if (or (>= pc (count instructions)) (> n 100000)) {:pc pc :reg reg}
      (let [[next-reg next-pc] (process-instruction (nth instructions pc) {:reg reg :pc pc})]
        (recur {:reg next-reg :pc next-pc} instructions (inc n))))))

(defn find-clock-signal [instructions]
  (loop [a-value 0]
    (let [result (process-instructions instructions {:pc 0 :reg {"a" a-value}})]
      (if (= (take 10 (get-in result [:reg :out])) '(0 1 0 1 0 1 0 1 0 1)) a-value
        (recur (inc a-value))))))

(println (format "The smallest positive integer that outputs a clock signal is %d." (find-clock-signal instructions)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

