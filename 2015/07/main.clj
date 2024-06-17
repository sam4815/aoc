(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def instructions (->> (str/split (slurp "input.txt") #"\n")
                       (map (fn [line] {:operator (first (or (re-seq #"[A-Z]+" line) '("PUT")))
                                        :operands (re-seq #"[a-z0-9]+" line)}))
                       (map (juxt (comp last :operands) identity))
                       (into {})))

(defn parse-operand [operand instructions]
  (if (integer? (parse-long operand))
    (parse-long operand)
    ((resolve 'find-signal-memo) operand instructions)))

(defn find-signal [signal instructions]
  (let [instruction (get instructions signal)
        operands (map #(parse-operand % instructions) (drop-last (get instruction :operands)))]
    (case (get instruction :operator)
      "PUT" (first operands)
      "NOT" (bit-and (bit-not (first operands)) 16rFFFF)
      "AND" (bit-and (first operands) (second operands) 16rFFFF)
      "OR" (bit-and (bit-or (first operands) (second operands)) 16rFFFF)
      "LSHIFT" (bit-and (bit-shift-left (first operands) (second operands)) 16rFFFF)
      "RSHIFT" (bit-and (bit-shift-right (first operands) (second operands)) 16rFFFF))))

(def find-signal-memo (memoize find-signal))

(def a-signal (find-signal-memo "a" instructions))
(def modified-a-signal (find-signal-memo "a" (assoc instructions "b" {:operator "PUT" :operands (list (str a-signal) "b")})))

(println (format "The signal provided to wire a is %s." a-signal))
(println (format "After modifying b, the signal provided to wire a is %s." modified-a-signal))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

