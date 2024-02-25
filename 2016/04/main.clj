(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def rooms (map (fn [line] (let [parts (str/split line #"-")] {:id (parse-long (re-find #"\d+" (last parts)))
                                                               :order (re-find #"[a-z]+" (last parts))
                                                               :words (butlast parts)}))
                (str/split (slurp "input.txt") #"\n")))

(defn is-real-room [{:keys [order words]}]
  (= order (str/join (take 5 (keys (sort-by (juxt (comp - val) key) (frequencies (apply concat words))))))))

(defn rotate-character [character n]
  (char (+ (mod (+ (- (int character) 97) n) 26) 97)))

(def real-rooms (filter is-real-room rooms))
(def real-room-sum (reduce + (map :id real-rooms)))

(def decrypted-rooms (map (fn [{:keys [id words]}] {:id id :words (map (fn [word] (str/join (map #(rotate-character % id) word))) words)}) real-rooms))

(def north-pole-room (first (filter (fn [{:keys [words]}] (some (partial = "northpole") words)) decrypted-rooms)))

(println (format "The sum of the vector IDs of the real rooms is %d." real-room-sum))
(println (format "The room containing the North Pole objects has ID %d." (get north-pole-room :id)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

