(ns main
  (:require [clojure.string :as str]
            [clojure.math :as math]))

(def start-time (System/currentTimeMillis))

(def numbers (map parse-long (re-seq #"\d+" (slurp "input.txt"))))
(def program (drop 3 numbers))

(defn get-combo [operand state]
  (case operand 4 (get state :a) 5 (get state :b) 6 (get state :c) operand))

(defn perform-instruction [[opcode operand] state]
  (as-> (case opcode
          0 (assoc state :a (math/floor-div (get state :a) (math/pow 2 (get-combo operand state)))) ;adv
          1 (assoc state :b (bit-xor (get state :b) operand)) ;bxl
          2 (assoc state :b (mod (get-combo operand state) 8)) ;bst
          3 (assoc state :pc (if (zero? (get state :a)) (get state :pc) (- operand 2))) ;jnz
          4 (assoc state :b (bit-xor (get state :b) (get state :c))) ;bxc
          5 (assoc state :out (conj (get state :out []) (mod (get-combo operand state) 8))) ;out
          6 (assoc state :b (math/floor-div (get state :a) (math/pow 2 (get-combo operand state)))) ;bdv
          7 (assoc state :c (math/floor-div (get state :a) (math/pow 2 (get-combo operand state))))) ;cdv
    s (assoc s :pc (+ (get s :pc) 2))))

(defn run [program a]
  (loop [state {:pc 0 :a a :b 0 :c 0}]
    (if (>= (get state :pc) (count program)) (get state :out)
      (recur (perform-instruction (take 2 (drop (get state :pc) program)) state)))))

(defn find-a [program]
  (loop [a 0]
    (let [output (run program a)]
      (if (= output program) a
        (if (= output (take-last (count output) program))
          (recur (* a 8))
          (recur (inc a)))))))

(println (format "After the program halts, its output is %s." (str/join "," (run program (first numbers)))))
(println (format "The program outputs itself if a is initialized to %d." (find-a program)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

