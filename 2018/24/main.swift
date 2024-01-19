import Foundation

let start = Date()

enum ArmyType { case ImmuneSystem, Infection }
enum AttackType: String { case Slashing, Bludgeoning, Cold, Radiation, Fire }

struct Army {
  var type: ArmyType
  var groups: [String: Group]
}

struct Group {
  var id: String
  var armyType: ArmyType
  var maxHealth: Int
  var numUnits: Int
  var weakTo: [AttackType]
  var immuneTo: [AttackType]
  var attack: Attack
  var initiative: Int
}

struct Attack {
  var type: AttackType
  var damage: Int
}

func parseGroup(_ str: String, armyType: ArmyType) -> Group {
  let numbers = str.components(
    separatedBy: CharacterSet.decimalDigits.union(CharacterSet(charactersIn: "-")).inverted
  ).filter({ !$0.isEmpty }).map({ Int($0)! })

  let weakTo: [AttackType] = str.matches(of: try! Regex("weak to [^;)]+")).flatMap {
    $0.0.split(separator: try! Regex("[^\\w]"))[2...].map {
      AttackType(rawValue: String($0).capitalized)!
    }
  }

  let immuneTo: [AttackType] = str.matches(of: try! Regex("immune to [^;)]+")).flatMap {
    $0.0.split(separator: try! Regex("[^\\w]"))[2...].map {
      AttackType(rawValue: String($0).capitalized)!
    }
  }

  let attackType: AttackType = AttackType(
    rawValue: String(
      str.matches(of: try! Regex("\\d+ \\w+ damage")).flatMap { $0.0.split(separator: " ") }[1]
    ).capitalized)!

  let id = String(str.matches(of: try! Regex("\\d+ \\w+ damage"))[0].0)

  return Group(
    id: id,
    armyType: armyType,
    maxHealth: numbers[1],
    numUnits: numbers[0],
    weakTo: weakTo, immuneTo: immuneTo,
    attack: Attack(type: attackType, damage: numbers[2]), initiative: numbers[3])
}

var armies = try String(contentsOfFile: "input.txt").split(separator: "\n\n").reduce(
  into: [ArmyType: Army]()
) {
  let armyType: ArmyType = $1.contains("Immune") ? .ImmuneSystem : .Infection
  let groups = $1.split(separator: "\n")[1...].enumerated().reduce(into: [String: Group]()) {
    groups, line in
    let group = parseGroup(String(line.element), armyType: armyType)
    groups[group.id] = group
  }

  $0[armyType] = Army(type: armyType, groups: groups)
}

public class Battle {
  var armies: [ArmyType: Army]
  var numRounds: Int = 0

  init(_ armies: [ArmyType: Army]) {
    self.armies = armies
  }

  func boostImmuneSystem(_ n: Int) {
    for id in self.armies[.ImmuneSystem]!.groups.keys {
      self.armies[.ImmuneSystem]!.groups[id]!.attack.damage += n
    }
  }

  func getEffectivePower(_ group: Group) -> Int {
    group.numUnits * group.attack.damage
  }

  func sortTargetSelection(a: Group, b: Group) -> Bool {
    let aEffectivePower = getEffectivePower(a)
    let bEffectivePower = getEffectivePower(b)

    if aEffectivePower == bEffectivePower { return a.initiative > b.initiative }
    return aEffectivePower > bEffectivePower
  }

  func sortTargets(attacker: Group, targetA: Group, targetB: Group) -> Bool {
    let aDamageDealt = calculateDamageDealt(attacker: attacker, defender: targetA)
    let bDamageDealt = calculateDamageDealt(attacker: attacker, defender: targetB)

    if aDamageDealt == bDamageDealt { return sortTargetSelection(a: targetA, b: targetB) }
    return aDamageDealt > bDamageDealt
  }

  func calculateDamageDealt(attacker: Group, defender: Group) -> Int {
    if defender.immuneTo.contains(attacker.attack.type) { return 0 }
    if defender.weakTo.contains(attacker.attack.type) { return getEffectivePower(attacker) * 2 }
    return getEffectivePower(attacker)
  }

  func listGroups(_ armies: [ArmyType: Army]) -> [Group] {
    armies.values.flatMap({ $0.groups.values })
  }

  func fight() {
    var targets: [String: Group] = [:]

    for group in listGroups(armies).sorted(by: { sortTargetSelection(a: $0, b: $1) }) {
      let opposingGroups =
        group.armyType == .ImmuneSystem
        ? armies[.Infection]!.groups.values : armies[.ImmuneSystem]!.groups.values

      let sortedOpposingGroups = opposingGroups.filter { g in
        targets.values.allSatisfy { $0.id != g.id }
      }.sorted(by: {
        sortTargets(attacker: group, targetA: $0, targetB: $1)
      })

      if let target = sortedOpposingGroups.first(where: {
        calculateDamageDealt(attacker: group, defender: $0) > 0
      }) {
        targets[group.id] = target
      }
    }

    for g in listGroups(armies).sorted(by: { $0.initiative > $1.initiative }) {
      if targets[g.id] == nil || armies[g.armyType]!.groups[g.id] == nil { continue }

      let attacker = armies[g.armyType]!.groups[g.id]!
      let defender = targets[attacker.id]!
      let damageDealt = calculateDamageDealt(attacker: attacker, defender: defender)
      let unitsLost = damageDealt / defender.maxHealth

      armies[defender.armyType]!.groups[defender.id]!.numUnits -= unitsLost

      if armies[defender.armyType]!.groups[defender.id]!.numUnits <= 0 {
        armies[defender.armyType]!.groups.removeValue(forKey: defender.id)
      }
    }
  }

  func getWinner() -> ArmyType? {
    if armies[.ImmuneSystem]!.groups.count == 0 { return .Infection }
    if armies[.Infection]!.groups.count == 0 { return .ImmuneSystem }
    return nil
  }

  func countUnits() -> Int {
    armies[.Infection]!.groups.values.map { $0.numUnits }.reduce(0, +)
      + armies[.ImmuneSystem]!.groups.values.map { $0.numUnits }.reduce(0, +)
  }

  func simulate() {
    while getWinner() == nil && numRounds < 5000 {
      fight()
      numRounds += 1
    }
  }
}

let battle = Battle(armies)
battle.simulate()
let winningArmyUnitsCount = battle.countUnits()

var boost = 1
var boostedArmyUnitsCount = 0

while boostedArmyUnitsCount == 0 {
  let boostedBattle = Battle(armies)
  boostedBattle.boostImmuneSystem(boost)
  boostedBattle.simulate()

  if boostedBattle.getWinner() == .ImmuneSystem {
    boostedArmyUnitsCount = boostedBattle.countUnits()
  }

  boost += 1
}

print(
  """
  At the end of the fight, the winning army will have \(winningArmyUnitsCount) units.
  At the end of the boosted fight, the winning army will have \(boostedArmyUnitsCount) units.
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
