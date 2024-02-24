(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def components (mapv (fn [line] (map #(parse-long %) (str/split line #"\/"))) (str/split (slurp "input.txt") #"\n")))

(defn find-matching-components [target available]
  (filter (fn [[index candidate]] (some #{target} candidate)) (map-indexed #(list %1 %2) available)))

(defn get-next-matches [{:keys [available bridge complete target]}]
  (if complete [{ :available available :bridge bridge :complete complete :target target }]
    (let [matches (find-matching-components target available)]
      (if (= 0 (count matches)) [{ :available available :bridge bridge :complete true :target target }]
        (map (fn [[index component]] {:available (vec (concat (subvec available 0 index) (subvec available (inc index))))
                                      :bridge (conj bridge component)
                                      :complete false
                                      :target (or (first (filter #(not= target %) component)) target) })
             matches)))))

(defn find-permutations [comps]
  (loop [bridges [{ :available comps :bridge [] :complete false :target 0 }]]
    (if (every? :complete bridges) bridges (recur (apply concat (map get-next-matches bridges))))))

(defn find-strength [bridge] (reduce + (map #(reduce + %) bridge)))

(def bridges (map :bridge (find-permutations components)))
(def max-strength (apply max (map find-strength bridges)))
(def max-length (apply max (map count bridges)))
(def max-length-strength (apply max (map find-strength (filter #(= max-length (count %)) bridges))))

(println (format "The strength of the strongest bridge is %d." max-strength))
(println (format "The strength of the strongest longest bridge is %d." max-length-strength))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

