(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def password (str/trim (slurp "input.txt")))

(defn drop-last-char [letters]
  (subs letters 0 (dec (count letters))))

(defn next-letter [letter]
  (char (inc (int letter))))

(defn next-letters [letters]
  (if (= (last letters) \z)
    (str (next-letters (drop-last-char letters)) \a)
    (str (drop-last-char letters) (next-letter (last letters)))))

(defn has-straight [letters]
  (loop [n 0]
    (if (= n (- (count letters) 2)) false
      (or (and (= (next-letter (get letters n)) (get letters (inc n)))
               (= (next-letter (get letters (inc n))) (get letters (+ n 2))))
          (recur (inc n))))))

(defn is-valid-password [letters]
  (and (nil? (re-seq #"[iol]" letters))
       (>= (count (re-seq #"(\w)\1" letters)) 2)
       (has-straight letters)))

(defn next-password [password]
  (loop [letters (next-letters password)]
    (if (is-valid-password letters) letters (recur (next-letters letters)))))

(println (format "Santa's next password should be %s." (next-password password)))
(println (format "Santa's next next password should be %s." (next-password (next-password password))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

