(ns main)

(def start-time (System/currentTimeMillis))

(def instructions (slurp "input.txt"))

(defn get-next-position [[x y] instruction]
  (case instruction
    \> [(inc x) y] \< [(dec x) y] \v [x (inc y)] \^ [x (dec y)]))

(defn deliver-present [[x y] houses]
  (assoc houses x (assoc (get houses x) y (inc (get (get houses x) y 0))))) 

(defn deliver-presents [instructions num-santas]
  (loop [n 0 positions (vec (repeat num-santas [0 0])) houses { 0 { 0 1 }}]
    (if (>= n (count instructions)) houses
      (let [next-position (get-next-position (get positions (mod n num-santas)) (get instructions n))]
        (recur (inc n) (assoc positions (mod n num-santas) next-position) (deliver-present next-position houses))))))

(def houses-visited (reduce + (map count (vals (deliver-presents instructions 1)))))
(def robo-houses-visited (reduce + (map count (vals (deliver-presents instructions 2)))))

(println (format "Normally, %d houses receive at least one present." houses-visited))
(println (format "With Robo-Santa, %d houses receive at least one present." robo-houses-visited))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

