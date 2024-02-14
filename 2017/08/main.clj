(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(defmacro inc [a b] `(+ ~a ~b))
(defmacro dec [a b] `(- ~a ~b))
(defmacro != [a b] `(not= ~a ~b))

(defmacro instruction
  [[register-a operation value-a _ register-b comparison value-b] registers]
  `(if (~comparison (or (get ~registers '~register-b) 0) ~value-b)
     (assoc ~registers '~register-a (~operation (or (get ~registers '~register-a) 0) ~value-a))
     ~registers))

(def final-registers (reduce (fn [[registers max-register] instr]
                               (def regs registers)
                               (list
                                 (eval (read-string (format "(instruction (%s) regs)" instr)))
                                 (max max-register (apply max (or (vals registers) '(0))))))
                             '({} 0) (str/split (slurp "input.txt") #"\n")))

(println (format "The largest register value after running the instructions is %d." (apply max (vals (first final-registers)))))
(println (format "The largest register value during this process is %d." (second final-registers)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

