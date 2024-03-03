(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(defn parse-instruction [string]
  (let [parts (str/split string #"\s")]
    (case (str/join " " (take 2 parts))
      "swap position" `(:swap-position ~(parse-long (nth parts 2)) ~(parse-long (nth parts 5)))
      "swap letter" `(:swap-letter ~(first (nth parts 2)) ~(first (nth parts 5)))
      "rotate based" `(:rotate-position ~(first (last parts)))
      "rotate left" `(:rotate-left ~(parse-long (nth parts 2)))
      "rotate right" `(:rotate-right ~(parse-long (nth parts 2)))
      "reverse positions" `(:reverse ~(parse-long (nth parts 2)) ~(parse-long (nth parts 4)))
      "move position" `(:move ~(parse-long (nth parts 2)) ~(parse-long (nth parts 5))))))

(def instructions (map parse-instruction (str/split (slurp "input.txt") #"\n")))

(defn rotate-right [word n]
  (vec (concat (subvec word (- (count word) (mod n (count word))) (count word))
               (subvec word 0 (- (count word) (mod n (count word)))))))

(defn scramble-instruction [word [instruction a b]]
  (case instruction
    :swap-position (assoc (assoc word b (nth word a)) a (nth word b))
    :swap-letter (mapv #(if (= a %) b (if (= b %) a %)) word)
    :rotate-position (let [index (.indexOf word a)] (rotate-right word (+ (if (>= index 4) 2 1) index)))
    :rotate-left (rotate-right word (- (count word) a))
    :rotate-right (rotate-right word a)
    :reverse (vec (concat (subvec word 0 a) (reverse (subvec word a (inc b))) (subvec word (inc b))))
    :move (let [letter (nth word a) [start end] (split-at b (filter #(not= letter %) word))]
            (vec (concat start `(~letter) end)))))

(defn determine-rotation [init-word a]
  (loop [word init-word]
    (if (= (scramble-instruction word [:rotate-position a]) init-word) word (recur (rotate-right word 1)))))

(defn unscramble-instruction [word [instruction a b]]
  (case instruction
    (:swap-position :move) (scramble-instruction word [instruction b a])
    (:swap-letter :reverse) (scramble-instruction word [instruction a b])
    :rotate-position (determine-rotation word a)
    :rotate-left (rotate-right word a)
    :rotate-right (rotate-right word (- (count word) a))))

(defn scramble [word instructions]
  (str/join (reduce scramble-instruction (vec word) instructions)))

(defn unscramble [word instructions]
  (str/join (reduce unscramble-instruction (vec word) (reverse instructions))))

(println (format "The result of scrambling abcdefgh is %s." (scramble "abcdefgh" instructions)))
(println (format "The result of unscrambling fbgdceah is %s." (unscramble "fbgdceah" instructions)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

