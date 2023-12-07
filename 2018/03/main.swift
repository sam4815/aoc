import Foundation

let start = Date()

let claims = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  $0.components(separatedBy: CharacterSet.decimalDigits.inverted).filter({ !$0.isEmpty }).map({
    Int($0)!
  })
}

let positions = claims.reduce(into: [String: Int]()) { posMap, claim in
  for x in (claim[1]..<(claim[1] + claim[3])) {
    for y in (claim[2]..<(claim[2] + claim[4])) {
      let inch = "\(x),\(y)"
      posMap[inch] = (posMap[inch] ?? 0) + 1
    }
  }
}

let inchesOverlapping = positions.filter { $0.1 > 1 }.count

let nonOverlappingClaim = claims.first(where: { claim in
  for x in (claim[1]..<(claim[1] + claim[3])) {
    for y in (claim[2]..<(claim[2] + claim[4])) {
      let inch = "\(x),\(y)"
      if positions[inch]! > 1 {
        return false
      }
    }
  }
  return true
})!

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The number of inches of fabric with overlapping claims is \(inchesOverlapping).
  The ID of the claim that doesn't overlap is \(nonOverlappingClaim[0]).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
