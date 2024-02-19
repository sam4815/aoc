(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def instructions
  (->> (str/split (slurp "input.txt") #"\n")
       (map #(str/split % #"\s"))
       (map #(list (first %) (or (parse-long (second %)) (second %)) (or (parse-long (last %)) (last %))))))

(defn update-registers [[opcode a b] registers]
  (case opcode
    "snd" (assoc registers :sent (concat (or (get registers :sent) '()) [(get registers a)]))
    "set" (assoc registers a b)
    "add" (assoc registers a (+ (or (get registers a) 0) b))
    "mul" (assoc registers a (* (or (get registers a) 0) b))
    "mod" (assoc registers a (rem (or (get registers a) 0) b))
    "rcv" (assoc (assoc registers :received (drop 1 (get registers :received))) a (first (get registers :received)))))

(defn is-deadlocked [instructions {:keys [reg pc]}]
  (and (= "rcv" (first (nth instructions pc)))
       (= (count (or (get reg :received) '())) 0)))

(defn process-instruction [[opcode a b-sym] {:keys [reg pc]}]
  (let [b (if (integer? b-sym) b-sym (get reg b-sym))]
    (case opcode
      "jgz" (if (> (or (get reg a) a) 0) (list reg (+ pc b)) (list reg (inc pc)))
      (list (update-registers (list opcode a b) reg) (inc pc)))))

(defn process-instructions [instructions init-program]
  (loop [program init-program]
    (if (is-deadlocked instructions program) program
      (let [[next-reg next-pc] (process-instruction (nth instructions (get program :pc)) program)]
        (recur { :reg next-reg :pc next-pc })))))

(defn pass-messages [p0 p1]
  (list
    (update-in p0 [:reg] assoc :received (get-in p1 [:reg :sent]) :sent '())
    (update-in p1 [:reg] assoc :received (get-in p0 [:reg :sent]) :sent '())))

(defn run-two-programs [instructions]
  (loop [p0 { :reg { "p" 0 } :pc 0 } p1 { :reg { "p" 1 } :pc 0 } p1-sent ()]
    (if (and (is-deadlocked instructions p0) (is-deadlocked instructions p1)) p1-sent
      (let [[next-p0 next-p1] (pass-messages (process-instructions instructions p0) (process-instructions instructions p1))]
        (recur next-p0 next-p1 (concat p1-sent (get-in p0 [:reg :received])))))))

(def single-program (process-instructions instructions { :reg {} :pc 0 }))
(def last-sent (last (get-in single-program [:reg :sent])))
(def messages (run-two-programs instructions))

(println (format "The first time a rcv instruction is executed, the last value sent is %d." last-sent))
(println (format "Once both programs have terminated, program 1 sent %d messages." (count messages)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

