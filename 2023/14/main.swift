import Foundation

enum Direction {
  case up, down, left, right
}

let start = Date()

var platform = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  $0.split(separator: "")
}

func findNorthLoad(_ platform: [[Substring]]) -> Int {
  platform.enumerated().reduce(0) { sum, row in
    sum + row.element.reduce(0) { $0 + ($1 == "O" ? platform.count - row.offset : 0) }
  }
}

func moveCharacters(_ platform: inout [[Substring]], direction: Direction) {
  let xRange = AnySequence(0..<platform.count)
  let yRange = (direction == .down || direction == .right) ? AnySequence(xRange.reversed()) : xRange
  let step = (direction == .down || direction == .right) ? -1 : 1

  for x in xRange {
    var openIndex = (direction == .down || direction == .right) ? platform.count - 1 : 0

    for y in yRange {
      let cell = (direction == .left || direction == .right) ? platform[x][y] : platform[y][x]

      if cell == "O" {
        switch direction {
        case .up, .down:
          platform[y][x] = "."
          platform[openIndex][x] = "O"
        case .left, .right:
          platform[x][y] = "."
          platform[x][openIndex] = "O"
        }
      }

      openIndex = cell == "O" ? (openIndex + step) : (cell == "#" ? (y + step) : openIndex)
    }
  }
}

func cycle(_ platform: [[Substring]], numCycles: Int) -> [Int] {
  var next = platform
  var northLoads: [Int] = []

  for _ in 0..<numCycles {
    moveCharacters(&next, direction: .up)
    moveCharacters(&next, direction: .left)
    moveCharacters(&next, direction: .down)
    moveCharacters(&next, direction: .right)

    northLoads += [findNorthLoad(next)]
  }

  return northLoads
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

moveCharacters(&platform, direction: .up)
let initialNorthLoad = findNorthLoad(platform)

let northLoads = cycle(platform, numCycles: 250)
let (offset, cycleLength) = detectCycle(Array(northLoads[200...]))
let targetIndex = (1_000_000_000 - 200 - offset) % cycleLength
let finalNorthLoad = northLoads[200 + targetIndex - 1]

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  After tilting the platform north, the load on the north support beams is \(initialNorthLoad).
  After 1000000000 cycles, the load on the north support beams is \(finalNorthLoad).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
