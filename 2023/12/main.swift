import Foundation

let start = Date()

let springs = try String(contentsOfFile: "input.txt").split(separator: "\n").map { line in
  let components = line.split(separator: " ")
  return (components[0].split(separator: ""), components[1].split(separator: ",").map { Int($0)! })
}

let unfoldedSprings = springs.map { spring in
  (
    Array([[Substring]](repeating: spring.0 + ["?"], count: 5).flatMap { $0 }.dropLast()),
    [[Int]](repeating: spring.1, count: 5).flatMap { $0 }
  )
}

var visited = [String: Int]()

func countPermutations(records: [Substring], groups: [Int]) -> Int {
  let hash = records.joined() + groups.map { String($0) }.joined(separator: ",")
  if visited[hash] != nil {
    return visited[hash]!
  }

  let group = groups.first!
  var sum = 0

  for index in 0..<records.count {
    if records[..<index].contains("#") || (group + index) > records.count {
      break
    }

    if !records[index..<(group + index)].allSatisfy({ $0 == "#" || $0 == "?" }) {
      continue
    }

    if records.count == group + index {
      sum += groups.count == 1 ? 1 : 0
      continue
    }

    if !["?", "."].contains(records[group + index]) {
      continue
    }

    if groups.count == 1 && !records[(group + index)...].contains("#") {
      sum += 1
    }

    if groups.count > 1 {
      sum += countPermutations(
        records: Array(records[(group + index + 1)...]), groups: Array(groups[1...])
      )
    }
  }

  visited[hash] = sum
  return sum
}

let sumArrangements = springs.reduce(0) { $0 + countPermutations(records: $1.0, groups: $1.1) }
let unfoldedSumArrangements = unfoldedSprings.reduce(0) {
  $0 + countPermutations(records: $1.0, groups: $1.1)
}

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The sum of all possible spring arrangements is \(sumArrangements).
  The sum of all possible unfolded spring arrangements is \(unfoldedSumArrangements).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
