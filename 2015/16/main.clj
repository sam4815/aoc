(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def reading { "children" 3 "cats" 7 "samoyeds" 2 "pomeranians" 3 "akitas" 0
              "vizslas" 0 "goldfish" 5 "trees" 3 "cars" 2 "perfumes" 1 })

(def sues (->> (str/split (slurp "input.txt") #"\n")
               (map #(zipmap (cons :id (re-seq #"(?<=\s)[a-z]+" %))
                             (map parse-long (re-seq #"\d+" %))))))

(defn exact-match [reading sue]
  (every? (fn [[k v]] (= v (get reading k))) (dissoc sue :id)))

(defn range-match [reading sue]
  (every? (fn [[k v]] (case k
                        ("cats" "trees") (> v (get reading k))
                        ("pomeranians" "goldfish") (< v (get reading k))
                        (= v (get reading k))))
          (dissoc sue :id)))

(def exact-sue (first (filter (partial exact-match reading) sues)))
(def range-sue (first (filter (partial range-match reading) sues)))

(println (format "At first, it seems that Sue %d got you the gift." (get exact-sue :id)))
(println (format "The real Aunt Sue is Sue %d." (get range-sue :id)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

