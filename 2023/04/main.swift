import Foundation

let start = Date()

let content = try String(contentsOfFile: "input.txt")
let lines = content.split(separator: "\n")

let victories = lines.map { line in
  let card = line.split(separator: ": ")[1].split(separator: " | ").map {
    $0.split(separator: " ").map({ Int($0)! })
  }
  return card[1].filter({ card[0].contains($0) }).count
}

let pointsSum = victories.map { $0 > 0 ? pow(2, $0 - 1) : 0 }.reduce(0, +)

var cardQuantities = lines.enumerated().reduce(into: [Int: Int]()) { acc, curr in
  acc[curr.0 + 1] = 1
}

for (index, victory) in victories.enumerated() {
  let cardNumber = index + 1
  for offset in ((cardNumber + 1)..<(cardNumber + 1 + victory)) {
    cardQuantities[offset]! += cardQuantities[cardNumber]!
  }
}

let sumCardQuantities = cardQuantities.values.reduce(0, +)

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The scratchcards are worth \(pointsSum) points total.
  The total number of scratchcards is \(sumCardQuantities).
  Solution generated in \(timeElapsed).
  """)
