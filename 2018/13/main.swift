import Foundation

let start = Date()

enum Turn { case Left, Straight, Right }
enum Direction { case Up, Left, Right, Down }

struct Cart {
  var position: (Int, Int)
  var nextTurn: Turn
  var direction: Direction
}

let (carts, tracks) = try String(contentsOfFile: "input.txt").split(separator: "\n").enumerated()
  .reduce(into: ([Cart](), [[Substring]]())) { items, line in
    var row = line.element.split(separator: "")

    for (y, cell) in row.enumerated() {
      if !["<", ">", "^", "v"].contains(cell) { continue }

      let direction: Direction
      switch cell {
      case "<": direction = .Left
      case "^": direction = .Up
      case "v": direction = .Down
      default: direction = .Right
      }

      items.0 += [Cart(position: (line.offset, y), nextTurn: .Left, direction: direction)]
      row[y] = [">", "<"].contains(cell) ? "-" : "|"
    }

    items.1 += [row]
  }

func getNextPosition(_ position: (Int, Int), direction: Direction) -> (Int, Int) {
  switch direction {
  case .Up: return (position.0 - 1, position.1)
  case .Right: return (position.0, position.1 + 1)
  case .Down: return (position.0 + 1, position.1)
  case .Left: return (position.0, position.1 - 1)
  }
}

func getNextDirection(cell: Substring, direction: Direction, nextTurn: Turn) -> Direction {
  switch (cell, direction, nextTurn) {
  case ("\\", .Left, _): return .Up
  case ("\\", .Up, _): return .Left
  case ("\\", .Right, _): return .Down
  case ("\\", .Down, _): return .Right

  case ("/", .Left, _): return .Down
  case ("/", .Up, _): return .Right
  case ("/", .Right, _): return .Up
  case ("/", .Down, _): return .Left

  case ("+", .Left, .Left): return .Down
  case ("+", .Left, .Right): return .Up
  case ("+", .Left, .Straight): return .Left

  case ("+", .Up, .Left): return .Left
  case ("+", .Up, .Right): return .Right
  case ("+", .Up, .Straight): return .Up

  case ("+", .Right, .Left): return .Up
  case ("+", .Right, .Right): return .Down
  case ("+", .Right, .Straight): return .Right

  case ("+", .Down, .Left): return .Right
  case ("+", .Down, .Right): return .Left
  case ("+", .Down, .Straight): return .Down

  default: return direction
  }
}

func getNextTurn(_ turn: Turn) -> Turn {
  switch turn {
  case .Left: return .Straight
  case .Straight: return .Right
  case .Right: return .Left
  }
}

func sortByPosition(cartA: Cart, cartB: Cart) -> Bool {
  if cartA.position.0 == cartB.position.0 {
    return cartA.position.1 < cartB.position.1
  } else {
    return cartA.position.0 < cartB.position.0
  }
}

func tick(_ carts: [Cart]) -> [Cart] {
  carts.sorted(by: { sortByPosition(cartA: $0, cartB: $1) }).map { cart in
    let position = getNextPosition(cart.position, direction: cart.direction)

    let direction = getNextDirection(
      cell: tracks[position.0][position.1], direction: cart.direction, nextTurn: cart.nextTurn)

    let nextTurn =
      tracks[position.0][position.1] == "+" ? getNextTurn(cart.nextTurn) : cart.nextTurn

    return Cart(position: position, nextTurn: nextTurn, direction: direction)
  }
}

func getPositionsMap(_ carts: [Cart]) -> [Int: [Int: Int]] {
  carts.enumerated().reduce(into: [Int: [Int: Int]]()) { positions, cart in
    positions[cart.element.position.0, default: [:]][cart.element.position.1] = cart.offset
  }
}

func findCrashedCarts(_ carts: [Cart], positions: inout [Int: [Int: Int]]) -> [Cart] {
  for (i, cart) in carts.enumerated() {
    if positions[cart.position.0, default: [:]][cart.position.1] != nil {
      return [cart, carts[positions[cart.position.0]![cart.position.1]!]]
    } else {
      positions[cart.position.0, default: [:]][cart.position.1] = i
    }
  }

  return []
}

func simulateCarts(_ initialCarts: [Cart], exitOnCrash: Bool) -> Cart {
  var carts = initialCarts

  while carts.count > 1 {
    var positions: [Int: [Int: Int]] = getPositionsMap(carts)

    carts = tick(carts)

    let crashedCarts = findCrashedCarts(carts, positions: &positions)

    carts = carts.filter { cart in
      !crashedCarts.contains(where: {
        $0.position.0 == cart.position.0 && $0.position.1 == cart.position.1
      })
    }

    if exitOnCrash && crashedCarts.count > 0 {
      return crashedCarts[0]
    }
  }

  return carts[0]
}

let crashedCart = simulateCarts(carts, exitOnCrash: true)
let finalCart = simulateCarts(carts, exitOnCrash: false)

print(
  """
  The location of the first crash is \(crashedCart.position.1),\(crashedCart.position.0).
  The location of the final cart is \(finalCart.position.1),\(finalCart.position.0).
  Solution generated in \(String(format: "%.4f", -start.timeIntervalSinceNow))s.
  """)
