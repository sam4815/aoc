(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def molecule (str/trim (second (str/split (slurp "input.txt") #"\n\n"))))

(def mappings (->> (str/split (first (str/split (slurp "input.txt") #"\n\n")) #"\n")
                   (map #(str/split % #" => "))))

(def forward-map (reduce (fn [all [k v]] (assoc all k (cons v (get all k [])))) {} mappings))
(def backward-map (into {} (map (comp vec reverse) mappings)))

(defn replace-at [string index length replacement]
  (str (subs string 0 index) replacement (subs string (+ index length))))

(defn find-molecules [molecule mappings]
  (loop [i 0 molecules []]
    (if (>= i (count molecule)) molecules
      (recur (inc i)
             (concat
               (map (partial replace-at molecule i 1) (get mappings (subs molecule i (+ i 1)) []))
               (map (partial replace-at molecule i 2) (get mappings (subs molecule i (min (count molecule) (+ i 2))) []))
               molecules)))))

(defn find-shortest-reduction [init-molecule mappings]
  (loop [queue [{:molecule init-molecule :steps 0}] visited {}]
    (let [{:keys [molecule steps]} (first queue)]
      (if (= molecule "e") steps
        (if (>= steps (get visited molecule Integer/MAX_VALUE))
          (recur (rest queue) visited)
          (recur (concat (map (fn [[k v]] {:molecule (str/replace-first molecule k v) :steps (inc steps)}) mappings)
                         (rest queue))
                 (assoc visited molecule steps)))))))

(def distinct-molecules (set (find-molecules molecule forward-map)))
(def min-steps (find-shortest-reduction molecule backward-map))

(println (format "%d distinct molecules can be created." (count distinct-molecules)))
(println (format "The fewest number of steps required for the medicine molecule is %d." min-steps))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

