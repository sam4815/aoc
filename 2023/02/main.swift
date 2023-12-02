import Foundation

let start = Date()

let contents = try? String(contentsOfFile: "input.txt")

let games = contents!.split(separator: "\n").reduce(into: [Int: [[String: Int]]]()) { games, line in
  let gameId = Int(line.firstMatch(of: /Game (\d+)/)!.1)!
  let cubes = line.split(separator: ": ")[1].split(separator: "; ")

  games[gameId] = cubes.map {
    $0.matches(of: /(\d+) (\w+)/).reduce(
      into: [String: Int]()
    ) { $0[String($1.2)] = Int($1.1) }
  }
}

let validGame = ["red": 12, "green": 13, "blue": 14]

let sumValidGameIds = games.reduce(0) { sum, game in
  game.value.allSatisfy({
    $0.allSatisfy({ (colour, quantity) in quantity <= validGame[colour]! })
  }) ? game.key + sum : sum
}

let sumMinimumPowers = games.reduce(0) { sum, game in
  let minimumCubes = game.value.reduce(into: ["red": 0, "green": 0, "blue": 0]) { minimums, cubes in
    minimums.forEach { minimums[$0.key] = max(minimums[$0.key]!, cubes[$0.key] ?? 0) }
  }
  return sum + minimumCubes.reduce(1) { $0 * $1.value }
}

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The sum of the valid game IDs is \(sumValidGameIds).
  The sum of the minimum powers for each game is \(sumMinimumPowers).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
