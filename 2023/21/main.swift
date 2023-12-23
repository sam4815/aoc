import Foundation

let start = Date()

struct Path {
  var point: (Int, Int)
  var steps: Int
}

let grid = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  $0.split(separator: "")
}

func getOneAway(_ point: (Int, Int), grid: [[Substring]]) -> [(Int, Int)] {
  var diffs: [(Int, Int)] = []

  for diff in [(-1, 0), (0, -1), (1, 0), (0, 1)] {
    let x = (point.0 + diff.0 + grid.count) % grid.count
    let y = (point.1 + diff.1 + grid[0].count) % grid[0].count
    if grid[x][y] == "#" { continue }

    diffs += [diff]
  }

  return diffs
}

func countPossible(_ numSteps: Int) -> Int {
  return (numSteps + 1) * (numSteps + 1)
}

func isEvenCell(_ point: (Int, Int)) -> Bool {
  (point.0 % 2 == 0 && point.1 % 2 == 0) || (point.0 % 2 == 1 && point.1 % 2 == 1)
}

func isOddCell(_ point: (Int, Int)) -> Bool {
  (point.0 % 2 == 1 && point.1 % 2 == 0) || (point.0 % 2 == 0 && point.1 % 2 == 1)
}

func isBlocked(_ point: (Int, Int)) -> Bool {
  grid[point.0][point.1] == "#" || getOneAway(point, grid: grid).count == 0
}

func isWithinSquare(_ point: (Int, Int)) -> Bool {
  (abs(65 - point.0) + abs(65 - point.1)) < 65
}

func countBoulders() -> [Int] {
  var counts: [Int] = [0, 0, 0, 0]

  for x in 0..<grid.count {
    for y in 0..<grid.count {
      if !isBlocked((x, y)) { continue }
      if isEvenCell((x, y)) && isWithinSquare((x, y)) { counts[0] += 1 }
      if isEvenCell((x, y)) && !isWithinSquare((x, y)) { counts[1] += 1 }
      if isOddCell((x, y)) && isWithinSquare((x, y)) { counts[2] += 1 }
      if isOddCell((x, y)) && !isWithinSquare((x, y)) { counts[3] += 1 }
    }
  }

  return counts
}

func findSquareSum(_ target: Int) -> Int {
  var i = 1
  while i * i + (i + 1) * (i + 1) < 81_850_984_600 { i += 1 }

  return i
}

let boulderCounts = countBoulders()

func countPlots(_ steps: Int) -> Int {
  let possible = countPossible(steps)
  let squaresWidth = ((steps - 65) / 131 * 2) + 1
  let numSquares = squaresWidth * squaresWidth

  let numInSquares = (numSquares - 1) / 2 + 1
  let numOutSquares = (numSquares - 1) / 4
  let squareSum = findSquareSum(numInSquares)
  let smallerSquare = squareSum * squareSum
  let largerSquare = (squareSum + 1) * (squareSum + 1)

  return possible - (boulderCounts[0] * smallerSquare) - (boulderCounts[1] * numOutSquares)
    - (boulderCounts[2] * largerSquare) - (boulderCounts[3] * numOutSquares)
}

let numPlotsReached = countPossible(64) - boulderCounts[0]
let numExpandedPlotsReached = countPlots(26_501_365)

print(
  """
  The number of garden plots that can be reached in 64 steps is \(numPlotsReached).
  The number of plots that can be reached in 26501365 steps is \(numExpandedPlotsReached).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
