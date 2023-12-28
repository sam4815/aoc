import Foundation

let start = Date()

let locationMap = try String(contentsOfFile: "input.txt").split(separator: "\n").enumerated()
  .reduce(
    into: [Int: [Int: Character]]()
  ) { locationMap, line in
    let numbers = line.element.split(separator: ", ").map { Int($0)! }
    locationMap[numbers[1], default: [:]][numbers[0]] = Character(UnicodeScalar(line.offset + 65)!)
  }

let xMin = locationMap.keys.min(by: <)!
let xMax = locationMap.keys.max(by: <)!
let yMin = locationMap.flatMap { $0.value.keys }.min(by: <)!
let yMax = locationMap.flatMap { $0.value.keys }.max(by: <)!

func findReachablePoints(_ i: Int, _ j: Int, distance: Int) -> [(Int, Int)] {
  var points: [(Int, Int)] = []
  for diff in (-distance)...distance {
    if abs(diff) != distance {
      points += [(i + diff, j + (distance - abs(diff))), (i + diff, j - (distance - abs(diff)))]
    } else {
      points += [(i + diff, j + (distance - abs(diff)))]
    }
  }

  return points
}

var area: [Int: [Int: (Character, Int)]] = [:]

func markNearestLocations(_ i: Int, _ j: Int, char: Character) {
  var distance = 0

  while distance < 300 {
    let points = findReachablePoints(i, j, distance: distance)

    for point in points {
      if area[point.0]?[point.1] == nil {
        area[point.0, default: [:]][point.1] = (char, distance)
      } else if area[point.0]![point.1]!.1 == distance {
        area[point.0, default: [:]][point.1] = (".", distance)
      } else if area[point.0]![point.1]!.1 > distance {
        area[point.0, default: [:]][point.1] = (char, distance)
      }
    }

    distance += 1
  }
}

func distanceToLocations(_ i: Int, _ j: Int, locations: [(Int, Int)]) -> Int {
  locations.reduce(0) { sum, location in
    sum + abs(i - location.0) + abs(j - location.1)
  }
}

let infiniteLocations = area.reduce(into: Set<Character>()) { locations, area in
  for (y, cell) in area.value {
    if area.key == xMin || area.key == xMax || y == yMin || y == yMax {
      locations.insert(cell.0)
    }
  }
}

let counts = area.flatMap { $0.value.values.map { $0.0 } }.reduce(into: [Character: Int]()) {
  counts, cell in
  counts[cell, default: 0] += infiniteLocations.contains(cell) ? 0 : 1
}

let largestArea = counts.values.max(by: <)!

let locations = locationMap.flatMap { row in row.value.keys.map { (row.key, $0) } }
let numSafeCoordinates = (xMin...xMax).reduce(0) { sum, x in
  (yMin...yMax).reduce(0) { ySum, y in
    ySum + (distanceToLocations(x, y, locations: locations) < 10000 ? 1 : 0)
  } + sum
}

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The size of the largest finite area is \(largestArea).
  The size of the safe region is \(numSafeCoordinates).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
