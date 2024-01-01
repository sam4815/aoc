import Foundation

let start = Date()

public class Marble {
  var value: Int
  var prev: Marble?
  var next: Marble?

  init(value: Int, prev: Marble? = nil, next: Marble? = nil) {
    self.value = value
    self.prev = prev
    self.next = next
  }

  func insert(value: Int) -> Marble {
    let insertionMarble = Marble(value: value, prev: self.next!, next: self.next!.next!)
    self.next!.next!.prev = insertionMarble
    self.next!.next = insertionMarble

    return insertionMarble
  }

  func getNCounterClockwise(_ n: Int) -> Marble {
    var marble = self.prev!
    for _ in 0..<(n - 1) {
      marble = marble.prev!
    }
    return marble
  }

  func remove() -> Marble {
    self.prev!.next = self.next!
    self.next!.prev = self.prev!

    return self.next!
  }
}

func playGame(numPlayers: Int, numMarbles: Int) -> Int {
  var scores = Array(repeating: 0, count: numPlayers)
  var nextMarbleValue = 1
  var currentPlayerIndex = 0

  var currentMarble = Marble(value: 0)
  currentMarble.prev = currentMarble
  currentMarble.next = currentMarble

  while nextMarbleValue <= numMarbles {
    if nextMarbleValue % 23 == 0 {
      scores[currentPlayerIndex] += nextMarbleValue

      currentMarble = currentMarble.getNCounterClockwise(7)
      scores[currentPlayerIndex] += currentMarble.value

      currentMarble = currentMarble.remove()
    } else {
      currentMarble = currentMarble.insert(value: nextMarbleValue)
    }

    nextMarbleValue += 1
    currentPlayerIndex = (currentPlayerIndex + 1) % numPlayers
  }

  return scores.max(by: <)!
}

let gameParams = try String(contentsOfFile: "input.txt").components(
  separatedBy: CharacterSet.decimalDigits.inverted
).filter { !$0.isEmpty }.map { Int($0)! }

let winningScore = playGame(numPlayers: gameParams[0], numMarbles: gameParams[1])
let winningScore100 = playGame(numPlayers: gameParams[0], numMarbles: gameParams[1] * 100)

print(
  """
  The winning Elf's score is \(winningScore).
  If the last marble is 100 times larger, the winning Elf's score is \(winningScore100).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
