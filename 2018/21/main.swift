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

let program = try String(contentsOfFile: "input.txt").split(separator: "\n")[1...].map {
  let components = $0.split(separator: " ")
  let numbers = components[1...].map { Int($0)! }
  return Instruction(
    opcode: Opcode(rawValue: String(components[0]))!, a: numbers[0], b: numbers[1], c: numbers[2])
}

let bitmaskInstructionIndex = program.firstIndex(where: { $0.opcode == .bori })!
let targetInstruction = program[bitmaskInstructionIndex + 1]
let mulInstruction = program.first(where: { $0.opcode == .muli })!

var targets: [Int] = []
var target = targetInstruction.a
var accumulator = 0x10000

while true {
  target += (accumulator & 0xFF)
  target *= mulInstruction.b
  target &= 0xFFFFFF

  if accumulator < 256 {
    if targets.contains(target) {
      break
    } else {
      targets += [target]
    }

    accumulator = target | 0x10000
    target = targetInstruction.a
  } else {
    accumulator /= 256
  }
}

let minInstructions = targets.first!
let maxInstructions = targets.last!

print(
  """
  The smallest value that causes the program to halt with the fewest instructions is \(minInstructions).
  The smallest value that causes the program to halt with the most instructions is \(maxInstructions).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)

// seti 0 1 4         r4 = 0
// bori 4 65536 1     r1 = r4 | 65536
// seti 16031208 7 4  r4 = 16031208

// bani 1 255 3       r3 = r1 & 255
// addr 4 3 4         r4 += r3
// bani 4 16777215 4  r4 &= 16777215
// muli 4 65899 4     r4 *= 65899
// bani 4 16777215 4  r4 &= 16777215

// gtir 256 1 3       if 256 > r1,
// addr 3 2 2           go to seti 27 3 2
// addi 2 1 2           else go to seti 0 9 3
// seti 27 3 2        go to eqrr 4 0 3

// seti 0 9 3         r3 = 0
// addi 3 1 5         r5 += 1
// muli 5 256 5       r5 *= 256
// gtrr 5 1 5         if r5 > r1,
// addr 5 2 2           go to addi 3 1 3
// addi 2 1 2           else go to seti 25 7 2

// seti 25 7 2        go to setr 3 1 1

// addi 3 1 3         r3 += 1
// seti 17 4 2        go to addi 3 1 5

// setr 3 1 1         r1 = r3
// seti 7 5 2         go to bani 1 255 3

// eqrr 4 0 3         if r0 == r4,
// addr 3 2 2            terminate program
// seti 5 1 2            else go to bori 4 65536 1
