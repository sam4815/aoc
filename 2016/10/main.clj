(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def instructions (map (fn [line] (re-seq #"\w+ \d+" line)) (str/split (slurp "input.txt") #"\n")))

(def bots (reduce #(merge-with into %1 %2) {}
                  (map #(do {(last %) (list (parse-long (subs (first %) 6)))}) (filter #(= 2 (count %)) instructions))))
(def rules (into {} (map #(do [(first %) {:low (second %) :high (last %)}]) (filter #(= 3 (count %)) instructions))))

(defn run-bots [init-bots rules]
  (loop [bots-bins init-bots]
    (let [two-chip-bot (first (filter (comp #{2} count val) bots-bins))]
      (if (nil? two-chip-bot) bots-bins
        (let [[low-val high-val] (sort (val two-chip-bot)) {:keys [low high]} (get rules (key two-chip-bot))]
          (if (and (= low-val 17) (= high-val 61))
            (println (format "The bot that compares values 17 and 61 is %s." (key two-chip-bot))))
          (recur (merge-with into (assoc bots-bins (key two-chip-bot) '()) {low `(~low-val) high `(~high-val)})))))))

(def executed-bots (run-bots bots rules))
(def chip-product (reduce * (apply concat (map (partial get executed-bots) ["output 0" "output 1" "output 2"]))))

(println (format "The product of the chips in the first three output bins is %d." chip-product))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

