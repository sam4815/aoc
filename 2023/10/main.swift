import Foundation

let start = Date()

let grid = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  $0.split(separator: "")
}

func findStart(grid: [[Substring]]) -> (Int, Int) {
  for (x, row) in grid.enumerated() {
    for (y, cell) in row.enumerated() {
      if cell == "S" { return (x, y) }
    }
  }
  return (0, 0)
}

func findNextPositions(position: (Int, Int)) -> [(Int, Int)] {
  if position.0 < 0 || position.1 < 0 { return [] }

  let symbol =
    grid[position.0][position.1] == "S"
    ? transformS(position: position) : grid[position.0][position.1]

  return [
    "|": [(position.0 - 1, position.1), (position.0 + 1, position.1)],
    "-": [(position.0, position.1 - 1), (position.0, position.1 + 1)],
    "J": [(position.0 - 1, position.1), (position.0, position.1 - 1)],
    "L": [(position.0 - 1, position.1), (position.0, position.1 + 1)],
    "7": [(position.0, position.1 - 1), (position.0 + 1, position.1)],
    "F": [(position.0, position.1 + 1), (position.0 + 1, position.1)],
    ".": [],
  ][symbol]!
}

func transformS(position: (Int, Int)) -> Substring {
  let connectedPositions = [
    ("Up", (position.0 - 1, position.1)), ("Down", (position.0 + 1, position.1)),
    ("Left", (position.0, position.1 - 1)), ("Right", (position.0, position.1 + 1)),
  ].filter { findNextPositions(position: $0.1).contains(where: { $0 == position }) }.reduce(
    ""
  ) { $0 + $1.0 }

  return [
    "UpDown": "|",
    "LeftRight": "-",
    "UpLeft": "J",
    "UpRight": "L",
    "DownLeft": "7",
    "DownRight": "F",
  ][connectedPositions]!
}

func walkLoop(position: (Int, Int)) -> [Int: [Int: Int]] {
  var path: [Int: [Int: Int]] = [:]
  var queue: [(Int, Int)] = findNextPositions(position: position)

  while queue.count > 0 {
    let currPos = queue.removeFirst()
    let currSteps = path[currPos.0, default: [:]][currPos.1, default: 1]

    for nextPos in findNextPositions(position: currPos) {
      if currSteps < path[nextPos.0, default: [:]][nextPos.1, default: Int.max] {
        path[nextPos.0, default: [:]][nextPos.1] = currSteps + 1
        queue.append(nextPos)
      }
    }
  }

  return path
}

func getPathSymbol(x: Int, y: Int, path: [Int: [Int: Int]]) -> Substring {
  let pathOccurrences = path[x, default: [:]][y, default: -1]
  let symbol = pathOccurrences >= 0 ? grid[x][y] : "."

  if symbol == "S" {
    return transformS(position: (x, y))
  }

  return symbol
}

func countEnclosed(path: [Int: [Int: Int]]) -> Int {
  grid.enumerated().reduce(into: 0) { sum, row in
    var enclosed = false
    let x = row.offset
    var y = 0

    while y < row.element.count {
      switch getPathSymbol(x: x, y: y, path: path) {
      case ".":
        sum += enclosed ? 1 : 0

      case "|":
        enclosed = !enclosed

      case "F":
        while ["F", "-"].contains(getPathSymbol(x: x, y: y, path: path)) { y += 1 }
        enclosed = getPathSymbol(x: x, y: y, path: path) == "J" ? !enclosed : enclosed

      case "L":
        while ["L", "-"].contains(getPathSymbol(x: x, y: y, path: path)) { y += 1 }
        enclosed = getPathSymbol(x: x, y: y, path: path) == "7" ? !enclosed : enclosed

      default: ()
      }

      y += 1
    }
  }
}

let startPosition = findStart(grid: grid)
let path = walkLoop(position: startPosition)

let farthestPosition = path.values.map { $0.values.max(by: <)! }.max(by: <)!
let numEnclosed = countEnclosed(path: path)

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The number of steps to the farthest position is \(farthestPosition).
  The number of tiles enclosed in this loop is \(numEnclosed).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
