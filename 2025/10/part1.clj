(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(defn parse-target [string]
  (Integer/parseInt (str/join (map #(case % \[ "" \] "" \. "0" \# "1") string)) 2))

(defn parse-buttons [strings]
  (let [length (- (count (first strings)) 2) init (repeat length "0")]
    (map #(Integer/parseInt (str/join %) 2)
         (map (fn [numbers]
                (map-indexed (fn [i digit]
                               (if (some #{i} (map parse-long (re-seq #"\d+" numbers))) "1" "0")) init))
              (rest strings)))))

(def machines (->> (str/split-lines (slurp "input.txt"))
                   (map #(str/split % #" "))
                   (map #(do {:target (parse-target (first %))
                              :bits (- (count (first %)) 2)
                              :buttons (parse-buttons (drop-last %))}))))

(defn get-permutations [n limit]
  (let [ranged (mapv #(do [%]) (range limit))]
    (if (= n 1) ranged (for [x ranged y (get-permutations (dec n) limit)] (vec (concat x y))))))

(defn sum-config [buttons numbers mask]
  (reduce (fn [total [i number]]
            (bit-xor total (if (even? number) 0 (nth buttons i))))
          0
          (map-indexed vector numbers)))

(defn find-config [{:keys [target buttons bits]}]
  (let [mask (Integer/parseInt (str/join (repeat bits "1")) 2)]
    (->> (get-permutations (count buttons) 2)
         (map (fn [numbers] [(sum-config buttons numbers mask) (reduce + numbers) numbers]))
         (filter (fn [[sum num-buttons]] (= sum target)))
         (apply min-key second)
         second)))

(def solutions (map find-config machines))

(println (format "The password is %d." (reduce + solutions)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

