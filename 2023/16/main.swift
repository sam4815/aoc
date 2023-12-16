import Foundation

let start = Date()

enum Direction {
  case up, down, left, right
}

struct Beam {
  var point: (Int, Int)
  var direction: Direction
}

let grid = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  $0.split(separator: "")
}

func countEnergized(_ grid: [[Substring]]) -> Int {
  grid.reduce(0) { $0 + $1.reduce(0) { $0 + ($1 == "#" ? 1 : 0) } }
}

func getNextLocation(_ beam: Beam) -> (Int, Int)? {
  let location: (Int, Int)
  switch beam.direction {
  case .up: location = (beam.point.0 - 1, beam.point.1)
  case .left: location = (beam.point.0, beam.point.1 - 1)
  case .down: location = (beam.point.0 + 1, beam.point.1)
  case .right: location = (beam.point.0, beam.point.1 + 1)
  }

  if location.0 < 0 || location.1 < 0 || location.0 >= grid.count || location.1 >= grid[0].count {
    return nil
  }

  return location
}

func traceBeam(_ grid: [[Substring]], start: Beam) -> [[Substring]] {
  var traced = grid
  var visited: [Direction: [Int: [Int: Bool]]] = [:]
  var queue: [Beam] = [start]

  while queue.count > 0 {
    let beam = queue.removeFirst()

    if visited[beam.direction]?[beam.point.0]?[beam.point.1] ?? false {
      continue
    } else {
      visited[beam.direction, default: [:]][beam.point.0, default: [:]][beam.point.1] = true
    }

    traced[beam.point.0][beam.point.1] = "#"

    switch (grid[beam.point.0][beam.point.1], beam.direction) {
    case ("|", .left), ("|", .right):
      if let upBeam = getNextLocation(Beam(point: beam.point, direction: .up)) {
        queue.append(Beam(point: upBeam, direction: .up))
      }
      if let downBeam = getNextLocation(Beam(point: beam.point, direction: .down)) {
        queue.append(Beam(point: downBeam, direction: .down))
      }

    case ("-", .up), ("-", .down):
      if let leftBeam = getNextLocation(Beam(point: beam.point, direction: .left)) {
        queue.append(Beam(point: leftBeam, direction: .left))
      }
      if let rightBeam = getNextLocation(Beam(point: beam.point, direction: .right)) {
        queue.append(Beam(point: rightBeam, direction: .right))
      }

    case ("\\", .right), ("/", .left):
      if let downBeam = getNextLocation(Beam(point: beam.point, direction: .down)) {
        queue.append(Beam(point: downBeam, direction: .down))
      }

    case ("/", .right), ("\\", .left):
      if let upBeam = getNextLocation(Beam(point: beam.point, direction: .up)) {
        queue.append(Beam(point: upBeam, direction: .up))
      }

    case ("\\", .up), ("/", .down):
      if let leftBeam = getNextLocation(Beam(point: beam.point, direction: .left)) {
        queue.append(Beam(point: leftBeam, direction: .left))
      }

    case ("/", .up), ("\\", .down):
      if let rightBeam = getNextLocation(Beam(point: beam.point, direction: .right)) {
        queue.append(Beam(point: rightBeam, direction: .right))
      }

    default:
      if let nextBeam = getNextLocation(Beam(point: beam.point, direction: beam.direction)) {
        queue.append(Beam(point: nextBeam, direction: beam.direction))
      }
    }
  }

  return traced
}

let numEnergizedTiles = countEnergized(
  traceBeam(grid, start: Beam(point: (0, 0), direction: .right)))

let maxEnergizedTiles = [Direction.up, .left, .down, .right].map { direction in
  (0..<grid.count).map { i in
    switch direction {
    case .down:
      return traceBeam(grid, start: Beam(point: (0, i), direction: .down))
    case .up:
      return traceBeam(grid, start: Beam(point: (grid.count - 1, i), direction: .up))
    case .right:
      return traceBeam(grid, start: Beam(point: (i, 0), direction: .right))
    case .left:
      return traceBeam(grid, start: Beam(point: (i, grid.count - 1), direction: .left))
    }
  }.map { countEnergized($0) }.max(by: <)!
}.max(by: <)!

print(
  """
  The number of energized tiles in \(numEnergizedTiles).
  The most tiles that can be energized is \(maxEnergizedTiles).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
