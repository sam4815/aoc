import Foundation

let start = Date()

let polymer = Array(
  try String(contentsOfFile: "input.txt").trimmingCharacters(
    in: .whitespacesAndNewlines
  )
).map { $0.asciiValue! }

func reactPolymer(polymer: [UInt8]) -> [UInt8] {
  var reactedPolymer = polymer
  var index = 0

  while index < reactedPolymer.count - 1 {
    let currCharCode = reactedPolymer[index]
    let nextCharCode = reactedPolymer[index + 1]

    if currCharCode == nextCharCode + 32 || nextCharCode == currCharCode + 32 {
      reactedPolymer.remove(at: index)
      reactedPolymer.remove(at: index)
      if index > 0 { index -= 1 }
    } else {
      index += 1
    }
  }

  return reactedPolymer
}

let polymerUnits = reactPolymer(polymer: polymer).count

let improvedPolymers = (65...90).map({ charCode in
  reactPolymer(polymer: polymer.filter { $0 != charCode && $0 != charCode + 32 })
})

let shortestPolymerUnits = improvedPolymers.min(by: { $0.count < $1.count })!.count

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  After all possible reactions, the resulting polymer contains \(polymerUnits) units.
  The shortest possible polymer contains \(shortestPolymerUnits) units.
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
