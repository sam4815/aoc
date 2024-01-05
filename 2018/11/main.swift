import Foundation

let start = Date()

let serialNumber = Int(try String(contentsOfFile: "input.txt").split(separator: "\n")[0])!

func findPowerLevel(_ x: Int, _ y: Int) -> Int {
  let powerLevel = (((x + 10) * y) + serialNumber) * (x + 10)
  let hundredsDigit = (powerLevel % 1000) / 100
  return hundredsDigit - 5
}

func sumSquare(_ point: (Int, Int), size: Int, levels: [Int: [Int: Int]]) -> Int {
  var sum = 0
  for x in 0..<size {
    for y in 0..<size {
      sum += levels[point.0 + x]![point.1 + y]!
    }
  }
  return sum
}

func findMaxForSize(size: Int, levels: [Int: [Int: Int]]) -> (
  point: (Int, Int), (power: Int, size: Int)
) {
  var maxPower = 0
  var maxCoordinates = (0, 0)

  for x in 1...(301 - size) {
    for y in 1...(301 - size) {
      let power = sumSquare((x, y), size: size, levels: powerLevels)

      if power > maxPower {
        maxPower = power
        maxCoordinates = (x, y)
      }
    }
  }

  return (point: maxCoordinates, (power: maxPower, size: size))
}

let powerLevels = (1...300).reduce(into: [Int: [Int: Int]]()) { levels, x in
  for y in 1...300 { levels[x, default: [:]][y] = findPowerLevel(x, y) }
}

let maxCoordinatesNxN = (1...15).map { findMaxForSize(size: $0, levels: powerLevels) }
let maxCoordinates3x3 = maxCoordinatesNxN[2]
let maxMaxCoordinatesNxN = maxCoordinatesNxN.max(by: { $0.1.power < $1.1.power })!

print(
  """
  The X,Y coordinates of the 3x3 square with the most power is \(maxCoordinates3x3.point).
  The X,Y coordinates of any square with the most power is \(maxMaxCoordinatesNxN.point) with size \(maxMaxCoordinatesNxN.1.size).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
