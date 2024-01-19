import Foundation

let start = Date()

struct Point { var x, y, z, w: Int, id: String }

var points = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  let numbers = $0.components(
    separatedBy: CharacterSet.decimalDigits.union(CharacterSet(charactersIn: "-")).inverted
  ).filter { !$0.isEmpty }.map { Int($0)! }

  return Point(x: numbers[0], y: numbers[1], z: numbers[2], w: numbers[3], id: String($0))
}

func findManhattanDistance(a: Point, b: Point) -> Int {
  abs(a.x - b.x) + abs(a.y - b.y) + abs(a.z - b.z) + abs(a.w - b.w)
}

func findConstellation(_ initialPoint: Point) -> Set<String> {
  var constellation = Set([initialPoint.id])
  var queue = [initialPoint]

  while queue.count > 0 {
    let point = queue.removeFirst()

    for anotherPoint in points {
      let manhattan = findManhattanDistance(a: point, b: anotherPoint)

      if manhattan <= 3 && !constellation.contains(anotherPoint.id) {
        constellation.insert(anotherPoint.id)
        queue += [anotherPoint]
      }
    }
  }

  return constellation
}

let constellations = points.reduce(into: [Set<String>]()) { constellations, point in
  if !constellations.contains(where: { $0.contains(point.id) }) {
    constellations += [findConstellation(point)]
  }
}

print(
  """
  The fixed points in spacetime form \(constellations.count) constellations.
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
