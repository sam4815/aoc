import Foundation

extension Substring {
  var isNumber: Bool {
    return self.allSatisfy { character in character.isNumber }
  }
}

let start = Date()

let contents = try String(contentsOfFile: "input.txt")
let contentsArray = Array(contents)

let width = contents.firstIndex(of: "\n")!.utf16Offset(in: contents)
let numbers = contents.matches(of: /\d+/)

let numsToCharMap = numbers.reduce(into: [(Int, [Character: [Int]])]()) { charMap, number in
  let start = number.range.lowerBound.utf16Offset(in: contents)
  let end = number.range.upperBound.utf16Offset(in: contents)

  let adjacentIndices =
    [start - 1, end]
    + Array((start - width - 2)...(end - (width + 1)))
    + Array((start + width)...(end + (width + 1)))

  let specialCharMap = adjacentIndices.reduce(into: [Character: [Int]]()) {
    chars, index in
    if index < 0 || index >= contentsArray.count { return }

    let char = contentsArray[index]
    if char == "." || char.isNumber || char == "\n" { return }

    chars[char] = (chars[char] ?? []) + [index]
  }

  charMap.append((Int(number.0)!, specialCharMap))
}

let sumPartNumbers = numsToCharMap.map { key, value in value.count > 0 ? key : 0 }.reduce(0, +)

let gearsToNumsMap = numsToCharMap.reduce(into: [Int: [Int]]()) { gearsMap, value in
  let (number, charMap) = value
  for position in (charMap["*"] ?? []) {
    gearsMap[position] = (gearsMap[position] ?? []) + [number]
  }
}

let sumGearRatios = gearsToNumsMap.values.reduce(0) { sum, numbers in
  sum + (numbers.count == 2 ? numbers.reduce(1, *) : 0)
}

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  The sum of the part numbers is \(sumPartNumbers).
  The sum of the gear ratios is \(sumGearRatios).
  Solution generated in \(timeElapsed)s.
  """)
