import Foundation

let start = Date()

enum Direction { case Left, Right }

let area = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  $0.split(separator: "").map { String($0) }
}

func findAdjacentAcres(_ position: (Int, Int), area: [[String]]) -> [String: Int] {
  var acres: [String: Int] = ["#": 0, ".": 0, "|": 0]

  for x in (position.0 - 1)...(position.0 + 1) {
    for y in (position.1 - 1)...(position.1 + 1) {
      if x < 0 || y < 0 || x >= area.count || y >= area[0].count { continue }
      if x == position.0 && y == position.1 { continue }

      acres[area[x][y]]! += 1
    }
  }

  return acres
}

func step(_ area: [[String]]) -> [[String]] {
  var nextArea: [[String]] = []
  for (x, row) in area.enumerated() {
    var nextRow: [String] = []

    for (y, cell) in row.enumerated() {
      let adjacent = findAdjacentAcres((x, y), area: area)

      if cell == "." {
        nextRow += [adjacent["|"]! >= 3 ? "|" : "."]
      }
      if cell == "|" {
        nextRow += [adjacent["#"]! >= 3 ? "#" : "|"]
      }
      if cell == "#" {
        nextRow += [(adjacent["|"]! >= 1 && adjacent["#"]! >= 1) ? "#" : "."]
      }
    }

    nextArea += [nextRow]
  }

  return nextArea
}

func findLumberValue(_ area: [[String]]) -> Int {
  let numLumb = area.reduce(0) { sum, row in sum + row.reduce(0) { $0 + ($1 == "#" ? 1 : 0) } }
  let numTrees = area.reduce(0) { sum, row in sum + row.reduce(0) { $0 + ($1 == "|" ? 1 : 0) } }
  return numLumb * numTrees
}

func findLumberValues(_ initialArea: [[String]], numMinutes: Int) -> [Int] {
  var area = initialArea
  var lumberValues: [Int] = [findLumberValue(initialArea)]

  var minute = 1
  while minute <= numMinutes {
    area = step(area)
    minute += 1
    lumberValues += [findLumberValue(area)]
  }

  return lumberValues
}

func detectCycle(_ ints: [Int]) -> (Int, Int) {
  var fast = 0
  var slow = 0

  while true {
    fast = (fast + 2) % ints.count
    slow = (slow + 1) % ints.count
    if ints[fast] == ints[slow] && fast != slow {
      return (fast, fast - slow)
    }
  }
}

let lumberValues = findLumberValues(area, numMinutes: 750)
let lumberValue10 = lumberValues[10]

let (offset, cycleLength) = detectCycle(Array(lumberValues[500...]))
let targetIndex = (1_000_000_000 - 500 - offset) % cycleLength
let lumberValue1000000000 = lumberValues[500 + targetIndex]

print(
  """
  After 10 minutes, the total resource value of the lumber is \(lumberValue10).
  After 1000000000 minutes, the total resource value of the lumber is \(lumberValue1000000000).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
