import Foundation

let start = Date()

let contents = try String(contentsOfFile: "input.txt").split(separator: "\n")

func getCharCounts(str: Substring) -> [Int] {
  return str.split(separator: "").reduce(
    into: [Substring: Int](),
    { charMap, char in
      charMap[char] = (charMap[char] ?? 0) + 1
    }
  ).map({ $0.value })
}

let hasTwiceAppearing = contents.filter { getCharCounts(str: $0).contains(2) }.count
let hasThriceAppearing = contents.filter { getCharCounts(str: $0).contains(3) }.count
let checksum = hasTwiceAppearing * hasThriceAppearing

func findSimilarLetters(strs: [[Substring]]) -> String {
  for (index, strA) in strs.enumerated() {
    outer: for strB in strs[(index + 1)...] {
      var diffIndex = -1

      for i in 0..<strA.count {
        if strA[i] != strB[i] && diffIndex >= 0 {
          continue outer
        } else if strA[i] != strB[i] {
          diffIndex = i
        }
      }

      return strA.enumerated().filter({ $0.offset != diffIndex }).map { $0.element }.joined()
    }
  }
  return ""
}

let commonLetters = findSimilarLetters(strs: contents.map { $0.split(separator: "") })

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The checksum of the box IDs is \(checksum).
  The common letters between the two correct box IDs are \(commonLetters).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
