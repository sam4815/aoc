import Foundation

let start = Date()

enum Opcode {
  case addr, addi, mulr, muli, banr, bani, borr,
    bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr
}

struct Instruction {
  var opcode: Int
  var a: Int
  var b: Int
  var c: Int
}

struct StateSample {
  var before: [Int]
  var instruction: Instruction
  var after: [Int]
}

func parseInts(_ line: Substring) -> [Int] {
  line.components(separatedBy: CharacterSet.decimalDigits.inverted).filter({ !$0.isEmpty }).map({
    Int($0)!
  })
}

let sections = try String(contentsOfFile: "input.txt").split(separator: "\n\n\n\n")

let samples = sections[0].split(separator: "\n\n").map {
  let lines = $0.split(separator: "\n").map { parseInts($0) }
  return StateSample(
    before: lines[0],
    instruction: Instruction(opcode: lines[1][0], a: lines[1][1], b: lines[1][2], c: lines[1][3]),
    after: lines[2])
}

let program = sections[1].split(separator: "\n").map {
  let line = parseInts($0)
  return Instruction(opcode: line[0], a: line[1], b: line[2], c: line[3])
}

func executeInstruction(_ state: [Int], instruction: Instruction, opcodes: [Int: Opcode]) -> [Int] {
  let registerA = state[instruction.a]
  let registerB = state[instruction.b]
  let valueA = instruction.a
  let valueB = instruction.b
  var result = state

  switch opcodes[instruction.opcode]! {
  case .addr: result[instruction.c] = registerA + registerB
  case .addi: result[instruction.c] = registerA + valueB
  case .mulr: result[instruction.c] = registerA * registerB
  case .muli: result[instruction.c] = registerA * valueB
  case .banr: result[instruction.c] = registerA & registerB
  case .bani: result[instruction.c] = registerA & valueB
  case .borr: result[instruction.c] = registerA | registerB
  case .bori: result[instruction.c] = registerA | valueB
  case .setr: result[instruction.c] = registerA
  case .seti: result[instruction.c] = valueA
  case .gtir: result[instruction.c] = (valueA > registerB ? 1 : 0)
  case .gtri: result[instruction.c] = (registerA > valueB ? 1 : 0)
  case .gtrr: result[instruction.c] = (registerA > registerB ? 1 : 0)
  case .eqir: result[instruction.c] = (valueA == registerB ? 1 : 0)
  case .eqri: result[instruction.c] = (registerA == valueB ? 1 : 0)
  case .eqrr: result[instruction.c] = (registerA == registerB ? 1 : 0)
  }

  return result
}

func testState(_ sample: StateSample) -> [Opcode] {
  var possible: [Opcode] = []
  let naiveOpcodes: [Int: Opcode] = [
    0: .addr, 1: .addi, 2: .mulr, 3: .muli, 4: .banr, 5: .bani, 6: .borr, 7: .bori,
    8: .setr, 9: .seti, 10: .gtir, 11: .gtri, 12: .gtrr, 13: .eqir, 14: .eqri, 15: .eqrr,
  ]

  for (i, opcode) in naiveOpcodes {
    let instruction = Instruction(
      opcode: i, a: sample.instruction.a, b: sample.instruction.b, c: sample.instruction.c)
    let result = executeInstruction(sample.before, instruction: instruction, opcodes: naiveOpcodes)

    if sample.after == result {
      possible += [opcode]
    }
  }

  return possible
}

func determineOpcodes(_ samples: [StateSample]) -> [Int: Opcode] {
  var opcodes: [Int: Opcode] = [:]

  while opcodes.count != 16 {
    for sample in samples {
      let possibleOpcodes = testState(sample).filter { !opcodes.values.contains($0) }
      if possibleOpcodes.count == 1 {
        opcodes[sample.instruction.opcode] = possibleOpcodes[0]
      }
    }
  }

  return opcodes
}

func executeProgram(_ program: [Instruction], samples: [StateSample]) -> [Int] {
  let opcodes = determineOpcodes(samples)
  var state = [0, 0, 0, 0]

  for instruction in program {
    state = executeInstruction(state, instruction: instruction, opcodes: opcodes)
  }

  return state
}

let multiOpcodesCount = samples.filter { testState($0).count >= 3 }.count
let register0 = executeProgram(program, samples: samples)[0]

print(
  """
  The number of samples that behave like 3 or more opcodes is \(multiOpcodesCount).
  After executing the test program, the value in register 0 is \(register0).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
