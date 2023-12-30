import Foundation

let start = Date()

struct Step {
  var name: String
  var requires: [String]
}

let steps = try String(contentsOfFile: "input.txt").split(separator: "\n").reduce(
  into: [String: Step]()
) { steps, line in
  let name = String(line.split(separator: "")[36])
  let requires = String(line.split(separator: "")[5])

  steps[name] = Step(name: name, requires: (steps[name]?.requires ?? []) + [requires])
  steps[requires] = Step(name: requires, requires: (steps[requires]?.requires ?? []) + [])
}

func completeSteps(_ steps: [String: Step], delay: Bool = true, totalWorkers: Int = 1) -> (
  String, Int
) {
  var stepOrder = ""
  var secondsPassed = 0

  var workRemaining: [String: Int] = [:]
  var availableWorkers = totalWorkers

  while stepOrder.count < steps.count {
    let candidates = steps.values.filter {
      $0.requires.allSatisfy({ workRemaining[$0, default: Int.max] <= 0 })
        && workRemaining[$0.name] == nil
    }

    for candidate in candidates.sorted(by: { $0.name < $1.name }) {
      if availableWorkers == 0 { break }

      workRemaining[candidate.name] = delay ? (Int(Character(candidate.name).asciiValue!) - 4) : 0
      availableWorkers -= 1
    }

    for step in workRemaining.keys {
      workRemaining[step]! -= 1

      if workRemaining[step]! == 0 {
        stepOrder += step
        availableWorkers += 1
      }
    }

    secondsPassed += 1
  }

  return (stepOrder, secondsPassed)
}

let (stepOrder, _) = completeSteps(steps)
let (_, secondsPassed) = completeSteps(steps, delay: true, totalWorkers: 5)

print(
  """
  The order the steps can be completed in is \(stepOrder).
  With five workers, it will take \(secondsPassed) seconds to complete all of the steps.
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
