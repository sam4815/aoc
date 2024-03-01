(ns main
  (:require [clojure.string :as str])
  (:import [java.security MessageDigest]))

(def start-time (System/currentTimeMillis))

(def passcode (str/trim (slurp "input.txt")))

(defn md5 [string]
  (let [raw (.digest (MessageDigest/getInstance "MD5") (.getBytes string))]
    (format "%032x" (BigInteger. 1 raw))))

(defn get-direction [index] (get { 0 \U 1 \D 2 \L 3 \R } index))

(defn is-door-open [door] (str/includes? "bcdef" (str door)))

(defn is-position-valid [[x y]] (and (>= x 0) (>= y 0) (< x 4) (< y 4)))

(defn apply-direction [[x y] direction]
  (case direction
    \U `(~x ~(dec y)) \D `(~x ~(inc y))
    \L `(~(dec x) ~y) \R `(~(inc x) ~y)))

(defn find-next-paths [current-position path passcode]
  (let [path-hash (md5 (str passcode (apply str path)))
        doors (map-indexed #(list (get-direction %1) (apply-direction current-position (get-direction %1)) %2) (take 4 path-hash))
        valid-doors (filter (fn [[direction position door]] (and (is-door-open door) (is-position-valid position))) doors)]
    (map (fn [[direction position]] { :position position :path (concat path [direction]) }) valid-doors)))

(defn find-best-path [comparison passcode]
  (loop [queue [{ :position '(0 0) :path [] }] best-path nil]
    (if (empty? queue) best-path
      (let [{:keys [position path]} (first queue)]
        (if (= position '(3 3))
          (if (or (comparison (count path) (count best-path)) (nil? best-path)) (recur (rest queue) path) (recur (rest queue) best-path))
          (recur (concat (find-next-paths position path passcode) (rest queue)) best-path))))))

(println (format "The shortest path to the vault is %s." (str/join (find-best-path < passcode))))
(println (format "The length of the longest path to the vault is %d." (count (find-best-path > passcode))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

