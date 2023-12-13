import Foundation

let start = Date()

let patterns = try String(contentsOfFile: "input.txt").split(separator: "\n\n").map { pattern in
  pattern.split(separator: "\n").map { $0.split(separator: "") }
}

func matrixTranspose<T>(_ matrix: [[T]]) -> [[T]] {
  (0..<matrix.first!.count).reduce(into: [[T]]()) { result, index in
    result += [matrix.map { $0[index] }]
  }
}

func findReflectionRow(_ pattern: [[Substring]], allowImperfections: Bool = false) -> Int {
  reflectionLoop: for x in 0..<(pattern.count - 1) {
    var imperfection = 0
    let reflectionHeight = min(x + 1, pattern.count - x - 1)

    for z in 0..<reflectionHeight {
      for y in 0..<pattern[0].count {
        if pattern[x - z][y] != pattern[x + z + 1][y] {
          if imperfection != 0 {
            continue reflectionLoop
          }
          imperfection = (x + 1)
        }
      }
    }

    if allowImperfections && imperfection != 0 {
      return imperfection
    }
    if !allowImperfections && imperfection == 0 {
      return (x + 1)
    }
  }

  return 0
}

func findReflectionValue(_ pattern: [[Substring]], allowImperfections: Bool = false) -> Int {
  findReflectionRow(pattern, allowImperfections: allowImperfections) * 100
    + findReflectionRow(matrixTranspose(pattern), allowImperfections: allowImperfections)
}

let patternSum = patterns.reduce(0, { $0 + findReflectionValue($1) })
let patternSmudgeSum = patterns.reduce(
  0, { $0 + findReflectionValue($1, allowImperfections: true) })

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The sum of the reflections is \(patternSum).
  After removing smudges, the sum of the reflections is \(patternSmudgeSum).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
