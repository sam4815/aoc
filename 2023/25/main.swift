import Foundation

let start = Date()

var graph = try String(contentsOfFile: "input.txt").split(separator: "\n").reduce(
  into: [String: [String: Int]]()
) { components, line in
  let name = String(line.split(separator: ": ")[0])
  let connected = line.split(separator: ": ")[1].split(separator: " ").map { String($0) }

  for comp in connected {
    graph[name, default: [:]][comp] = 1
    graph[comp, default: [:]][name] = 1
  }
}

func findPath(_ graph: [String: [String: Int]], source: String, sink: String) -> [String: String] {
  var path = [source: source]
  var queue = [source]

  while queue.count > 0 {
    let node = queue.removeFirst()
    for (edge, capacity) in graph[node]! {
      if capacity > 0 && path[edge] == nil {
        path[edge] = node
        queue += [edge]
      }
    }
  }

  return path
}

func findMinCut(_ graph: [String: [String: Int]], source: String, sink: String) -> (Int, Int) {
  var subgraph = graph
  var maxFlow = 0

  var path = findPath(subgraph, source: source, sink: sink)

  while path[sink] != nil && sink != source {
    var flow = Int.max
    var node = sink

    while node != source {
      flow = min(flow, subgraph[path[node]!]![node]!)
      subgraph[path[node]!]![node]! -= flow
      subgraph[node]![path[node]!]! += flow

      node = path[node]!
    }

    maxFlow += flow
    path = findPath(subgraph, source: source, sink: sink)
  }

  return (maxFlow, path.count)
}

let source = graph.first!.key
let sink = graph.keys.first(where: { findMinCut(graph, source: source, sink: $0).0 == 3 })!

let groupSize = findMinCut(graph, source: source, sink: sink).1
let componentsProduct = groupSize * (graph.count - groupSize)

print(
  """
  The product of the component groups is \(componentsProduct).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
