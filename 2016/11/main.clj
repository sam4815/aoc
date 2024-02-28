(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(defn parse-component [string]
  (let [words (str/split string #"[^a-z]")] {:component (last words) :element (first words)}))

(def floors (map (fn [line] (map parse-component (remove nil? (re-seq #"[a-z\-]+\s(?:generator|microchip)" line))))
                 (str/split (slurp "input.txt") #"\n")))

(def facility {:level 0 :steps 0 :floors floors})

(def additional-components '({:element "elerium" :component "generator"} {:element "elerium" :component "microchip"}
                             {:element "dilithium" :component "generator"} {:element "dilithium" :component "microchip"}))

(def facility-plus {:steps 0 :level 0 :floors (map-indexed #(if (= %1 0) (concat %2 additional-components) %2) floors)})

(defn ready-for-assembly [{:keys [floors]}] (every? empty? (drop-last floors)))

(defn is-floor-safe [floor]
  (let [groups (group-by :component floor) generators (get groups "generator") chips (get groups "microchip")]
    (every? (fn [{:keys [element]}] (or (some (comp #{element} :element) generators) (= 0 (count generators)))) chips)))

(defn is-facility-safe [{:keys [floors]}] (every? is-floor-safe floors))

(defn get-combos [floor]
  (concat (map list floor)
          (apply concat (map-indexed (fn [i component] (map #(list component %) (drop (inc i) floor))) (drop-last floor)))))

(defn move-floor [facility components source target]
  (map-indexed (fn [i floor] (if (= i target) (concat floor components)
                               (if (= i source) (filter (fn [c] (not (some #{c} components))) floor) floor))) facility))

(defn hash-facility [{:keys [level floors]}] {:level level :floors (map (fn [floor] (frequencies (map :component floor))) floors)})

(defn get-valid-moves [{:keys [level floors steps]} visited]
  (let [combos (get-combos (nth floors level))
        valid-levels (filter #(and (>= % 0) (<= % 3)) `(~(dec level) ~(inc level)))
        all-moves (apply concat (map (fn [target] (map #(do {:level target :floors (move-floor floors % level target) :steps (inc steps)}) combos)) valid-levels))]
    (filter #(and (is-facility-safe %) (< steps (get visited (hash-facility %) 100))) all-moves)))

(defn explore-facility [facility]
  (loop [queue [facility] visited {} min-steps 100]
    (if (= 0 (count queue)) min-steps
      (let [current-facility (first queue) next-queue (drop 1 queue)]

        (if (or (>= (get current-facility :steps) min-steps)
                (>= (get current-facility :steps) (get visited (hash-facility current-facility) min-steps))) (recur next-queue visited min-steps)

          (if (ready-for-assembly current-facility)
            (if (< (get current-facility :steps) min-steps)
              (recur next-queue visited (get current-facility :steps))
              (recur next-queue visited min-steps))

            (let [next-visited (assoc visited (hash-facility current-facility) (get current-facility :steps))]
              (recur (concat queue (get-valid-moves current-facility visited)) next-visited min-steps))))))))

(println (format "The number of steps required to assemble some of the components is %d." (explore-facility facility)))
(println (format "The number of steps required to assemble all of the components is %d." (explore-facility facility-plus)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

