import Foundation

let start = Date()

let sequences = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  $0.split(separator: " ").map { Int($0)! }
}

func deriveSequence(sequence: [Int]) -> [Int] {
  var nextSequence: [Int] = []
  for i in 1..<sequence.count {
    nextSequence += [sequence[i] - sequence[i - 1]]
  }
  return nextSequence
}

func findExtrapolatedBounds(sequence: [Int]) -> (Int, Int) {
  var derivatives: [[Int]] = [sequence]

  while derivatives.last!.contains(where: { $0 != 0 }) {
    derivatives += [deriveSequence(sequence: derivatives.last!)]
  }

  return (
    derivatives.reversed().reduce(0) { $1.first! - $0 },
    derivatives.reduce(0) { $0 + $1.last! }
  )
}

let extrapolatedBounds = sequences.map { findExtrapolatedBounds(sequence: $0) }

let nextExtrapolatedSum = extrapolatedBounds.reduce(0) { $0 + $1.1 }
let prevExtrapolatedSum = extrapolatedBounds.reduce(0) { $0 + $1.0 }

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The sum of the next extrapolated values is \(nextExtrapolatedSum).
  The sum of the previous extrapolated values is \(prevExtrapolatedSum).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
