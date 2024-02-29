(ns main
  (:require [clojure.string :as str])
  (:import [java.security MessageDigest]))

(def start-time (System/currentTimeMillis))

(def salt (str/trim (slurp "input.txt")))

(defn md5 [string]
  (let [raw (.digest (MessageDigest/getInstance "MD5") (.getBytes string))]
    (format "%032x" (BigInteger. 1 raw))))

(defn find-hash [target salt stretch-factor]
  (let [hashes (mapv #(nth (iterate md5 (str salt %)) stretch-factor) (range 25000))]
    (loop [current-index 0 num-hashes-found 0]
      (if (= target num-hashes-found) (dec current-index)
        (let [current-hash (nth hashes current-index) triple (re-find #"(\w)\1\1" current-hash)]
          (if (some? triple)
            (let [pattern (re-pattern (str "(" (first (first triple)) ")" "\\1\\1\\1\\1"))
                  quintuple (re-find pattern (str/join "," (subvec hashes (inc current-index) (+ current-index 1000))))]
              (if (some? quintuple)
                (recur (inc current-index) (inc num-hashes-found))
                (recur (inc current-index) num-hashes-found)))
            (recur (inc current-index) num-hashes-found)))))))

(println (format "The index of the 64th hash is %d." (find-hash 64 salt 1)))
(println (format "With key stretching, the index of the 64th hash is %d." (find-hash 64 salt 2017)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

