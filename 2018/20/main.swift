import Foundation

let start = Date()

struct RegexPath {
  var position: (Int, Int)
  var index: Int
}

struct DoorPath {
  var position: (Int, Int)
  var numDoors: Int
  var doorsPassed: [Int: [Int: Bool]]
}

let regex = try String(contentsOfFile: "input.txt").split(separator: "\n")[0].split(separator: "")
  .map { String($0) }

var visited: [Int: [Int: [Int: Bool]]] = [:]
var area: [Int: [Int: String]] = [:]
area[0, default: [:]][0] = "X"

func explore(_ regex: [String], initial: RegexPath) -> [RegexPath] {
  var paths = [initial]
  var subpaths: [(Int, Int)] = []

  while paths.count > 0 {
    let path = paths.removeFirst()

    if visited[path.position.0]?[path.position.1]?[path.index] ?? false {
      continue
    } else {
      visited[path.position.0, default: [:]][path.position.1, default: [:]][path.index] = true
    }

    if path.index >= regex.count {
      continue
    }

    switch regex[path.index] {
    case "|":
      subpaths += [path.position]
      paths += [RegexPath(position: initial.position, index: path.index + 1)]

    case ")":
      subpaths += [path.position]
      return subpaths.map { RegexPath(position: $0, index: path.index + 1) }

    case "(":
      paths += explore(
        regex, initial: RegexPath(position: path.position, index: path.index + 1))

    case "E", "W":
      let step = (regex[path.index] == "E" ? 1 : -1)
      area[path.position.0, default: [:]][path.position.1 + step] = "|"
      area[path.position.0, default: [:]][path.position.1 + step * 2] = "."
      paths += [
        RegexPath(
          position: (path.position.0, path.position.1 + step * 2), index: path.index + 1)
      ]

    case "N", "S":
      let step = (regex[path.index] == "S" ? 1 : -1)
      area[path.position.0 + step, default: [:]][path.position.1] = "-"
      area[path.position.0 + step * 2, default: [:]][path.position.1] = "."
      paths += [
        RegexPath(
          position: (path.position.0 + step * 2, path.position.1), index: path.index + 1)
      ]

    default:
      paths += [RegexPath(position: path.position, index: path.index + 1)]
    }
  }

  return [RegexPath(position: (0, 0), index: 0)]
}

func findMostDoors(_ area: [Int: [Int: String]]) -> (Int, Int) {
  var max = 0
  var min1000: [Int: [Int: Bool]] = [:]

  var queue = [DoorPath(position: (0, 0), numDoors: 0, doorsPassed: [:])]

  while queue.count > 0 {
    var path = queue.removeFirst()

    if path.numDoors > max {
      max = path.numDoors
    }

    if path.numDoors >= 1000 {
      min1000[path.position.0, default: [:]][path.position.1] = true
    }

    for offset in [(1, 0), (0, 1), (-1, 0), (0, -1)] {
      let doorPosition = (path.position.0 + offset.0, path.position.1 + offset.1)
      let doorCell = area[doorPosition.0]?[doorPosition.1] ?? "#"

      if doorCell != "|" && doorCell != "-" {
        continue
      }

      if path.doorsPassed[doorPosition.0]?[doorPosition.1] ?? false {
        continue
      } else {
        path.doorsPassed[doorPosition.0, default: [:]][doorPosition.1] = true
      }

      queue += [
        DoorPath(
          position: (path.position.0 + offset.0 * 2, path.position.1 + offset.1 * 2),
          numDoors: path.numDoors + 1, doorsPassed: path.doorsPassed)
      ]
    }
  }

  return (max, min1000.values.reduce(0) { $0 + $1.count })
}

let _ = explore(regex, initial: RegexPath(position: (0, 0), index: 0))
let (maxDoors, numMin1000) = findMostDoors(area)

print(
  """
  The largest number of doors required to reach any room is \(maxDoors).
  The number of rooms that are at least 1000 doors away is \(numMin1000).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
