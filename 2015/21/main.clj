(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def boss-stats (zipmap [:hp :damage :armor] (map parse-long (re-seq #"\d+" (slurp "input.txt")))))

(def weapons [{:cost 8 :damage 4 :armor 0} {:cost 10 :damage 5 :armor 0} {:cost 25 :damage 6 :armor 0}
              {:cost 40 :damage 7 :armor 0} {:cost 74 :damage 8 :armor 0}])

(def armors [{:cost 0 :damage 0 :armor 0} {:cost 13 :damage 0 :armor 1} {:cost 31 :damage 0 :armor 2}
             {:cost 53 :damage 0 :armor 3} {:cost 75 :damage 0 :armor 4} {:cost 102 :damage 0 :armor 5}])

(def rings [{:cost 0 :damage 0 :armor 0} {:cost 25 :damage 1 :armor 0} {:cost 50 :damage 2 :armor 0}
            {:cost 100 :damage 3 :armor 0} {:cost 20 :damage 0 :armor 1} {:cost 40 :damage 0 :armor 2}
            {:cost 80 :damage 0 :armor 3}])

(def loadouts (for [weapon weapons
                    armor armors
                    ring-a rings
                    ring-b (filter (partial not= ring-a) rings)]
                (reduce (fn [stats item] (merge-with + stats item)) {:hp 100} [weapon armor ring-a ring-b])))

(defn find-winner [boss-init player-init]
  (let [player-damage (max 1 (- (get player-init :damage) (get boss-init :armor)))
        boss-damage (max 1 (- (get boss-init :damage) (get player-init :armor)))]
    (loop [n 0 boss boss-init player player-init]
      (if (or (<= (get boss :hp) 0) (<= (get player :hp) 0))
        (if (> (get player :hp) 0) :player :boss)
        (if (even? n)
          (recur (inc n) (assoc boss :hp (- (get boss :hp) player-damage)) player)
          (recur (inc n) boss (assoc player :hp (- (get player :hp) boss-damage))))))))

(def winners (group-by (partial find-winner boss-stats) loadouts))

(def min-gold (apply min (map :cost (get winners :player))))
(def max-gold (apply max (map :cost (get winners :boss))))

(println (format "The least amount of gold required to win is %d." min-gold))
(println (format "The most amount of gold that still results in a loss is %d." max-gold))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

