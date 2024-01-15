import Foundation

let start = Date()

enum RegionType { case Rocky, Wet, Narrow }
enum Tool { case Gear, Torch, Neither }

struct Region {
  var type: RegionType
  var coord: (Int, Int)
}

struct Path {
  var region: Region
  var equipped: Tool
  var minutes: Int
}

let numbers = try String(contentsOfFile: "input.txt").components(
  separatedBy: CharacterSet.decimalDigits.inverted
).filter({ !$0.isEmpty }).map({ Int($0)! })

let depth = numbers[0]
let target = (numbers[1], numbers[2])

func findRegionType(_ region: (Int, Int)) -> RegionType {
  switch findRiskLevel(region) {
  case 0: return .Rocky
  case 1: return .Wet
  default: return .Narrow
  }
}

func findRiskLevel(_ region: (Int, Int)) -> Int {
  findErosionLevel(region) % 3
}

func findErosionLevel(_ region: (Int, Int)) -> Int {
  (findGeologicalIndex(region) + depth) % 20183
}

var geoIndices: [Int: [Int: Int]] = [:]
func findGeologicalIndex(_ region: (Int, Int)) -> Int {
  if geoIndices[region.0]?[region.1] ?? -1 != -1 { return geoIndices[region.0]![region.1]! }

  let index: Int
  switch region {
  case (0, 0): index = 0
  case (target.0, target.1): index = 0
  case (_, 0): index = region.0 * 16807
  case (0, _): index = region.1 * 48271
  case (_, _):
    index = findErosionLevel((region.0 - 1, region.1)) * findErosionLevel((region.0, region.1 - 1))
  }

  geoIndices[region.0, default: [:]][region.1] = index
  return index
}

let riskLevelSum = (0...target.0).reduce(0) { sum, x in
  sum + (0...target.1).reduce(0) { sum, y in sum + findRiskLevel((x, y)) }
}

func findAdjacentRegions(_ region: Region) -> [Region] {
  var adjacent: [Region] = []

  for adj in [
    (region.coord.0 + 1, region.coord.1), (region.coord.0 - 1, region.coord.1),
    (region.coord.0, region.coord.1 + 1), (region.coord.0, region.coord.1 - 1),
  ] {
    if adj.0 < 0 || adj.1 < 0 { continue }
    adjacent += [Region(type: findRegionType((adj.0, adj.1)), coord: adj)]
  }

  return adjacent
}

func distanceToTarget(_ region: Region) -> Int {
  abs(region.coord.0 - target.0) + abs(region.coord.1 - target.1)
}

func getToolOptions(_ path: Path) -> [Tool] {
  switch path.region.type {
  case .Rocky: return [.Torch, .Gear].filter { $0 != path.equipped }
  case .Wet: return [.Gear, .Neither].filter { $0 != path.equipped }
  case .Narrow: return [.Torch, .Neither].filter { $0 != path.equipped }
  }
}

func hasVisited(_ path: Path, visited: inout [Int: [Int: [Tool: Int]]]) -> Bool {
  if visited[path.region.coord.0]?[path.region.coord.1]?[path.equipped] ?? Int.max <= path.minutes {
    return true
  } else {
    visited[path.region.coord.0, default: [:]][path.region.coord.1, default: [:]][path.equipped] =
      path.minutes
    return false
  }
}

func insertPath(_ path: Path, queue: inout [Path]) {
  let minDuration = distanceToTarget(path.region) + path.minutes

  for i in 0..<queue.count {
    if minDuration < distanceToTarget(queue[i].region) + queue[i].minutes {
      queue.insert(path, at: i)
      return
    }
  }

  queue += [path]
}

func findShortestPath(_ target: (Int, Int)) -> Int {
  var minMinutes = Int.max
  var queue = [Path(region: Region(type: .Rocky, coord: (0, 0)), equipped: .Torch, minutes: 0)]
  var visited: [Int: [Int: [Tool: Int]]] = [:]

  while queue.count > 0 {
    var path = queue.removeFirst()

    if path.region.coord.0 == target.0 && path.region.coord.1 == target.1 {
      if path.equipped == .Gear { path.minutes += 7 }

      minMinutes = min(minMinutes, path.minutes)
      continue
    }

    if distanceToTarget(path.region) + path.minutes > minMinutes { continue }
    if hasVisited(path, visited: &visited) { continue }

    for region in findAdjacentRegions(path.region) {
      switch (path.equipped, region.type) {
      case (.Torch, .Rocky), (.Torch, .Narrow):
        insertPath(Path(region: region, equipped: .Torch, minutes: path.minutes + 1), queue: &queue)
      case (.Torch, .Wet):
        for tool in getToolOptions(path) {
          insertPath(Path(region: region, equipped: tool, minutes: path.minutes + 8), queue: &queue)
        }

      case (.Gear, .Rocky), (.Gear, .Wet):
        insertPath(Path(region: region, equipped: .Gear, minutes: path.minutes + 1), queue: &queue)
      case (.Gear, .Narrow):
        for tool in getToolOptions(path) {
          insertPath(Path(region: region, equipped: tool, minutes: path.minutes + 8), queue: &queue)
        }

      case (.Neither, .Wet), (.Neither, .Narrow):
        insertPath(
          Path(region: region, equipped: .Neither, minutes: path.minutes + 1), queue: &queue)
      case (.Neither, .Rocky):
        for tool in getToolOptions(path) {
          insertPath(Path(region: region, equipped: tool, minutes: path.minutes + 8), queue: &queue)
        }
      }
    }
  }

  return minMinutes
}

let numMinutes = findShortestPath(target)

print(
  """
  The sum of the risk levels is \(riskLevelSum).
  The shortest path to the target takes \(numMinutes) minutes.
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
