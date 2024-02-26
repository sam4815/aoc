(ns main
  (:require [clojure.string :as str]))

(def start-time (System/currentTimeMillis))

(def ip-addresses (str/split (slurp "input.txt") #"\n"))

(defn supports-tls [ip]
  (and (nil? (re-find #"\[\w*(\w)(\w)(?!\1)\2\1\w*\]" ip))
       (some? (re-find #"(\w)(\w)(?!\1)\2\1" ip))))

(defn supports-ssl [ip]
  (or (some? (re-find #"\[\w*(\w)(\w)\1.*\]\w*\2(?!\2)\1\2" ip))
      (some? (re-find #"(?<=\]|^)\w*(\w)(\w)\1.*\[\w*\2(?!\2)\1\2" ip))))

(println (format "%d IPs support TLS." (count (filter supports-tls ip-addresses))))
(println (format "%d IPs support SSL." (count (filter supports-ssl ip-addresses))))
(println (format "Solution generated in %.3fs." (float (/ (- (System/currentTimeMillis) start-time) 1000))))

