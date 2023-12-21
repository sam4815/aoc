import Foundation

let start = Date()

enum ModuleType { case flipflop, conjunction, broadcast }
enum PulseType { case high, low }
enum State { case on, off }

struct Pulse {
  var type: PulseType
  var source: String
  var destination: String
}

struct Module {
  var label: String
  var type: ModuleType
  var destinations: [String] = []
  var pulses: [String: PulseType] = [:]
  var state: State = .off
}

var modules = try String(contentsOfFile: "input.txt").split(separator: "\n").reduce(
  into: [String: Module]()
) { moduleMap, line in
  let components = line.split(separator: " -> ")
  let destinations = components[1].split(separator: ", ").map { String($0) }

  var label = String(components[0])
  let type = ["%": ModuleType.flipflop, "&": .conjunction, "b": .broadcast][label.first!]!
  if type != .broadcast { label.removeFirst() }

  moduleMap[label] = Module(label: label, type: type, destinations: destinations)
}

let target = modules.first(where: { $0.value.destinations.contains("rx") })!.key
var rxDependent: [String: Int] = [:]

modules.values.forEach { module in
  module.destinations.forEach { dest in
    if dest == target { rxDependent[module.label] = 0 }
    if modules[dest] != nil && modules[dest]!.type == .conjunction {
      modules[dest]!.pulses[module.label] = .low
    }
  }
}

var pulseCounts: [PulseType: Int] = [:]
var buttonPush = 0

let initialPulse = Pulse(type: .low, source: "button", destination: "broadcaster")
func pushButton() {
  var pulseQueue = [initialPulse]
  buttonPush += 1

  while pulseQueue.count > 0 {
    let pulse = pulseQueue.removeFirst()
    pulseCounts[pulse.type, default: 0] += 1

    if pulse.destination == target && pulse.type == .high {
      rxDependent[pulse.source] = buttonPush
    }

    if modules[pulse.destination] == nil { continue }

    var module = modules[pulse.destination]!

    switch module.type {
    case .broadcast:
      pulseQueue += module.destinations.map {
        Pulse(type: pulse.type, source: module.label, destination: $0)
      }
    case .flipflop:
      if pulse.type == .low {
        pulseQueue += module.destinations.map {
          Pulse(type: module.state == .on ? .low : .high, source: module.label, destination: $0)
        }
        module.state = module.state == .on ? .off : .on
      }
    case .conjunction:
      module.pulses[pulse.source] = pulse.type
      pulseQueue += module.destinations.map {
        Pulse(
          type: module.pulses.values.allSatisfy { $0 == .high } ? .low : .high,
          source: module.label,
          destination: $0)
      }
    }

    modules[pulse.destination] = module
  }
}

func gcd(a: Int, b: Int) -> Int {
  if b == 0 { return a }
  return gcd(a: b, b: a % b)
}

func lcm(a: Int, b: Int) -> Int {
  return (a * b) / gcd(a: a, b: b)
}

for _ in 1...1000 { pushButton() }
let pulsesProduct = pulseCounts[.high]! * pulseCounts[.low]!

while rxDependent.values.contains(0) { pushButton() }
let numButtonsForLowRx = rxDependent.values.reduce(1) { lcm(a: $0, b: $1) }

print(
  """
  The product of the low pulses and high pulses is \(pulsesProduct).
  The number of button presses required to deliver a low pulse to rx is \(numButtonsForLowRx).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
