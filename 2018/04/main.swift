import Foundation

let start = Date()

let records: [(Date, String)] = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  record in
  let parts = record.split(separator: "] ").map { String($0) }

  let dateFormatter = DateFormatter()
  dateFormatter.dateFormat = "[yyyy-MM-dd HH:mm"

  let date = dateFormatter.date(from: parts[0])!

  return (date, parts[1])
}

var activeGuardId = 0
var lastMinute = 0

let guardMap = records.sorted(by: { $0.0 < $1.0 }).reduce(into: [Int: [Int: Int]]()) {
  guards, record in
  if record.1.contains("Guard") {
    activeGuardId = Int(
      record.1.components(separatedBy: CharacterSet.decimalDigits.inverted).joined())!
  }

  if record.1.contains("falls asleep") {
    lastMinute = Calendar.current.component(.minute, from: record.0)
  }

  if record.1.contains("wakes up") {
    let minute = Calendar.current.component(.minute, from: record.0)
    for i in lastMinute..<minute {
      guards[activeGuardId, default: [:]][i, default: 0] += 1
    }
  }
}

let sleepiestGuard = guardMap.max(by: { a, b in
  a.value.values.reduce(0, +) < b.value.values.reduce(0, +)
})!
let sleepiestMinute = sleepiestGuard.value.max(by: { $0.value < $1.value })!
let sleepiestProduct = sleepiestGuard.key * sleepiestMinute.key

let secondSleepiestGuard = guardMap.max(by: { a, b in
  a.value.values.max(by: <)! < b.value.values.max(by: <)!
})!
let secondSleepiestMinute = secondSleepiestGuard.value.max(by: { $0.value < $1.value })!
let secondSleepiestProduct = secondSleepiestGuard.key * secondSleepiestMinute.key

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The sleepiest guard ID multiplied by the sleepiest minute is \(sleepiestProduct).
  The second sleepiest guard ID multiplied by the second sleepiest minute is \(secondSleepiestProduct).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
