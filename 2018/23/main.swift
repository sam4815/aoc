import Foundation

let start = Date()

struct Coordinates {
  var x: Int
  var y: Int
  var z: Int
}

struct Nanobot {
  var position: Coordinates
  var radius: Int
}

let nanobots = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  let numbers = $0.components(
    separatedBy: CharacterSet.decimalDigits.union(CharacterSet(charactersIn: "-")).inverted
  ).filter({ !$0.isEmpty }).map({ Int($0)! })

  return Nanobot(
    position: Coordinates(x: numbers[0], y: numbers[1], z: numbers[2]), radius: numbers[3])
}

func findManhattanDistance(a: Coordinates, b: Coordinates) -> Int {
  abs(a.x - b.x) + abs(a.y - b.y) + abs(a.z - b.z)
}

func countNanobotsInRange(_ nanobot: Nanobot) -> Int {
  nanobots.filter({ findManhattanDistance(a: $0.position, b: nanobot.position) <= nanobot.radius })
    .count
}

let strongestNanobot = nanobots.max(by: { $0.radius < $1.radius })!
let numNanobotsInRange = countNanobotsInRange(strongestNanobot)

let orderedBotBounds = nanobots.flatMap { bot in
  let distanceToOrigin = findManhattanDistance(a: Coordinates(x: 0, y: 0, z: 0), b: bot.position)
  return [(distanceToOrigin - bot.radius, "open"), (distanceToOrigin + bot.radius + 1, "close")]
}.sorted(by: { $0.0 < $1.0 })

var numNanobots = 0
var maxNanobots = 0
var maxNanobotsPoint = 0

for (boundary, boundaryType) in orderedBotBounds {
  numNanobots += (boundaryType == "open" ? 1 : -1)

  if numNanobots > maxNanobots {
    maxNanobotsPoint = boundary
    maxNanobots = numNanobots
  }
}

print(
  """
  There are \(numNanobotsInRange) nanobots in range of the strongest nanobot.
  The manhattan distance to the point in range of the most nanobots is \(maxNanobotsPoint).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
