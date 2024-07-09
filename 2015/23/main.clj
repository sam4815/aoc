(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def instructions (->> (str/split (slurp "input.txt") #"\n")
                       (map (fn [line] (do {:opcode (subs line 0 3)
                                            :register (re-find #"a|b" line)
                                            :value (parse-long (or (re-find #"-?\d+" line) ""))})))))

(defn run [instructions init-registers]
  (loop [pc 0 registers init-registers]
    (if (>= pc (count instructions)) registers
      (let [{:keys [opcode register value]} (nth instructions pc)]
        (case opcode
          "hlf" (recur (inc pc) (update registers register #(/ % 2)))
          "tpl" (recur (inc pc) (update registers register #(* % 3)))
          "inc" (recur (inc pc) (update registers register inc))
          "jmp" (recur (+ pc value) registers)
          "jie" (recur (if (even? (get registers register)) (+ pc value) (inc pc)) registers)
          "jio" (recur (if (= 1 (get registers register)) (+ pc value) (inc pc)) registers))))))

(println (format "With register a initialized to 0, the value in register b is %d." (get (run instructions {"a" 0 "b" 0}) "b")))
(println (format "With register a initialized to 1, the value in register b is %d." (get (run instructions {"a" 1 "b" 0}) "b")))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

