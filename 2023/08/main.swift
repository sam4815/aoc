import Foundation

let start = Date()

let contents = try String(contentsOfFile: "input.txt").split(separator: "\n\n")
let directions = contents[0].split(separator: "")
let network = contents[1].split(separator: "\n").reduce(into: [String: [Substring: String]]()) {
  networks, line in
  let nodes = line.components(separatedBy: CharacterSet.alphanumerics.inverted).filter {
    !$0.isEmpty
  }
  networks[nodes[0]] = ["L": nodes[1], "R": nodes[2]]
}

func gcd(a: Int, b: Int) -> Int {
  if b == 0 { return a }
  return gcd(a: b, b: a % b)
}

func lcm(a: Int, b: Int) -> Int {
  return (a * b) / gcd(a: a, b: b)
}

func countSteps(start: String, endsWith: String) -> Int {
  var currPos = start
  var numSteps = 0

  while !currPos.hasSuffix(endsWith) {
    let direction = directions[numSteps % directions.count]
    currPos = network[currPos]![direction]!
    numSteps += 1
  }

  return numSteps
}

let numStepsToZZZ = countSteps(start: "AAA", endsWith: "ZZZ")

let numGhostStepsToZ = network.keys.filter({ $0.last! == "A" }).map {
  countSteps(start: $0, endsWith: "Z")
}.reduce(1) {
  lcm(a: $0, b: $1)
}

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The number of steps required to reach ZZZ is \(numStepsToZZZ).
  The number of ghost steps required to reach Z is \(numGhostStepsToZ).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
