(ns main
  (:require [clojure.string :as str])
  (:import [java.security MessageDigest]))

(def start-time (System/currentTimeMillis))

(def secret-key (str/trim (slurp "input.txt")))

(defn md5 [string]
  (let [raw (.digest (MessageDigest/getInstance "MD5") (.getBytes string))]
    (format "%032x" (BigInteger. 1 raw))))

(defn find-pattern [secret-key pattern]
  (loop [n 0]
    (let [current-hash (md5 (str secret-key n))]
      (if (str/starts-with? current-hash pattern) n (recur (inc n))))))

(println (format "The lowest number that produces 00000 is %d." (find-pattern secret-key "00000")))
(println (format "The lowest number that produces 000000 is %d." (find-pattern secret-key "000000")))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

