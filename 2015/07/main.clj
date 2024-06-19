(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def instructions (->> (str/split (slurp "input.txt") #"\n")
                       (map (fn [line] {:operator (first (or (re-seq #"[A-Z]+" line) '("PUT")))
                                        :operands (re-seq #"[a-z0-9]+" line)}))
                       (map (juxt (comp last :operands) identity))
                       (into {})))

(defn find-signal [instructions]
  (def memoized (memoize (fn [signal]
                           (let [{:keys [operands operator]} (get instructions signal)
                                 parse-operand (fn [[parsed operand]] (if (integer? parsed) parsed (memoized operand)))
                                 [a b] (map (comp parse-operand (juxt parse-long identity)) (drop-last operands))]
                             (bit-and
                               0xFFFF
                               (case operator
                                 "PUT" a
                                 "NOT" (bit-not a)
                                 "AND" (bit-and a b)
                                 "OR" (bit-or a b)
                                 "LSHIFT" (bit-shift-left a b)
                                 "RSHIFT" (bit-shift-right a b))))))))

(def a-signal ((find-signal instructions) "a"))
(def modified-a-signal ((find-signal (assoc instructions "b" {:operator "PUT" :operands [(str a-signal) "b"]})) "a"))

(println (format "The signal provided to wire a is %d." a-signal))
(println (format "After modifying b, the signal provided to wire a is %d." modified-a-signal))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

