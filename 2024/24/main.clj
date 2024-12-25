(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def parts (str/split (slurp "input.txt") #"\n\n"))

(def init-values (->> (str/split (first parts) #"\n")
                      (map (fn [line] (re-seq #"[\w\d]+" line)))
                      (map (fn [[a b]] [a {:op "IN" :gates [(parse-long b)]}]))
                      (into {})))

(def init-ops (->> (str/split (second parts) #"\n")
                   (map (fn [line] (re-seq #"[\w\d]+" line)))
                   (map (fn [[a op b c]] [c {:op op :gates [a b]}]))
                   (into {})))

(defn to-gates [binary letter]
  (map-indexed (fn [i digit] [(str letter (format (str "%0" 2 "d") i))
                              {:op "IN" :gates [(parse-long (str digit))]}])
               (reverse binary)))

(defn swap [[a-gate b-gate] ops]
  (into ops {a-gate (get ops b-gate) b-gate (get ops a-gate)}))

(defn with-input [x y ops]
  (let [x-binary (format (str "%0" 45 "d") (BigInteger. (Long/toString x 2)))
        y-binary (format (str "%0" 45 "d") (BigInteger. (Long/toString y 2)))]
    (into (into ops (to-gates x-binary "x")) (to-gates y-binary "y"))))

(defn solve [ops visited gate]
  (if (get visited gate) 0
    (let [{:keys [op gates]} (get ops gate)]
      (case op
        "IN" (first gates)
        "XOR" (apply bit-xor (map #(solve ops (assoc visited gate true) %) gates))
        "AND" (apply bit-and (map #(solve ops (assoc visited gate true) %) gates))
        "OR" (apply bit-or (map #(solve ops (assoc visited gate true) %) gates)) 0))))

(defn find-digits-by-letter [letter ops]
  (->> (filter (fn [[label gate]] (= (first label) letter)) ops)
       (sort-by (fn [[label gate]] label))
       reverse
       keys
       (map (partial solve ops {}))
       str/join))

(defn find-dependents [gate ops]
  (loop [queue [gate] dependents []]
    (if (empty? queue) (distinct dependents)
      (let [curr (get ops (first queue))]
        (if (not curr)
          (recur (rest queue) dependents)
          (recur (concat (rest queue) (get curr :gates)) (conj dependents (first queue))))))))

(defn calc [x y ops]
  (let [ops (with-input x y ops)]
    (map #(Long/parseLong (find-digits-by-letter % ops) 2) [\x \y \z])))

(defn find-broken-indices [offset ops]
  (loop [i 0 n 1 indices []]
    (if (> n (Math/pow 2 45)) indices
      (let [[x y z] (calc (+ n offset) 1 ops)]
        (if (= (+ x y) z)
          (recur (inc i) (* n 2) indices)
          (recur (inc i) (* n 2) (conj indices i)))))))

(defn find-swaps [i ops]
  (let [number (Math/pow 2 i) broken-gate (format (str "z" "%0" 2 "d") i)]
    (->> (for [a-gate (concat [broken-gate] (get (get ops broken-gate) :gates)) b-gate (keys ops)] [a-gate b-gate])
         (filter (fn [[a-gate b-gate]] (not= a-gate b-gate)))
         (filter (fn [[a-gate b-gate]] (let [[x y z] (calc number 1 (swap [a-gate b-gate] ops))] (= (+ x y) z))))
         (map (fn [[a-gate b-gate]] [a-gate b-gate])))))

(defn find-fix [ops [as bs cs ds]]
  (->> (for [a as b bs c cs d ds] [a b c d])
       (filter (fn [[a b c d]] (zero? (count (find-broken-indices -1 (swap d (swap c (swap b (swap a ops)))))))))
       first))

(def z (find-digits-by-letter \z (into init-values init-ops)))
(def fixes (find-fix init-ops (map #(find-swaps % init-ops) (find-broken-indices 0 init-ops))))

(println (format "The system outputs %d." (Long/parseLong z 2)))
(println (format "The eight wires that need swapping are %s." (str/join "," (sort (apply concat fixes)))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

