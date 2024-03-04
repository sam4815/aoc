(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def instructions
  (->> (str/split (slurp "input.txt") #"\n")
       (map #(str/split % #"\s"))
       (mapv #(list (first %) (or (parse-long (second %)) (second %)) (or (parse-long (last %)) (last %))))))

(defn toggle-instruction [[opcode x y]]
  (case opcode
    "inc" ["dec" x y]
    "dec" ["inc" x y]
    "tgl" ["inc" x y]
    "jnz" ["cpy" x y]
    "cpy" ["jnz" x y]))

(defn update-registers [[opcode x y] registers]
  (case opcode
    "cpy" (assoc registers y (if (int? x) x (get registers x 0)))
    "inc" (assoc registers x (inc (get registers x 0)))
    "dec" (assoc registers x (dec (get registers x 0)))))

(defn process-instruction [[opcode x y] {:keys [reg pc]}]
  (case opcode
    "jnz" (if (not= (if (int? x) x (get reg x 0)) 0) [reg (+ pc (if (int? y) y (get reg y 0)))] [reg (inc pc)])
    [(update-registers [opcode x y] reg) (inc pc)]))

(defn process-instructions [init-instructions init-program]
  (loop [{:keys [reg pc]} init-program instructions init-instructions]
    (if (>= pc (count instructions)) {:pc pc :reg reg}
      (let [[opcode x y] (nth instructions pc)]
        (case opcode
          "tgl" (let [offset (+ pc (get reg x))]
                  (recur {:pc (inc pc) :reg reg}
                         (if (< offset (count instructions))
                           (assoc instructions offset (toggle-instruction (nth instructions offset)))
                           instructions)))
          (let [[next-reg next-pc] (process-instruction [opcode x y] {:reg reg :pc pc})]
            (recur {:reg next-reg :pc next-pc} instructions)))))))

(def program (process-instructions instructions {:reg {"a" 7} :pc 0}))

(def product (* (second (nth instructions 19)) (second (nth instructions 20))))
(def result-of-a-12 (+ product (reduce * (range 1 (inc 12)))))

(println (format "After executing the assembunny code, the value in register a is %d." (get-in program [:reg "a"])))
(println (format "The value that needs to be sent to the safe is %d." result-of-a-12))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

