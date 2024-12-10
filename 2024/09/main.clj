(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def disk-map (mapv parse-long (re-seq #"\d" (slurp "input.txt"))))

(defn find-sum [system-pos file-id file-size]
  (reduce + (map * (range system-pos (+ system-pos file-size)) (repeat file-size file-id))))

(defn get-file [position disk-map]
  (let [id (/ position 2)]
    [id (get disk-map position 0)]))

(defn get-prev-file [id disk-map]
  (get-file (* (dec id) 2) disk-map))

(defn get-last-file [disk-map]
  (get-file (dec (count disk-map)) disk-map))

(defn find-checksum [init-disk-map]
  (loop [map-position 0 system-position 0 checksum 0
         file-under-move (get-last-file disk-map)
         disk-map init-disk-map]
    (if (> map-position (* 2 (first file-under-move)))
      checksum
      (if (even? map-position)
        (let [[file-id file-size] (get-file map-position disk-map)]
          (recur (inc map-position)
                 (+ system-position file-size)
                 (+ checksum (find-sum system-position file-id file-size))
                 file-under-move
                 disk-map))

        (let [space-size (get disk-map map-position) [move-file-id move-file-size] (identity file-under-move)]
          (if (>= space-size move-file-size)
            (recur (if (= space-size move-file-size) (inc map-position) map-position)
                   (+ system-position move-file-size)
                   (+ checksum (find-sum system-position move-file-id move-file-size))
                   (get-prev-file move-file-id disk-map)
                   (assoc disk-map map-position (- space-size move-file-size)))

            (recur (inc map-position)
                   (+ system-position space-size)
                   (+ checksum (find-sum system-position move-file-id space-size))
                   [move-file-id (- move-file-size space-size)]
                   (assoc disk-map (* 2 move-file-id) (- move-file-size space-size)))))))))

(defn find-system-position [map-position disk-map]
  (reduce + (subvec disk-map 0 map-position)))

(defn find-space [size map-position disk-map]
  (first (filter (fn [[i space]] (and (odd? i) (>= space size)))
                 (map-indexed #(list %1 %2) (subvec disk-map 0 map-position)))))

(defn remove-file [file-id file-size space-position disk-map]
  (let [file-position (* file-id 2)
        file-removed (assoc disk-map file-position 0)]
    (assoc file-removed (dec file-position) (+ file-size (get disk-map (dec file-position))))))

(defn find-new-checksum [init-disk-map]
  (loop [checksum 0
         [move-file-id move-file-size] (get-last-file disk-map)
         disk-map init-disk-map]
    (if (neg? move-file-id)
      checksum
      (let [space (find-space move-file-size (* move-file-id 2) disk-map)]
        (if (some? space)
          (let [[space-position space-size] (identity space)
                space-taken (- (get init-disk-map space-position) space-size)]
            (recur (+ checksum (find-sum (+ space-taken (find-system-position space-position init-disk-map))
                                         move-file-id move-file-size))
                   (get-prev-file move-file-id disk-map)
                   (assoc (remove-file move-file-id move-file-size space-position disk-map)
                          space-position
                          (- space-size move-file-size))))

          (recur (+ checksum (find-sum (find-system-position (* move-file-id 2) init-disk-map) move-file-id move-file-size))
                 (get-prev-file move-file-id disk-map)
                 disk-map))))))

(println (format "The filesystem checksum is %d." (find-checksum disk-map)))
(println (format "Using the new method, the filesystem checksum is %d." (find-new-checksum disk-map)))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

