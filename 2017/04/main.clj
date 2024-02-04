(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def passphrases (->> (str/split (slurp "input.txt") #"\n")
                      (map #(str/split % #" "))))

(defn filter-valid-v1 [passes]
  (filter #(= (count %) (count (set %))) passes))

(defn are-anagrams [word-a word-b] (= (frequencies word-a) (frequencies word-b)))

(defn has-no-anagrams [words]
  (every? (fn [word] (not-any? #(are-anagrams word %) (remove #(= word %) words))) words))

(defn filter-valid-v2 [passes]
  (filter #(has-no-anagrams %) (filter-valid-v1 passes)))

(println (format "The number of valid passphrases under the existing policy is %d." (count (filter-valid-v1 passphrases))))
(println (format "The number of valid passphrases under the new policy is %d." (count (filter-valid-v2 passphrases))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

