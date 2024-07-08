(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def boss-stats (zipmap [:hp :damage] (map parse-long (re-seq #"\d+" (slurp "input.txt")))))

(def spells [{:cost 53 :duration 1 :name "Magic Missile"}
             {:cost 73 :duration 1 :name "Drain"}
             {:cost 113 :duration 6 :name "Shield"}
             {:cost 173 :duration 6 :name "Poison"}
             {:cost 229 :duration 5 :name "Recharge"}])

(defn apply-effect [{:keys [player boss]} effect]
  (case (get effect :name)
    "Magic Missile" {:player player
                     :boss (assoc boss :hp (- (get boss :hp) 4))}
    "Drain" {:player (assoc player :hp (+ (get player :hp) 2))
             :boss (assoc boss :hp (- (get boss :hp) 2))}
    "Shield" {:player (assoc player :armor 7)
              :boss boss}
    "Poison" {:player player
              :boss (assoc boss :hp (- (get boss :hp) 3))}
    "Recharge" {:player (assoc player :mana (+ (get player :mana) 101))
                :boss boss}))

(defn apply-effects [round]
  (if (> (get-in round [:player :hp]) 0)
    (let [initial {:player (assoc (get round :player) :armor 0) :boss (get round :boss)}
          {:keys [player boss]} (reduce apply-effect initial (get round :effects))]
      {:player player
       :boss boss
       :mana-spent (get round :mana-spent)
       :effects (filter #(not= 0 (get % :duration)) (map #(update % :duration dec) (get round :effects)))})
    round))

(defn apply-attack [round]
  (if (> (get-in round [:boss :hp]) 0)
    (update-in round [:player :hp] (fn [hp] (- hp (- (get-in round [:boss :damage]) (get-in round [:player :armor])))))
    round))

(defn get-spell-options [{:keys [player boss effects mana-spent]}]
  (map (fn [spell] (do {:player (update (update player :mana (fn [mana] (- mana (get spell :cost))))
                                        :hp
                                        (fn [hp] (- hp (get boss :penalty))))
                        :boss boss
                        :effects (cons spell effects)
                        :mana-spent (+ (get spell :cost) mana-spent)}))
       (filter (fn [spell] (and (>= (get player :mana) (get spell :cost))
                                (not (some #(= (get % :name) (get spell :name)) effects)))) spells)))

(defn find-least-mana [boss-init]
  (loop [queue (get-spell-options {:player {:hp 50 :mana 500 :armor 0} :boss boss-init :effects [] :mana-spent 0})
         least-mana Integer/MAX_VALUE]
    (if (zero? (count queue)) least-mana
      (let [round (apply-effects (apply-attack (apply-effects (first queue))))]
        (if (<= (get-in round [:boss :hp]) 0)
          (recur (rest queue) (min (get round :mana-spent) least-mana))
          (if (or (>= (get round :mana-spent) least-mana)
                  (<= (get-in round [:player :hp]) 0))
            (recur (rest queue) least-mana)
            (recur (concat (get-spell-options round) (rest queue))
                   least-mana)))))))

(println (format "The least amount of mana required to win is %d." (find-least-mana (assoc boss-stats :penalty 0))))
(println (format "In hard mode, the least amount of mana required to win is %d." (find-least-mana (assoc boss-stats :penalty 1))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

