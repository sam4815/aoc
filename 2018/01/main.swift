import Foundation

let start = Date()

let frequencies = try String(contentsOfFile: "input.txt").split(separator: "\n").map { Int($0)! }
let frequencySum = frequencies.reduce(0, +)

func findFirstDuplicate(nums: [Int]) -> Int {
  var seen = [Int: Bool]()
  var sum = 0
  var currentIndex = 0

  while true {
    sum += nums[currentIndex % nums.count]
    if seen[sum] != nil {
      return sum
    } else {
      seen[sum] = true
      currentIndex += 1
    }
  }
}

let firstDuplicate = findFirstDuplicate(nums: frequencies)

let timeElapsed = start.timeIntervalSinceNow * -1

print(
  """
  After applying all changes, the resulting frequency is \(frequencySum).
  The first frequency reached twice is \(firstDuplicate).
  Solution generated in \(String(format: "%.4f", timeElapsed))s.
  """)
