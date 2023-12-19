import Foundation

let start = Date()

enum Direction {
  case U, D, L, R
}

struct Instruction {
  var direction: Direction
  var size: Int
  var colour: String
}

let lines = try String(contentsOfFile: "input.txt").split(separator: "\n")
let directionMap: [String: Direction] = [
  "0": .R, "1": .D, "2": .L, "3": .U, "R": .R, "D": .D, "L": .L, "U": .U,
]

let instructions = lines.map { line in
  let components = line.split(separator: " ")
  return Instruction(
    direction: directionMap[String(components[0])]!, size: Int(components[1])!,
    colour: components[2].trimmingCharacters(in: CharacterSet(charactersIn: "()#")))
}

let bigInstructions = lines.map {
  var colour = $0.split(separator: " ")[2].trimmingCharacters(in: CharacterSet(charactersIn: "()#"))
  let directionNum = String(colour.removeLast())
  return Instruction(
    direction: directionMap[directionNum]!, size: Int(colour, radix: 16)!, colour: colour)
}

public class InstructionNode {
  var value: Instruction
  var prev: InstructionNode?
  var next: InstructionNode?

  init(value: Instruction, prev: InstructionNode? = nil, next: InstructionNode? = nil) {
    self.value = value
    self.prev = prev
    self.next = next
  }

  func equals(_ compareTo: InstructionNode) -> Bool {
    return
      self.value.colour == compareTo.value.colour
      && self.value.direction == compareTo.value.direction
      && self.value.size == compareTo.value.size
  }

  func size() -> Int {
    var i = 1
    var next = self.next!
    while !self.equals(next) {
      i += 1
      next = next.next!
    }
    return i
  }

  func area() -> Int {
    (self.value.size + 1) * (self.next!.value.size + 1)
  }

  func combineWithNext() {
    self.value.size += self.next!.value.size
    self.next = self.next!.next!
    self.next!.prev = self
  }

  func counteractWithNext() -> Int {
    let next = self.next!
    let diff = min(self.value.size, next.value.size)
    switch (self.value.direction, next.value.direction) {
    case (.U, .D):
      if self.value.size > next.value.size {
        self.value.size -= next.value.size
      } else {
        self.value.size = next.value.size - self.value.size
        self.value.direction = .D
      }
    case (.D, .U):
      if self.value.size > next.value.size {
        self.value.size -= next.value.size
      } else {
        self.value.size = next.value.size - self.value.size
        self.value.direction = .U
      }
    case (.L, .R):
      if self.value.size > next.value.size {
        self.value.size -= next.value.size
      } else {
        self.value.size = next.value.size - self.value.size
        self.value.direction = .R
      }
    case (.R, .L):
      if self.value.size > next.value.size {
        self.value.size -= next.value.size
      } else {
        self.value.size = next.value.size - self.value.size
        self.value.direction = .L
      }
    default: ()
    }

    self.next = next.next!
    self.next!.prev = self

    return diff
  }

  func opposesNext() -> Bool {
    let next = self.next!
    return
      (self.value.direction == .L && next.value.direction == .R)
      || (self.value.direction == .R && next.value.direction == .L)
      || (self.value.direction == .U && next.value.direction == .D)
      || (self.value.direction == .D && next.value.direction == .U)
  }

  func print() {
    Swift.print(self.value)
    var next = self.next!
    while !self.equals(next) {
      Swift.print(next.value)
      next = next.next!
    }
  }
}

func convertToLinkedList(_ instructions: [Instruction]) -> InstructionNode {
  let initial = InstructionNode(value: instructions[0])
  var curr = initial
  for i in 1..<instructions.count {
    curr.next = InstructionNode(value: instructions[i], prev: curr)
    curr = curr.next!
  }
  curr.next = initial
  initial.prev = curr

  return initial
}

func reduceList(_ list: InstructionNode) -> Int {
  var reduced = list
  var area = 0
  let patterns = [
    [Direction.U, .R, .D], [Direction.R, .D, .L], [Direction.D, .L, .U], [Direction.L, .U, .R],
  ]

  while reduced.size() > 4 {
    for pattern in patterns {
      if reduced.size() <= 4 {
        break
      }
      if reduced.value.direction == pattern[0] && reduced.next!.value.direction == pattern[1]
        && reduced.next!.next!.value.direction == pattern[2]
      {
        let up = reduced
        var mid = reduced.next!
        let down = reduced.next!.next!
        let height = min(up.value.size, down.value.size)

        area += height * (mid.value.size + 1)

        if up.value.size == height {
          let preUp = up.prev!
          preUp.next = mid
          mid.prev = preUp

          if preUp.value.direction == mid.value.direction {
            preUp.combineWithNext()
            mid = preUp
          } else if preUp.opposesNext() {
            area += preUp.counteractWithNext()
            mid = preUp
          }
        } else {
          up.value.size -= height
        }

        if down.value.size == height {
          let postDown = down.next!
          mid.next = postDown
          postDown.prev = mid

          if mid.value.direction == postDown.value.direction {
            mid.combineWithNext()
          } else if mid.opposesNext() {
            area += mid.counteractWithNext()
          }
        } else {
          down.value.size -= height
        }

        reduced = mid
      }
    }

    reduced = reduced.next!
  }

  return area + reduced.area()
}

let smallLagoonVolume = reduceList(convertToLinkedList(instructions))
let bigLagoonVolume = reduceList(convertToLinkedList(bigInstructions))

print(
  """
  The small lagoon can hold \(smallLagoonVolume) cubic metres of lava.
  The big lagoon can hold \(bigLagoonVolume) cubic metres of lava.
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
