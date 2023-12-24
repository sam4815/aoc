import Foundation

let start = Date()

struct Point {
  var x: Int
  var y: Int
  var z: Int
}

struct Brick {
  var key: Int
  var start: Point
  var end: Point
  var supports: [Int]
  var dependsOn: [Int]
}

var bricks = try String(contentsOfFile: "input.txt").split(separator: "\n").enumerated().reduce(
  into: [Int: Brick]()
) { bricks, line in
  let points = line.element.split(separator: "~").map { $0.split(separator: ",").map { Int($0)! } }

  bricks[line.offset] = Brick(
    key: line.offset,
    start: Point(x: points[0][0], y: points[0][1], z: points[0][2]),
    end: Point(x: points[1][0], y: points[1][1], z: points[1][2]),
    supports: [],
    dependsOn: []
  )
}

var grid = bricks.reduce(into: [Int: [Int: [Int: Int]]]()) { grid, brick in
  for x in brick.value.start.x...brick.value.end.x {
    for y in brick.value.start.y...brick.value.end.y {
      for z in brick.value.start.z...brick.value.end.z {
        grid[x, default: [:]][y, default: [:]][z] = brick.key
      }
    }
  }
}

func findPoints(_ brick: Brick) -> [Point] {
  var points: [Point] = []
  for x in brick.start.x...brick.end.x {
    for y in brick.start.y...brick.end.y {
      for z in brick.start.z...brick.end.z {
        points += [Point(x: x, y: y, z: z)]
      }
    }
  }
  return points
}

func oneRowDown(_ points: [Point]) -> [Point] {
  points.map { Point(x: $0.x, y: $0.y, z: $0.z - 1) }
}

func isValid(_ points: [Point], key: Int) -> Bool {
  points.allSatisfy {
    let cell = grid[$0.x]?[$0.y]?[$0.z] ?? nil
    return (cell == nil || cell == key) && $0.z >= 1
  }
}

func findBricks(_ points: [Point], key: Int) -> [Int] {
  Array(Set(points.map { grid[$0.x]?[$0.y]?[$0.z] ?? -1 }.filter { $0 != key && $0 != -1 }))
}

func addPoint(_ point: Point, value: Int) {
  grid[point.x, default: [:]][point.y, default: [:]][point.z] = value
}

func removePoint(_ point: Point) {
  grid[point.x]![point.y]!.removeValue(forKey: point.z)

  if grid[point.x]![point.y]!.count == 0 {
    grid[point.x]!.removeValue(forKey: point.y)

    if grid[point.x]!.count == 0 {
      grid.removeValue(forKey: point.x)
    }
  }
}

for key in bricks.keys.sorted(by: { bricks[$0]!.start.z < bricks[$1]!.start.z }) {
  var points = findPoints(bricks[key]!)
  var next = oneRowDown(points)

  while isValid(next, key: key) {
    points.forEach({ removePoint($0) })
    next.forEach({ addPoint($0, value: key) })
    points = next
    next = oneRowDown(points)
  }

  let dependents = findBricks(next, key: key)
  for d in dependents {
    bricks[d]!.supports += [key]
  }

  bricks[key] = Brick(
    key: key,
    start: points.first!,
    end: points.last!,
    supports: [],
    dependsOn: dependents
  )
}

var numDisintegratable = 0
var brickfall = 0

for (key, brick) in bricks {
  var fallen: [Int: Bool] = [key: true]
  var queue = brick.supports

  while queue.count > 0 {
    let curr = bricks[queue.removeFirst()]!
    if curr.dependsOn.count > 0 && curr.dependsOn.allSatisfy({ fallen[$0] != nil }) {
      fallen[curr.key] = true
      queue += curr.supports
    }
  }

  brickfall += fallen.count - 1

  if brick.supports.allSatisfy({ bricks[$0]!.dependsOn.count > 1 }) {
    numDisintegratable += 1
  }
}

print(
  """
  The number of bricks that can be disintegrated is \(numDisintegratable).
  The sum of all the bricks that would fall if any one is removed is \(brickfall).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
