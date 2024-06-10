(ns main)

(def start-time (System/currentTimeMillis))

(def instructions (slurp "input.txt"))

(defn track-floor [instructions]
  (loop [floor 0 n 0]
    (if (>= n (count instructions)) floor
      (recur (if (= (get instructions n) \() (inc floor) (dec floor)) (inc n)))))

(defn find-first-basement [instructions]
  (loop [floor 0 n 0]
    (if (< floor 0) n
      (recur (if (= (get instructions n) \() (inc floor) (dec floor)) (inc n)))))

(println (format "The instructions take Santa to floor %d." (track-floor instructions)))
(println (format "Santa enters the basement at position %d." (find-first-basement instructions)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

