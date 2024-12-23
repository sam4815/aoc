(ns main
  (:require [clojure.string :as str]
            [clojure.set :as set]))

(def start-time (System/currentTimeMillis))

(def codes (map seq (str/split (slurp "input.txt") #"\n" )))

(def numerical
  { \7 [0 0] \8 [0 1] \9 [0 2] \4 [1 0] \5 [1 1] \6 [1 2] \1 [2 0] \2 [2 1] \3 [2 2] \0 [3 1] \A [3 2] })

(defn get-manhattan [[x y] [i j]]
  (+ (abs (- x i)) (abs (- y j))))

(defn find-numeric [start end]
  (let [[x y] (get numerical start)
        moves [[\^ (dec x) y] [\v (inc x) y] [\< x (dec y)] [\> x (inc y)]]
        min-dist (apply min (map (fn [[move i j]] (get-manhattan (get numerical end) [i j])) moves))]
    (->> (filter (fn [[move i j]] (and (not= [i j] [3 0])
                                       (= min-dist (get-manhattan [i j] (get numerical end))))) moves)
         (map (fn [[move i j]] [move (get (set/map-invert numerical) [i j])])))))

(defn find-subpaths [start end steps-fn]
  (loop [queue [{:button start :path []}] paths []]
    (if (empty? queue) paths
      (let [{:keys [button path]} (first queue)]
        (if (= button end) (recur (rest queue) (conj paths path))
          (recur (concat (map (fn [[step button]] {:button button :path (conj path step)}) (steps-fn button end))
                         (rest queue)) paths))))))

(def memo-find-subpaths (memoize find-subpaths))

(defn get-subpaths [start end]
  (case [start end]
    [\A \^] [\< \A] [\A \v] [\< \v \A] [\A \>] [\v \A] [\A \<] [\v \< \< \A]
    [\< \A] [\> \> \^ \A] [\< \^] [\> \^ \A] [\< \v] [\> \A] [\< \>] [\> \> \A]
    [\^ \A] [\> \A] [\^ \v] [\v \A] [\^ \>] [\v \> \A] [\^ \<] [\v \< \A]
    [\> \A] [\^ \A] [\> \^] [\< \^ \A] [\> \v] [\< \A] [\> \<] [\< \< \A]
    [\v \A] [\^ \> \A] [\v \^] [\^ \A] [\v \>] [\> \A] [\v \<] [\< \A]
    [\A]))

(defn find-abs-paths [code]
  (concat (concat (get-subpaths \A (first code))
                  (apply concat (map-indexed (fn [i c] (get-subpaths (nth code i) c)) (drop 1 code)))
                  (get-subpaths (last code) \A))))

(defn score-transformation [code]
  (reduce + (map (fn [[k v]] (* (count k) v)) code)))

(defn transform [code]
  (let [freq (frequencies (str/split (str/join code) #"A"))]
    (assoc freq "l" (reduce + (vals freq)))))

(defn transform-next [code]
  (reduce (fn [a [k v]]
            (reduce (fn [a pair] (assoc a pair (+ (get a pair 0) v)))
                    a
                    (str/split (str/join (find-abs-paths (seq k))) #"A")))
          {"l" (score-transformation code)}
          (filter (fn [[k v]] (not= k "l")) code)))

(defn find-paths [code steps-fn]
  (loop [queue [{:button \A :index 0 :path []}] paths []]
    (if (empty? queue) paths
      (let [{:keys [button index path]} (first queue)]
        (if (>= index (count code)) (recur (rest queue) (conj paths path))
          (recur (concat (map (fn [subpath] {:button (nth code index) :index (inc index) :path (concat path (conj subpath \A))})
                              (memo-find-subpaths button (nth code index) steps-fn))
                         (rest queue))
                 paths))))))

(defn find-shortest-sequence [num-robots code]
  (->> (find-paths code find-numeric)
       (map transform)
       (map #(nth (iterate transform-next %) num-robots))
       (map score-transformation)
       (apply min)))

(defn get-complexity [num-robots code]
  (* (find-shortest-sequence num-robots code)
     (parse-long (apply str (drop-last code)))))

(println (format "With 2 keypads, the sum of the complexities is %d." (reduce + (map (partial get-complexity 2) codes))))
(println (format "With 25 keypads, the sum of the complexities is %d." (reduce + (map (partial get-complexity 25) codes))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

