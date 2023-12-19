import Foundation

let start = Date()

enum Operation {
  case GT, LT
}

enum Part {
  case x, m, a, s
}

struct Parts {
  var x: Int
  var m: Int
  var a: Int
  var s: Int
}

struct PartsRange {
  var x: ClosedRange<Int>
  var m: ClosedRange<Int>
  var a: ClosedRange<Int>
  var s: ClosedRange<Int>
}

struct Workflow {
  var label: Substring
  var outcome: String
  var rules: [Rule]
}

struct Rule {
  var part: Part
  var operation: Operation
  var number: Int
  var next: String
}

let sections = try String(contentsOfFile: "input.txt").split(separator: "\n\n")

let workflows = sections[0].split(separator: "\n").reduce(into: [String: Workflow]()) {
  workflows, line in
  let components = String(line).trimmingCharacters(in: CharacterSet(charactersIn: "}")).split(
    separator: "{")
  let rules = components[1].split(separator: ",")

  workflows[String(components[0])] = Workflow(
    label: components[0], outcome: String(rules.last!),
    rules: rules.dropLast().map { rule in
      let number = Int(
        String(rule).components(separatedBy: CharacterSet.decimalDigits.inverted).joined())!
      let part: Part = ["x": .x, "m": .m, "a": .a, "s": .s][rule.first!]!
      let operation: Operation = [">": .GT, "<": .LT][rule.split(separator: "")[1]]!
      let next = String(rule.split(separator: ":")[1])

      return Rule(part: part, operation: operation, number: number, next: next)
    })
}

let parts = sections[1].split(separator: "\n").map { line in
  String(line).trimmingCharacters(in: CharacterSet(charactersIn: "{}")).split(separator: ",")
    .reduce(
      into: Parts(x: 0, m: 0, a: 0, s: 0)
    ) {
      let components = $1.split(separator: "=")
      switch components[0] {
      case "x": $0.x = Int(components[1])!
      case "m": $0.m = Int(components[1])!
      case "a": $0.a = Int(components[1])!
      case "s": $0.s = Int(components[1])!
      default: ()
      }
    }
}

func arePartsAccepted(_ parts: Parts) -> Bool {
  var workflowLabel = "in"
  workflowProcess: while workflowLabel != "A" && workflowLabel != "R" {
    let workflow = workflows[workflowLabel]!
    for rule in workflow.rules {
      let passed: Bool
      switch rule.part {
      case .x: passed = rule.operation == .GT ? (parts.x > rule.number) : (parts.x < rule.number)
      case .m: passed = rule.operation == .GT ? (parts.m > rule.number) : (parts.m < rule.number)
      case .a: passed = rule.operation == .GT ? (parts.a > rule.number) : (parts.a < rule.number)
      case .s: passed = rule.operation == .GT ? (parts.s > rule.number) : (parts.s < rule.number)
      }

      if passed {
        workflowLabel = rule.next
        continue workflowProcess
      }
    }

    workflowLabel = workflow.outcome
  }

  return workflowLabel == "A"
}

func splitRange(range: ClosedRange<Int>, rule: Rule) -> ([ClosedRange<Int>], [ClosedRange<Int>]) {
  var pass: [ClosedRange<Int>] = []
  var fail: [ClosedRange<Int>] = []
  if rule.operation == .GT {
    if range.last! > rule.number && range.first! > rule.number {
      pass += [range]
    } else if range.last! > rule.number {
      pass += [(rule.number + 1)...range.last!]
      fail += [range.first!...rule.number]
    } else {
      fail += [range]
    }
  } else {
    if range.first! < rule.number && range.last! < rule.number {
      pass += [range]
    } else if range.first! < rule.number {
      pass += [range.first!...(rule.number - 1)]
      fail += [rule.number...range.last!]
    } else {
      fail += [range]
    }
  }

  return (pass, fail)
}

func splitRanges(rangeParts: PartsRange, rule: Rule) -> ([PartsRange], [PartsRange]) {
  var passed: [PartsRange] = []
  var failed: [PartsRange] = []

  let range: ClosedRange<Int>
  switch rule.part {
  case .x: range = rangeParts.x
  case .m: range = rangeParts.m
  case .a: range = rangeParts.a
  case .s: range = rangeParts.s
  }

  let (passing, failing) = splitRange(range: range, rule: rule)

  for subrange in passing {
    var parts = rangeParts
    switch rule.part {
    case .x: parts.x = subrange
    case .m: parts.m = subrange
    case .a: parts.a = subrange
    case .s: parts.s = subrange
    }
    passed += [parts]
  }
  for subrange in failing {
    var parts = rangeParts
    switch rule.part {
    case .x: parts.x = subrange
    case .m: parts.m = subrange
    case .a: parts.a = subrange
    case .s: parts.s = subrange
    }
    failed += [parts]
  }

  return (passed, failed)
}

func findAcceptedRanges(_ rangeParts: PartsRange) -> [PartsRange] {
  var acceptedRanges: [PartsRange] = []
  var queue: [(String, PartsRange)] = [("in", rangeParts)]

  while queue.count > 0 {
    let (label, range) = queue.removeFirst()

    if label == "A" || label == "R" {
      acceptedRanges += label == "A" ? [range] : []
      continue
    }

    let workflow = workflows[label]!
    var processing: [PartsRange] = [range]

    for rule in workflow.rules {
      processing = processing.reduce(into: [PartsRange]()) { processing, subrange in
        let (passing, failing) = splitRanges(rangeParts: subrange, rule: rule)
        for passed in passing { queue += [(rule.next, passed)] }
        for failed in failing { processing += [failed] }
      }
    }

    for subrange in processing {
      queue += [(workflow.outcome, subrange)]
    }
  }

  return acceptedRanges
}

let sumAccepted = parts.filter { arePartsAccepted($0) }.reduce(0) {
  $0 + $1.x + $1.m + $1.a + $1.s
}

let acceptedRanges = findAcceptedRanges(
  PartsRange(x: 1...4000, m: 1...4000, a: 1...4000, s: 1...4000))

let sumPossibleAccepted = acceptedRanges.reduce(0) {
  $0 + ($1.x.count * $1.m.count * $1.a.count * $1.s.count)
}

print(
  """
  The sum of the accepted parts is \(sumAccepted).
  The sum of all possible accepted parts is \(sumPossibleAccepted).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
