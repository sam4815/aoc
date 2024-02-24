(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def sections (str/split (slurp "input.txt") #"\n\n"))

(defn parse-state [state]
  {:label (re-find #"[A-Z](?=:)" state) 
   :directions (re-seq #"left|right" state)
   :values (map #(parse-long %) (re-seq #"(?<=value )\d+" state))
   :states (re-seq #"[A-Z](?=\.)" state)})

(def states (into {} (map #(vector (get % :label) %) (map parse-state (rest sections)))))

(defn get-next-cursor [cursor direction] (get { "left" (dec cursor) "right" (inc cursor) } direction))

(defn run-turing [total-steps states]
  (loop [steps 0 tape {} cursor 0 current-state "A"]
    (if (= steps total-steps) tape
      (let [{:keys [directions values states]} (get states current-state) value (or (get tape cursor) 0)]
        (if (= value 0) (recur (inc steps) (assoc tape cursor (first values)) (get-next-cursor cursor (first directions)) (first states))
          (recur (inc steps) (assoc tape cursor (second values)) (get-next-cursor cursor (second directions)) (second states)))))))

(def turing-state (run-turing (parse-long (re-find #"\d+" (first sections))) states))
(def diagnostic-checksum (reduce + (vals turing-state)))

(println (format "Once the Turing machine is working again, its diagnostic checksum is %d." diagnostic-checksum))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

