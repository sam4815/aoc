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
  ranges.flatMap { (min, max) in
    for rule in rules {
      // Rule bounds the range
      if min >= rule[1] && max < (rule[1] + rule[2]) {
        return [(rule[0] + (min - rule[1]), rule[0] + (max - rule[1]))]
      }
      // Range bounds the rule
      if min < rule[1] && max >= (rule[1] + rule[2]) {
        return mapRanges(ranges: [(min, rule[1] - 1)], rules: rules)
          + mapRanges(ranges: [(rule[1] + rule[2], max)], rules: rules) + [
            (rule[0], rule[0] + rule[2])
          ]
      }
      // Range partially intersects left side of rule
      if min < rule[1] && max >= rule[1] && max < (rule[1] + rule[2]) {
        return mapRanges(ranges: [(min, rule[1] - 1)], rules: rules) + [
          (rule[0], rule[0] + (max - rule[1]))
        ]
      }
      // Range partially intersects right side of rule
      if min > rule[1] && min < (rule[1] + rule[2]) && max >= (rule[1] + rule[2]) {
        return mapRanges(ranges: [(rule[1] + rule[2], max)], rules: rules) + [
          (rule[0] + (min - rule[1]), rule[0] + rule[2])
        ]
      }
    }
    // Range does not intersect rule at all
    return [(min, max)]
  }
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
