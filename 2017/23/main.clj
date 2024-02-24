(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def instructions
  (->> (str/split (slurp "input.txt") #"\n")
       (map #(str/split % #"\s"))
       (map #(list (first %) (or (parse-long (second %)) (second %)) (or (parse-long (last %)) (last %))))))

(defn update-registers [[opcode a b] registers]
  (case opcode
    "set" (assoc registers a b)
    "sub" (assoc registers a (- (or (get registers a) 0) b))
    "mul" (assoc (assoc registers :mul (inc (get registers :mul))) a (* (or (get registers a) 0) b))))

(defn process-instruction [[opcode a b-sym] {:keys [reg pc]}]
  (let [b (if (integer? b-sym) b-sym (get reg b-sym))]
    (case opcode
      "jnz" (if (not= (or (get reg a) a) 0) (list reg (+ pc b)) (list reg (inc pc)))
      (list (update-registers (list opcode a b) reg) (inc pc)))))

(defn process-instructions [instructions init-program]
  (loop [program init-program n 0]
    (if (or (>= (get program :pc) (count instructions)) (>= n 100000)) program
      (let [[next-reg next-pc] (process-instruction (nth instructions (get program :pc)) program)]
        (recur { :reg next-reg :pc next-pc } (inc n))))))

(defn is-divisible [target]
  (loop [n 2] (if (= target n) false (if (= (mod target n) 0) true (recur (inc n))))))

(def debug-program (process-instructions instructions { :reg { "a" 0 :mul 0 } :pc 0 }))
(def fiery-program (process-instructions instructions { :reg { "a" 1 :mul 0 } :pc 0 }))

(let [start (get-in fiery-program [:reg "b"]) end (inc (get-in fiery-program [:reg "c" ]))]
  (def h-value (count (filter is-divisible (range start end 17)))))

(println (format "In debug mode, the mul instruction is performed %d times." (get-in debug-program [:reg :mul])))
(println (format "After the unoptimized program finishes, the value in register h is %d." h-value))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

