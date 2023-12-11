import Foundation

let start = Date()

let grid = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  $0.split(separator: "")
}

var occupiedColumns: [Int: Bool] = [:]
var occupiedRows: [Int: Bool] = [:]

let galaxies = grid.enumerated().reduce(into: [(Int, Int)]()) { galaxies, row in
  row.element.enumerated().forEach { (index, cell) in
    if cell == "#" {
      occupiedColumns[index] = true
      occupiedRows[row.offset] = true
      galaxies += [(row.offset, index)]
    }
  }
}

func findShortestPath(galaxyA: (Int, Int), galaxyB: (Int, Int), emptyLength: Int) -> Int {
  var xDistance = abs(galaxyA.0 - galaxyB.0)
  var yDistance = abs(galaxyA.1 - galaxyB.1)

  for i in (min(galaxyA.0, galaxyB.0)..<max(galaxyA.0, galaxyB.0)) {
    if occupiedRows[i] != true { xDistance += (emptyLength - 1) }
  }
  for i in (min(galaxyA.1, galaxyB.1)..<max(galaxyA.1, galaxyB.1)) {
    if occupiedColumns[i] != true { yDistance += (emptyLength - 1) }
  }

  return xDistance + yDistance
}

let youngSumShortestPaths = galaxies.enumerated().reduce(0) { sum, galaxyA in
  sum
    + galaxies[(galaxyA.offset + 1)...].reduce(0) { subsum, galaxyB in
      subsum + findShortestPath(galaxyA: galaxyA.element, galaxyB: galaxyB, emptyLength: 2)
    }
}

let oldSumShortestPaths = galaxies.enumerated().reduce(0) { sum, galaxyA in
  sum
    + galaxies[(galaxyA.offset + 1)...].reduce(0) { subsum, galaxyB in
      subsum + findShortestPath(galaxyA: galaxyA.element, galaxyB: galaxyB, emptyLength: 1_000_000)
    }
}

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The sum of the shortest path between every pair of younger galaxies is \(youngSumShortestPaths).
  The sum of the shortest path between every pair of older galaxies is \(oldSumShortestPaths).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
