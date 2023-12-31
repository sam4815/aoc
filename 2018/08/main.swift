import Foundation

let start = Date()

struct Node {
  var metadata: [Int]
  var children: [Node]
}

func parseNode(_ numbers: [Int]) -> (Node, Int) {
  var offset = 2
  var children: [Node] = []

  for _ in 0..<numbers[0] {
    let (child, length) = parseNode(Array(numbers[offset...]))
    children += [child]
    offset += length
  }

  let metadata = Array(numbers[offset..<(offset + numbers[1])])
  return (Node(metadata: metadata, children: children), offset + numbers[1])
}

func sumMetadata(_ node: Node) -> Int {
  node.metadata.reduce(0, +) + node.children.reduce(0, { $0 + sumMetadata($1) })
}

func findValue(_ node: Node) -> Int {
  if node.children.count == 0 {
    return node.metadata.reduce(0, +)
  }

  return node.metadata.reduce(0) { sum, value in
    if value == 0 || value > node.children.count { return sum }
    return sum + findValue(node.children[value - 1])
  }
}

let (node, _) = parseNode(
  try String(contentsOfFile: "input.txt").split(separator: "\n")[0].split(separator: " ").map {
    Int($0)!
  })

let metadataSum = sumMetadata(node)
let nodeValue = findValue(node)

print(
  """
  The sum of the node metadata is \(metadataSum).
  The value of the root node is \(nodeValue).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
