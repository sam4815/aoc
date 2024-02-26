(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def file (str/replace (slurp "input.txt") #"\n" ""))

(defn parse-marker [file]
  (let [marker (subs file 0 (str/index-of file ")"))
        nums (map #(parse-long %) (re-seq #"\d+" marker))]
    (list (first nums) (second nums) (inc (str/index-of file ")")))))

(defn decompress [file]
  (loop [i 0 decompressed ""]
    (if (>= i (count file)) decompressed
      (let [current-char (nth file i)]
        (if (= current-char \()
          (let [[length n offset] (parse-marker (subs file i))]
            (recur (+ i offset length) (str decompressed (apply str (repeat n (subs file (+ i offset) (+ i offset length)))))))
          (recur (inc i) (str decompressed current-char)))))))

(defn find-decompressed-length [file multiplier]
  (loop [i 0 file-length 0]
    (if (>= i (count file)) file-length
      (let [current-char (nth file i)]
        (if (= current-char \()
          (let [[marker-length n offset] (parse-marker (subs file i))
                subseq-length (find-decompressed-length (subs file (+ i offset) (+ i offset marker-length)) (* n multiplier))]
            (recur (+ i offset marker-length) (+ file-length subseq-length)))
          (recur (inc i) (+ file-length multiplier)))))))

(println (format "The decompressed length of the file is %d." (count (decompress file))))
(println (format "Using the improved format, the decompressed length of the file is %d." (find-decompressed-length file 1)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

