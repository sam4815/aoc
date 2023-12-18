import Foundation

let start = Date()

enum Direction {
  case up, down, left, right
}

struct Path {
  var point: (Int, Int)
  var direction: Direction = .right
  var cost: Int = 0
}

let grid = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  $0.split(separator: "").map { Int($0)! }
}

let gridSize = grid.count

func getNextLocations(_ path: Path, range: ClosedRange<Int>) -> [((Int, Int), Int)] {
  var locations: [((Int, Int), Int)] = []
  var cost = path.cost

  for i in (1...range.last!) {
    let next: (Int, Int)
    switch path.direction {
    case .up: next = (path.point.0 - i, path.point.1)
    case .left: next = (path.point.0, path.point.1 - i)
    case .down: next = (path.point.0 + i, path.point.1)
    case .right: next = (path.point.0, path.point.1 + i)
    }

    if next.0 < 0 || next.1 < 0 || next.0 >= gridSize || next.1 >= gridSize {
      break
    }

    cost += grid[next.0][next.1]
    if i >= range.first! {
      locations += [(next, cost)]
    }
  }

  return locations
}

func findMinimumCostPath(_ grid: [[Int]], range: ClosedRange<Int>) -> Int {
  var min = 9 * (gridSize * 2)
  var visited: [Int: [Int: [Direction: Int]]] = [:]
  var queue: [Path] = [Path(point: (0, 0)), Path(point: (0, 0), direction: .down)]

  while queue.count > 0 {
    let path = queue.removeFirst()

    if path.point.0 == (gridSize - 1) && path.point.1 == (gridSize - 1) {
      if path.cost < min {
        min = path.cost
      }
      continue
    }

    if visited[path.point.0]?[path.point.1]?[path.direction] ?? Int.max <= path.cost {
      continue
    } else {
      visited[path.point.0, default: [:]][path.point.1, default: [:]][path.direction] =
        path.cost
    }

    if (path.cost) >= min {
      continue
    }

    for (nextPoint, nextCost) in getNextLocations(path, range: range) {
      switch path.direction {
      case .up, .down:
        queue.insert(
          contentsOf: [
            Path(point: nextPoint, direction: .right, cost: nextCost),
            Path(point: nextPoint, direction: .left, cost: nextCost),
          ],
          at: 0)
      case .left, .right:
        queue.insert(
          contentsOf: [
            Path(point: nextPoint, direction: .down, cost: nextCost),
            Path(point: nextPoint, direction: .up, cost: nextCost),
          ],
          at: 0)
      }
    }
  }

  return min
}

let minHeatLoss = findMinimumCostPath(grid, range: (1...3))
let minUltraHeatLoss = findMinimumCostPath(grid, range: (4...10))

print(
  """
  The minimum heat loss that can be incurred with the regular crucible is \(minHeatLoss).
  The minimum heat loss that can be incurred with the ultra crucible is \(minUltraHeatLoss).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
