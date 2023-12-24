import Foundation

let start = Date()

struct Dimensions {
  var x: Double
  var y: Double
  var z: Double
}

struct Hailstone {
  var position: Dimensions
  var velocity: Dimensions
}

struct LinearEquation {
  var xCoefficient: Decimal
  var constant: Decimal
}

let hailstones = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  let parts = $0.split(separator: " @ ").map {
    $0.split(separator: ",").map { Double($0.trimmingCharacters(in: .whitespacesAndNewlines))! }
  }

  let position = Dimensions(x: parts[0][0], y: parts[0][1], z: parts[0][2])
  let velocity = Dimensions(x: parts[1][0], y: parts[1][1], z: parts[1][2])

  return Hailstone(
    position: position,
    velocity: velocity
  )
}

func findEquation(hailstone: Hailstone) -> LinearEquation {
  let xCoefficient = Decimal(hailstone.velocity.y) / Decimal(hailstone.velocity.x)
  let constant = Decimal(hailstone.position.y) - (Decimal(hailstone.position.x) * xCoefficient)

  return LinearEquation(xCoefficient: xCoefficient, constant: constant)
}

func findIntersection(hailstoneA: Hailstone, hailstoneB: Hailstone) -> (Double, Double) {
  let a = findEquation(hailstone: hailstoneA)
  let b = findEquation(hailstone: hailstoneB)

  let LHS = a.xCoefficient - b.xCoefficient
  let RHS = b.constant - a.constant
  let x = RHS / LHS
  let y = a.xCoefficient * x + a.constant

  return (NSDecimalNumber(decimal: x).doubleValue, NSDecimalNumber(decimal: y).doubleValue)
}

func isPointOutsideTestArea(_ point: (Double, Double)) -> Bool {
  let testAreaMin: Double = 200000000000000.0
  let testAreaMax: Double = 400000000000000.0

  return point.0 < testAreaMin || point.0 > testAreaMax
    || point.1 < testAreaMin || point.1 > testAreaMax
}

func isPointInFuture(hailstone: Hailstone, intersection: (Double, Double)) -> Bool {
  let xInFuture =
    (hailstone.velocity.x > 0.0 && intersection.0 > hailstone.position.x
      || hailstone.velocity.x < 0.0 && intersection.0 < hailstone.position.x)
  let yInFuture =
    (hailstone.velocity.y > 0.0 && intersection.1 > hailstone.position.y
      || hailstone.velocity.y < 0.0 && intersection.1 < hailstone.position.y)

  return xInFuture && yInFuture
}

func find2DIntersections(_ hailstones: [Hailstone]) -> Int {
  var numTestIntersections = 0

  for (indexA, hailstoneA) in hailstones.enumerated() {
    let hailstonesForComparison = hailstones[(indexA + 1)...]

    for hailstoneB in hailstonesForComparison {
      let intersection = findIntersection(hailstoneA: hailstoneA, hailstoneB: hailstoneB)

      if isPointOutsideTestArea(intersection) { continue }
      if !isPointInFuture(hailstone: hailstoneA, intersection: intersection) { continue }
      if !isPointInFuture(hailstone: hailstoneB, intersection: intersection) { continue }

      numTestIntersections += 1
    }
  }

  return numTestIntersections
}

func addVelocity(_ hailstone: Hailstone, dx: Double, dy: Double) -> Hailstone {
  return Hailstone(
    position: Dimensions(
      x: hailstone.position.x, y: hailstone.position.y, z: hailstone.position.z),
    velocity: Dimensions(
      x: hailstone.velocity.x + dx, y: hailstone.velocity.y + dy, z: hailstone.velocity.z))
}

func find2DCollisionStone(_ hailstones: [Hailstone]) -> ((Double, Double), (Double, Double)) {
  for dx in -250...250 {
    for dy in -250...250 {
      let hailstoneA = addVelocity(hailstones.first!, dx: Double(dx), dy: Double(dy))
      let hailstoneRange = 30...34
      var intersections: (Set<Double>, Set<Double>) = (Set(), Set())
      var numCollisions = 0

      for hailB in hailstones[hailstoneRange] {
        let hailstoneB = addVelocity(hailB, dx: Double(dx), dy: Double(dy))
        let intersection = findIntersection(hailstoneA: hailstoneA, hailstoneB: hailstoneB)
        let roundedIntersection = (round(intersection.0), round(intersection.1))

        if !isPointInFuture(hailstone: hailstoneA, intersection: roundedIntersection) { continue }
        if !isPointInFuture(hailstone: hailstoneB, intersection: roundedIntersection) { continue }

        numCollisions += 1

        intersections.0.insert(roundedIntersection.0)
        intersections.1.insert(roundedIntersection.1)
      }

      if intersections.0.count == 1 && numCollisions == hailstoneRange.count {
        return ((intersections.0.first!, intersections.1.first!), (Double(-dx), Double(-dy)))
      }
    }
  }

  return ((0, 0), (0, 0))
}

func findCollisionStone(_ hailstones: [Hailstone]) -> Hailstone {
  let yzSwappedHailstones = hailstones.map {
    Hailstone(
      position: Dimensions(x: $0.position.x, y: $0.position.z, z: $0.velocity.y),
      velocity: Dimensions(x: $0.velocity.x, y: $0.velocity.z, z: $0.velocity.y))
  }

  let xy = find2DCollisionStone(hailstones)
  let xz = find2DCollisionStone(yzSwappedHailstones)

  return Hailstone(
    position: Dimensions(x: xy.0.0, y: xy.0.1, z: xz.0.1),
    velocity: Dimensions(x: xy.1.0, y: xy.1.1, z: xz.1.1))
}

let numTestIntersections = find2DIntersections(hailstones)
let collisionStone = findCollisionStone(hailstones)
let collisionStoneSum =
  Int(collisionStone.position.x + collisionStone.position.y + collisionStone.position.z)

print(
  """
  The number of interestions that occur within the test area is \(numTestIntersections).
  Adding up the x, y, and z position of the collision stone gives \(collisionStoneSum).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
