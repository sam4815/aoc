import Foundation

let start = Date()

struct Light {
  var position: (Int, Int)
  var velocity: (Int, Int)
}

let lights = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  let numbers = $0.components(
    separatedBy: CharacterSet.decimalDigits.union(CharacterSet(charactersIn: "-")).inverted
  ).filter { !$0.isEmpty }.map { Int($0)! }

  return Light(position: (numbers[1], numbers[0]), velocity: (numbers[3], numbers[2]))
}

func calculatePositions(lights: [Light], time: Int) -> [Int: [Int: Bool]] {
  lights.reduce(into: [Int: [Int: Bool]]()) { area, light in
    let position = (
      light.position.0 + light.velocity.0 * time, light.position.1 + light.velocity.1 * time
    )
    area[position.0, default: [:]][position.1] = true
  }
}

func findXYRanges(_ positions: [Int: [Int: Bool]]) -> ((Int, Int), (Int, Int)) {
  let xMin = positions.keys.min(by: <)!
  let xMax = positions.keys.max(by: <)!
  let yMin = positions.flatMap { $0.value.keys }.min(by: <)!
  let yMax = positions.flatMap { $0.value.keys }.max(by: <)!

  return ((xMin, xMax), (yMin, yMax))
}

let timeWhenMessageAppears = (0...12000).map {
  let positions = calculatePositions(lights: lights, time: $0)
  let ((xMin, xMax), (yMin, yMax)) = findXYRanges(positions)
  return abs(xMin - xMax) + abs(yMin - yMax)
}.enumerated().min { $0.element < $1.element }!

let positions = calculatePositions(lights: lights, time: timeWhenMessageAppears.offset)
let ((xMin, xMax), (yMin, yMax)) = findXYRanges(positions)

var message = ""
for x in xMin...xMax {
  var row = ""
  for y in yMin...yMax {
    row += (positions[x]?[y] ?? false) ? "#" : "."
  }
  message += "\n" + row
}

print(
  """
  The message that appears is \(message)
  The message appears after \(timeWhenMessageAppears.offset) seconds.
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
