import Foundation

let start = Date()

enum Opcode: String {
  case addr, addi, mulr, muli, banr, bani, borr,
    bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr
}

struct Instruction {
  var opcode: Opcode
  var a: Int
  var b: Int
  var c: Int
}

let sections = try String(contentsOfFile: "input.txt").split(separator: "\n")

let instuctionPointer = Int(sections[0].split(separator: " ")[1])!
let program = sections[1...].map {
  let components = $0.split(separator: " ")
  let numbers = components[1...].map { Int($0)! }
  return Instruction(
    opcode: Opcode(rawValue: String(components[0]))!, a: numbers[0], b: numbers[1], c: numbers[2])
}

func findFactors(_ num: Int) -> [Int] {
  (1...num).filter { num % $0 == 0 }
}

func executeInstruction(_ state: [Int], instruction: Instruction) -> [Int] {
  var result = state

  switch instruction.opcode {
  case .addr: result[instruction.c] = state[instruction.a] + state[instruction.b]
  case .addi: result[instruction.c] = state[instruction.a] + instruction.b
  case .mulr: result[instruction.c] = state[instruction.a] * state[instruction.b]
  case .muli: result[instruction.c] = state[instruction.a] * instruction.b
  case .banr: result[instruction.c] = state[instruction.a] & state[instruction.b]
  case .bani: result[instruction.c] = state[instruction.a] & instruction.b
  case .borr: result[instruction.c] = state[instruction.a] | state[instruction.b]
  case .bori: result[instruction.c] = state[instruction.a] | instruction.b
  case .setr: result[instruction.c] = state[instruction.a]
  case .seti: result[instruction.c] = instruction.a
  case .gtir: result[instruction.c] = (instruction.a > state[instruction.b] ? 1 : 0)
  case .gtri: result[instruction.c] = (state[instruction.a] > instruction.b ? 1 : 0)
  case .gtrr: result[instruction.c] = (state[instruction.a] > state[instruction.b] ? 1 : 0)
  case .eqir: result[instruction.c] = (instruction.a == state[instruction.b] ? 1 : 0)
  case .eqri: result[instruction.c] = (state[instruction.a] == instruction.b ? 1 : 0)
  case .eqrr: result[instruction.c] = (state[instruction.a] == state[instruction.b] ? 1 : 0)
  }

  return result
}

func executeProgram(_ program: [Instruction], initialState: [Int], ipRegister: Int) -> [Int] {
  var state = initialState
  var ip = 0

  while ip < program.count && state[4] < 100000 {
    let instruction = program[ip]
    state[ipRegister] = ip

    state = executeInstruction(state, instruction: instruction)

    ip = state[ipRegister]
    ip += 1
  }

  return state
}

let firstResult = executeProgram(
  program, initialState: [0, 0, 0, 0, 0, 0], ipRegister: instuctionPointer)[0]

let partialSecondResult = executeProgram(
  program, initialState: [1, 0, 0, 0, 0, 0], ipRegister: instuctionPointer)

let secondResult = findFactors(partialSecondResult[4]).reduce(0, +)

print(
  """
  After executing the first program, the value in register 0 is \(firstResult).
  After executing the second program, the value in register 0 is \(secondResult).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
