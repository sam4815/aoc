import Foundation

let start = Date()

struct Path {
  var point: (Int, Int)
  var steps: Int
  var grid: [[Substring]]
}

let grid = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  $0.split(separator: "")
}

let begin = (0, 1)
let end = (grid.count - 1, grid.count - 2)

func getNextPoints(_ point: (Int, Int)) -> [(Int, Int)] {
  [(point.0 - 1, point.1), (point.0, point.1 - 1), (point.0 + 1, point.1), (point.0, point.1 + 1)]
    .filter { $0.0 >= 0 && $0.1 >= 0 && $0.0 < grid.count && $0.1 < grid.count }
}

func getNextPaths(_ path: Path, followSlopes: Bool) -> [Path] {
  var paths: [Path] = []

  for point in getNextPoints(path.point) {
    var nextPoint = (point.0, point.1)
    var nextCost = path.steps + 1
    var nextGrid = path.grid

    if followSlopes {
      while [">", "<", "v", "^"].contains(nextGrid[nextPoint.0][nextPoint.1]) {
        let cell = nextGrid[nextPoint.0][nextPoint.1]
        nextGrid[nextPoint.0][nextPoint.1] = "O"
        nextCost += 1

        switch cell {
        case ">": nextPoint.1 += 1
        case "<": nextPoint.1 -= 1
        case "v": nextPoint.0 += 1
        case "^": nextPoint.0 -= 1
        default: ()
        }
      }
    }

    if nextGrid[nextPoint.0][nextPoint.1] == "#" || nextGrid[nextPoint.0][nextPoint.1] == "O" {
      continue
    }
    nextGrid[nextPoint.0][nextPoint.1] = "O"

    paths += [Path(point: nextPoint, steps: nextCost, grid: nextGrid)]
  }

  if paths.count == 1 && paths[0].point != end {
    return getNextPaths(paths[0], followSlopes: followSlopes)
  }

  return paths
}

func findLongestHike(_ grid: [[Substring]], followSlopes: Bool) -> Int {
  var max: Int = 0
  var visited: [Int: [Int: Int]] = [:]
  var queue: [Path] = [Path(point: begin, steps: 0, grid: grid)]

  while queue.count > 0 {
    let path = queue.removeFirst()

    if path.point == end {
      if path.steps >= max { max = path.steps }
      continue
    }

    if visited[path.point.0]?[path.point.1] ?? 0 > (path.steps + 65) {
      continue
    } else {
      visited[path.point.0, default: [:]][path.point.1] = path.steps
    }

    queue += getNextPaths(path, followSlopes: followSlopes)
  }

  return max
}

let longestHikeWithSlopes = findLongestHike(grid, followSlopes: true)
let longestHikeWithoutSlopes = findLongestHike(grid, followSlopes: false)

print(
  """
  With slopes, the longest possible hike takes \(longestHikeWithSlopes) steps.
  Without slopes, the longest possible hike takes \(longestHikeWithoutSlopes) steps.
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
