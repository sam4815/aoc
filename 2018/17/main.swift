import Foundation

let start = Date()

enum Direction { case Left, Right }

var ground = try String(contentsOfFile: "input.txt").split(separator: "\n").reduce(
  into: [Int: [Int: String]]()
) { ground, line in
  let numbers = line.components(separatedBy: CharacterSet.decimalDigits.inverted).filter({
    !$0.isEmpty
  }).map({ Int($0)! })

  let xVein = line.first! == "x" ? numbers[0]...numbers[0] : numbers[1]...numbers[2]
  let yVein = line.first! == "x" ? numbers[1]...numbers[2] : numbers[0]...numbers[0]

  for x in xVein {
    for y in yVein {
      ground[y, default: [:]][x] = "#"
    }
  }
}

let xMin = ground.keys.min(by: <)!
let xMax = ground.keys.max(by: <)!

func isRestedWater(_ position: (Int, Int)) -> Bool {
  ground[position.0, default: [:]][position.1] == "~"
}

func isWater(_ position: (Int, Int)) -> Bool {
  ground[position.0, default: [:]][position.1] == "|" || isRestedWater(position)
}

func isClay(_ position: (Int, Int)) -> Bool {
  ground[position.0, default: [:]][position.1] == "#"
}

func isSupported(_ position: (Int, Int)) -> Bool {
  isRestedWater((position.0 + 1, position.1)) || isClay((position.0 + 1, position.1))
}

func isBlocked(_ position: (Int, Int), step: Int) -> Bool {
  isRestedWater((position.0, position.1 + step)) || isClay((position.0, position.1 + step))
}

func goDown(_ position: inout (Int, Int)) {
  while !isSupported(position) {
    position.0 += 1
    ground[position.0, default: [:]][position.1] = "|"

    if position.0 >= xMax {
      return
    }
  }
}

func goDirection(_ position: inout (Int, Int), direction: Direction) -> (Int, Int)? {
  let step = (direction == .Left ? -1 : 1)

  while !isBlocked(position, step: step) {
    position.1 += step
    ground[position.0, default: [:]][position.1] = "|"

    if !isSupported(position) {
      return position
    }
  }

  return nil
}

var waters = [(xMin - 1, 500)]

while waters.count > 0 {
  var downStream = waters.removeFirst()
  goDown(&downStream)
  if downStream.0 >= xMax {
    continue
  }

  var leftStream = downStream
  let leftFall = goDirection(&leftStream, direction: .Left)
  waters += leftFall == nil ? [] : [leftFall!]

  var rightStream = downStream
  let rightFall = goDirection(&rightStream, direction: .Right)
  waters += rightFall == nil ? [] : [rightFall!]

  if leftFall == nil && rightFall == nil {
    for y in leftStream.1...rightStream.1 {
      ground[downStream.0, default: [:]][y] = "~"

      if ground[downStream.0 - 1, default: [:]][y] == "|" {
        waters += [(downStream.0 - 1, y)]
      }
    }
  }
}

let waterQuantity = ground.values.reduce(0) { sum, row in
  sum + row.values.reduce(0) { $0 + ($1 == "|" || $1 == "~" ? 1 : 0) }
}

let restedWaterQuantity = ground.values.reduce(0) { sum, row in
  sum + row.values.reduce(0) { $0 + ($1 == "~" ? 1 : 0) }
}

print(
  """
  The number of tiles that can be reached by all of the water is \(waterQuantity).
  The number of tiles that can be reached by resting water is \(restedWaterQuantity).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
