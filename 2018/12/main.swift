import Foundation

let start = Date()

let parts = try String(contentsOfFile: "input.txt").split(separator: "\n\n")

let initialState = parts[0].split(separator: " ")[2].split(separator: "").map { String($0) }

let rules = parts[1].split(separator: "\n").reduce(into: [String: String]()) { rules, line in
  let components = line.split(separator: " => ").map { String($0) }
  rules[components[0]] = components[1]
}

func step(_ state: [String]) -> [String] {
  let padded = [".", ".", ".", "."] + state + [".", ".", ".", "."]
  var next: [String] = []

  for i in 2..<(padded.count - 2) {
    let str = padded[(i - 2)...(i + 2)].joined()
    next += [rules[str] ?? "."]
  }

  return next
}

func countPlants(_ plants: ([String], Int)) -> Int {
  plants.0.enumerated().reduce(0) { sum, plant in
    sum + (plant.element == "#" ? (plant.offset - plants.1) : 0)
  }
}

func applySteps(_ initialState: [String], numSteps: Int) -> ([String], Int) {
  var state = initialState
  var zeroIndex = 0
  for _ in 1...numSteps {
    state = step(state)
    zeroIndex += 2
  }
  return (state, zeroIndex)
}

let plantSum20 = countPlants(applySteps(initialState, numSteps: 20))

let plantSum99 = countPlants(applySteps(initialState, numSteps: 99))
let plantSum100 = countPlants(applySteps(initialState, numSteps: 100))
let plantSum50000000000 = (50_000_000_000 - 100) * (plantSum100 - plantSum99) + plantSum100

print(
  """
  After 20 generations, the sum of all plants is \(plantSum20).
  After 50000000000 generations, the sum of all plants is \(plantSum50000000000).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
