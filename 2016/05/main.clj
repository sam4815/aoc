(ns main
  (:require [clojure.string :as str])
  (:import [java.security MessageDigest]))

(def start-time (System/currentTimeMillis))

(def door-id (str/trim (slurp "input.txt")))

(defn md5 [string]
  (let [raw (.digest (MessageDigest/getInstance "MD5") (.getBytes string))]
    (format "%032x" (BigInteger. 1 raw))))

(defn find-hashes [string]
  (loop [hashes '() n 0]
    (if (= 20 (count hashes)) hashes
      (let [current-hash (md5 (str string n))]
        (recur (if (str/starts-with? current-hash "00000") (concat hashes [current-hash]) hashes) (inc n))))))

(def hashes (find-hashes door-id))

(def first-password (str/join (map #(nth % 5) (take 8 hashes))))

(def key-hashes (map (fn [i] (first (filter #(= (nth % 5) i) hashes))) "01234567"))
(def second-password (str/join (map #(nth % 6) key-hashes)))

(println (format "The first password is %s." first-password))
(println (format "The second password is %s." second-password))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

