import Foundation

let start = Date()

enum Faction { case Goblin, Elf }

struct Unit {
  var label: String
  var HP: Int
  var attackPower: Int
  var position: (Int, Int)
  var faction: Faction
}

public class Fight {
  var cave: [[String]]
  var units: [String: Unit] = [:]
  var round: Int = 0
  var hasElfDied: Bool = false
  var hasBattleEnded: Bool = false

  init(cave: [[String]], elfAttackPower: Int) {
    self.cave = cave

    for (x, row) in cave.enumerated() {
      for (y, cell) in row.enumerated() {
        if !["G", "E"].contains(cell) { continue }

        let label = String(Character(UnicodeScalar(self.units.count + 65)!))

        self.units[label] = Unit(
          label: label, HP: 200,
          attackPower: cell == "G" ? 3 : elfAttackPower,
          position: (x, y),
          faction: (cell == "G" ? .Goblin : .Elf))

        self.cave[x][y] = label
      }
    }
  }

  func getManhattanDistance(start: (Int, Int), end: (Int, Int)) -> Int {
    abs(start.0 - end.0) + abs(start.1 - end.1)
  }

  func hasLowerManhattanDistance(_ a: (Int, Int), _ b: (Int, Int), start: (Int, Int)) -> Bool {
    getManhattanDistance(start: start, end: a) < getManhattanDistance(start: start, end: b)
  }

  func getReadingNumber(_ position: (Int, Int)) -> Int {
    position.0 * self.cave[0].count + position.1
  }

  func hasLowerReadingNumber(_ a: (Int, Int), _ b: (Int, Int)) -> Bool {
    getReadingNumber(a) < getReadingNumber(b)
  }

  func killUnit(_ unit: Unit) {
    self.cave[unit.position.0][unit.position.1] = "."
    self.units.removeValue(forKey: unit.label)

    if unit.faction == .Elf {
      self.hasElfDied = true
    }
  }

  func findAdjacentTiles(_ position: (Int, Int)) -> [(Int, Int)] {
    [
      (position.0 - 1, position.1), (position.0, position.1 - 1),
      (position.0, position.1 + 1), (position.0 + 1, position.1),
    ]
  }

  func findAdjacentEnemies(_ unit: Unit) -> [Unit] {
    var enemies: [Unit] = []

    for (x, y) in findAdjacentTiles(unit.position) {
      if let enemy = self.units[self.cave[x][y]] {
        enemies += unit.faction == enemy.faction ? [] : [enemy]
      }
    }

    return enemies.sorted(by: {
      $0.HP == $1.HP ? getReadingNumber($0.position) < getReadingNumber($1.position) : $0.HP < $1.HP
    })
  }

  func findShortestPath(start: (Int, Int), end: (Int, Int)) -> [(Int, Int)]? {
    var queue: [[(Int, Int)]] = [[start]]
    var visited: [Int: [Int: Bool]] = [:]

    while queue.count > 0 {
      let currentPath = queue.removeFirst()
      let lastPosition = currentPath.last!

      if lastPosition == end { return currentPath }

      if visited[lastPosition.0]?[lastPosition.1] ?? false {
        continue
      } else {
        visited[lastPosition.0, default: [:]][lastPosition.1] = true
      }

      for move in findAdjacentTiles(lastPosition) {
        if self.cave[move.0][move.1] != "." { continue }
        queue += [currentPath + [move]]
      }
    }

    return nil
  }

  func findTargets(_ unit: Unit) -> [(Int, Int)] {
    var targets: [(Int, Int)] = []

    let enemies = self.units.values.filter({ $0.faction != unit.faction })
    if enemies.count == 0 {
      self.hasBattleEnded = true
    }
    for enemy in enemies {
      targets += findAdjacentTiles(enemy.position).filter { self.cave[$0.0][$0.1] == "." }
    }

    return targets.sorted(by: { hasLowerManhattanDistance($0, $1, start: unit.position) })
  }

  func performMove(_ unit: Unit) {
    let adjacentEnemies = findAdjacentEnemies(unit)
    if adjacentEnemies.count > 0 { return }

    let targets = findTargets(unit)
    var path: [(Int, Int)] = []
    var chosenTarget: (Int, Int)?

    for target in targets {
      if path.count < getManhattanDistance(start: unit.position, end: target) {
        if chosenTarget != nil { break }
      }

      if let targetPath = findShortestPath(start: unit.position, end: target) {
        if targetPath.count < path.count || chosenTarget == nil
          || (path.count == targetPath.count
            && hasLowerReadingNumber(target, chosenTarget ?? (200, 200)))
        {
          path = targetPath
          chosenTarget = target
        }
      }
    }

    if chosenTarget != nil {
      self.cave[unit.position.0][unit.position.1] = "."
      self.units[unit.label]!.position = path[1]
      self.cave[path[1].0][path[1].1] = unit.label
    }
  }

  func performAttack(_ unit: Unit) {
    let adjacentEnemies = findAdjacentEnemies(unit)
    if adjacentEnemies.count == 0 { return }

    let enemyToAttack = adjacentEnemies[0]
    self.units[enemyToAttack.label]!.HP -= unit.attackPower

    if self.units[enemyToAttack.label]!.HP <= 0 {
      killUnit(enemyToAttack)
    }
  }

  func performRound() {
    for unit in self.units.values.sorted(by: { hasLowerReadingNumber($0.position, $1.position) }) {
      if self.units[unit.label] == nil { continue }

      performMove(self.units[unit.label]!)
      if self.hasBattleEnded { return }
      performAttack(self.units[unit.label]!)
    }

    self.round += 1
  }

  func simulateCombat(allowElfDeath: Bool = true) {
    while !self.hasBattleEnded {
      self.performRound()

      if !allowElfDeath && self.hasElfDied { break }
    }
  }

  func combatOutcome() -> Int {
    self.round * self.units.values.reduce(0) { sum, unit in sum + unit.HP }
  }
}

let cave = try String(contentsOfFile: "input.txt").split(separator: "\n").map {
  $0.split(separator: "").map { String($0) }
}

let fight = Fight(cave: cave, elfAttackPower: 3)
fight.simulateCombat()
let combatOutcome = fight.combatOutcome()

var zeroDeathOutcome = 0
var attackPower = 4

while zeroDeathOutcome == 0 {
  let fight = Fight(cave: cave, elfAttackPower: attackPower)
  fight.simulateCombat(allowElfDeath: false)

  if fight.hasElfDied {
    attackPower += 1
  } else {
    zeroDeathOutcome = fight.combatOutcome()
  }
}

print(
  """
  The outcome of the initial combat is \(combatOutcome).
  The outcome of the combat in which no Elves die is \(zeroDeathOutcome).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
