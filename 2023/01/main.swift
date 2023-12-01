import Foundation
import SwiftUI

let start = NSDate()

let contents = try? String(contentsOfFile: "input.txt")
let lines = contents!.split(separator: "\n")

let digitNames = [
  "zero": "0", "one": "1", "two": "2", "three": "3", "four": "4",
  "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
]

let digitsRegex = try? NSRegularExpression(pattern: "(?=(\\d))")
let digitsLettersRegex = try? NSRegularExpression(
  pattern: "(?=(one|two|three|four|five|six|seven|eight|nine|\\d))")

func findCalibrationValue(line: Substring, pattern: NSRegularExpression) -> Int {
  let digits = pattern.matches(
    in: String(line),
    range: NSRange(line.startIndex..., in: line))

  let firstMatch = String(line[Range(digits.first!.range(at: 1), in: line)!])
  let lastMatch = String(line[Range(digits.last!.range(at: 1), in: line)!])

  let firstDigit = digitNames[firstMatch] ?? firstMatch
  let lastDigit = digitNames[lastMatch] ?? lastMatch

  return Int(firstDigit + lastDigit)!
}

var calibrationSumDigits = 0
var calibrationSumLetters = 0

for line in lines {
  calibrationSumDigits += findCalibrationValue(line: line, pattern: digitsRegex!)
  calibrationSumLetters += findCalibrationValue(line: line, pattern: digitsLettersRegex!)
}

let timeElapsed = NSDate().timeIntervalSince(start as Date)

print(
  """
  Counting only digits, the sum of the calibration values is \(calibrationSumDigits).
  Counting both digits and letters, the sum of the calibration values is \(calibrationSumLetters).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
