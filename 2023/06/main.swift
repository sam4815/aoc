import Foundation

let start = Date()

let contents = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  $0.split(separator: ": ")[1].split(separator: " ")
}

func getNumberRecordBeaters(time: Double, distance: Double) -> Int {
  // Quadratic formula
  let b = time * -1.0
  let c = distance

  let lowerBound = (-b - sqrt(pow(b, 2) - 4 * c)) / 2
  let upperBound = (-b + sqrt(pow(b, 2) - 4 * c)) / 2

  return Int(upperBound.rounded(.towardZero) - lowerBound.rounded(.towardZero))
}

let recordProduct = contents[0].enumerated().map({ (index, time) in
  getNumberRecordBeaters(time: Double(time)!, distance: Double(contents[1][index])!)
}).reduce(1, *)

let totalRecordBeaters = getNumberRecordBeaters(
  time: Double(contents[0].joined())!, distance: Double(contents[1].joined())!)

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The product of all the ways to beat every record is \(recordProduct).
  The number of ways to beat the single record is \(totalRecordBeaters).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
