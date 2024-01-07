import Foundation

let start = Date()

let numRecipes = Int(try String(contentsOfFile: "input.txt").split(separator: "\n")[0])!
let pattern = try String(contentsOfFile: "input.txt").split(separator: "\n")[0]
  .split(separator: "").map { Int($0)! }

func endsWithPattern(_ numbers: [Int], pattern: [Int]) -> Bool {
  if numbers.count < pattern.count { return false }
  return pattern.enumerated().allSatisfy({ (index, digit) in
    let offset = pattern.count - 1 - index
    return digit == numbers[numbers.count - 1 - offset]
  })
}

var recipes = [3, 7]
var firstIndex = 0
var secondIndex = 1

while recipes.count < (numRecipes + 10) {
  let sum = recipes[firstIndex] + recipes[secondIndex]

  if sum >= 10 { recipes += [1] }
  recipes += [sum % 10]

  firstIndex = (firstIndex + recipes[firstIndex] + 1) % recipes.count
  secondIndex = (secondIndex + recipes[secondIndex] + 1) % recipes.count
}

let final10 = recipes[(recipes.count - 10)...].map { String($0) }.joined()

while !endsWithPattern(recipes, pattern: pattern) {
  let sum = recipes[firstIndex] + recipes[secondIndex]

  if sum >= 10 { recipes += [1] }
  if endsWithPattern(recipes, pattern: pattern) { break }
  recipes += [sum % 10]

  firstIndex = (firstIndex + recipes[firstIndex] + 1) % recipes.count
  secondIndex = (secondIndex + recipes[secondIndex] + 1) % recipes.count
}

let numRecipesBeforePattern = recipes.count - pattern.count

print(
  """
  The scores of the final 10 recipes is \(final10).
  The number of recipes that appear to the left of the score sequence in \(numRecipesBeforePattern).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
