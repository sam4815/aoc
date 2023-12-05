import Foundation

let start = Date()

let contents = try String(contentsOfFile: "input.txt").split(separator: "\n\n")
let seeds = contents[0].split(separator: ": ")[1].split(separator: " ").map { Int($0)! }

let maps = contents[1...].map {
  $0.split(separator: "\n")[1...].map({
    $0.split(separator: " ").map { Int($0)! }
  })
}

func mapRanges(ranges: [(Int, Int)], rules: [[Int]]) -> [(Int, Int)] {
  var currRanges: [(Int, Int)] = ranges
  var nextRanges: [(Int, Int)] = []

  outer: while currRanges.count > 0 {
    let (min, max) = currRanges.removeFirst()

    for rule in rules {
      // Rule bounds the range
      if min >= rule[1] && max < (rule[1] + rule[2]) {
        nextRanges.append((rule[0] + (min - rule[1]), rule[0] + (max - rule[1])))
        continue outer
      }
      // Range bounds the rule
      if min < rule[1] && max >= (rule[1] + rule[2]) {
        currRanges.append((min, rule[1] - 1))
        currRanges.append((rule[1] + rule[2], max))

        nextRanges.append((rule[0], rule[0] + rule[2]))
        continue outer
      }
      // Range partially intersects left side of rule
      if min < rule[1] && max >= rule[1] && max < (rule[1] + rule[2]) {
        currRanges.append((min, rule[1] - 1))

        nextRanges.append((rule[0], rule[0] + (max - rule[1])))
        continue outer
      }
      // Range partially intersects right side of rule
      if min > rule[1] && min < (rule[1] + rule[2]) && max >= (rule[1] + rule[2]) {
        currRanges.append((rule[1] + rule[2], max))

        nextRanges.append((rule[0] + (min - rule[1]), rule[0] + rule[2]))
        continue outer
      }
    }

    // Range does not intersect rule at all
    nextRanges.append((min, max))
  }

  return nextRanges
}

let partialSeeds: [(Int, Int)] = stride(from: 0, to: seeds.count, by: 1).map {
  (seeds[$0], seeds[$0])
}
let allSeeds: [(Int, Int)] = stride(from: 0, to: seeds.count, by: 2).map {
  (seeds[$0], seeds[$0] + seeds[$0 + 1] - 1)
}

let partialRanges = maps.reduce(partialSeeds) { mapRanges(ranges: $0, rules: $1) }
let allRanges = maps.reduce(allSeeds) { mapRanges(ranges: $0, rules: $1) }

let partialLocationMin = partialRanges.map { $0.0 }.reduce(partialRanges[0].0, min)
let locationMin = allRanges.map { $0.0 }.reduce(allRanges[0].0, min)

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  Counting only the bounding seeds, the lowest location number is \(partialLocationMin).
  Counting all seeds, the lowest location number is \(locationMin).
  Solution generated in \(timeElapsed).
  """)
