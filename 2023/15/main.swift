import Foundation

let start = Date()

struct Step {
  var label: Substring
  var hash: Int
  var operation: String
  var focalLength: Int = 0
}

func hash(_ str: Substring) -> Int {
  Array(str).reduce(into: 0) { hash, char in
    hash += Int(exactly: char.asciiValue!)!
    hash *= 17
    hash %= 256
  }
}

let steps = try String(contentsOfFile: "input.txt").split(separator: ",")
let hashSum = steps.reduce(0) { $0 + hash($1) }

let decodedSteps = try String(contentsOfFile: "input.txt").split(separator: ",").map { step in
  if String(step).hasSuffix("-") {
    let label = step.split(separator: "-")[0]
    return Step(label: label, hash: hash(label), operation: "-")
  }

  let components = step.split(separator: "=")
  return Step(
    label: components[0], hash: hash(components[0]), operation: "=",
    focalLength: Int(components[1])!)
}

let hashMap = decodedSteps.reduce(into: [Int: [(Substring, Int)]]()) { hashMap, step in
  var box = hashMap[step.hash, default: []]

  if step.operation == "=" {
    if let index = box.firstIndex(where: { $0.0 == step.label }) {
      box[index] = (step.label, step.focalLength)
    } else {
      box += [(step.label, step.focalLength)]
    }
  } else if let index = box.firstIndex(where: { $0.0 == step.label }) {
    box.remove(at: index)
  }

  hashMap[step.hash] = box
}

let focusSum = hashMap.reduce(into: 0) { sum, hash in
  for (index, lens) in hash.value.enumerated() {
    sum += (hash.key + 1) * (index + 1) * lens.1
  }
}

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The sum of the HASH results is \(hashSum).
  The sum of the focusing powers is \(focusSum).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
