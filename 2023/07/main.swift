import Foundation

struct Hand {
  var score: Int = 0
  var jokerScore: Int = 0
  var cards: [Substring] = []
  var bid: Int = 0
}

let start = Date()
let basicRankings = ["A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"]
let jokerRankings = ["A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"]

func compareHands(handA: Hand, handB: Hand, isJokerWild: Bool = false) -> Bool {
  if handA.score != handB.score && !isJokerWild {
    return handA.score < handB.score
  }
  if handA.jokerScore != handB.jokerScore && isJokerWild {
    return handA.jokerScore < handB.jokerScore
  }

  let rankings = isJokerWild ? jokerRankings : basicRankings
  for i in 0...4 {
    let indexA = rankings.firstIndex(where: { $0 == handA.cards[i] })!
    let indexB = rankings.firstIndex(where: { $0 == handB.cards[i] })!

    if indexA != indexB {
      return indexA > indexB
    }
  }

  return false
}

let hands = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  let components = $0.split(separator: " ")
  let cards = components[0].split(separator: "")
  let cardQuantities = cards.reduce(into: [Substring: Int]()) { $0[$1, default: 0] += 1 }

  let scores = cardQuantities.values.sorted(by: >)
  let paddedScores = scores + Array(repeating: 0, count: 5 - scores.count)
  let score = Int(paddedScores.map { String($0) }.joined())!

  let nonJokerScores = cardQuantities.filter { $0.key != "J" }.values.sorted(by: >)
  var nonJokerPaddedScores = nonJokerScores + Array(repeating: 0, count: 5 - nonJokerScores.count)
  nonJokerPaddedScores[0] += cardQuantities["J"] ?? 0
  let jokerScore = Int(nonJokerPaddedScores.map { String($0) }.joined())!

  return Hand(score: score, jokerScore: jokerScore, cards: cards, bid: Int(components[1])!)
}

let totalWinnings = hands.sorted(by: { handA, handB in
  compareHands(handA: handA, handB: handB)
}).enumerated().reduce(0) { total, item in
  total + (item.offset + 1) * item.element.bid
}

let jokerTotalWinnings = hands.sorted(by: { handA, handB in
  compareHands(handA: handA, handB: handB, isJokerWild: true)
}).enumerated().reduce(0) { total, item in
  total + (item.offset + 1) * item.element.bid
}

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The total winnings are \(totalWinnings).
  With the Joker rule, the total winnings are \(jokerTotalWinnings).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
