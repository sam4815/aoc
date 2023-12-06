import Foundation

let start = Date()

struct Rule {
  var min: Int = 0
  var max: Int = 0
  var diff: Int = 0
}

let contents = try String(contentsOfFile: "input.txt").split(separator: "\n\n")
let seeds = contents[0].split(separator: ": ")[1].split(separator: " ").map { Int($0)! }

let maps = contents[1...].map {
  $0.split(separator: "\n")[1...].map({
    $0.split(separator: " ").map { Int($0)! }
  }).map { Rule(min: $0[1], max: $0[1] + $0[2], diff: $0[0] - $0[1]) }
}

func mapRanges(ranges: [(Int, Int)], rules: [Rule]) -> [(Int, Int)] {
  ranges.flatMap { (min, max) in
    for rule in rules {
      // Rule bounds the range
      if min >= rule.min && max < rule.max {
        return [(min + rule.diff, max + rule.diff)]
      }
      // Range bounds the rule
      if min < rule.min && max >= rule.max {
        return mapRanges(ranges: [(min, rule.min - 1), (rule.max, max)], rules: rules) + [
          (rule.min + rule.diff, rule.max + rule.diff)
        ]
      }
      // Range intersects left side of rule
      if min < rule.min && max >= rule.min && max < rule.max {
        return mapRanges(ranges: [(min, rule.min - 1)], rules: rules) + [
          (rule.min + rule.diff, max + rule.diff)
        ]
      }
      // Range intersects right side of rule
      if min > rule.min && min < rule.max && max >= rule.max {
        return mapRanges(ranges: [(rule.max, max)], rules: rules) + [
          (min + rule.diff, rule.max + rule.diff)
        ]
      }
    }
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
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
