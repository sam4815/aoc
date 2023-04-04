from time import time
from math import inf

start = time()

with open("input.txt", encoding="utf8") as file:
    initial_grid = [[*row] for row in file.read().splitlines()]


def get_cell(grid, position):
    if not 0 <= position[0] < len(grid) or not 0 <= position[1] < len(grid[0]):
        return "#"

    return grid[position[0]][position[1]]


def is_upper_case(cell):
    return 65 <= ord(cell) <= 90


def insert_path(paths, path_to_insert):
    for i, path in enumerate(paths):
        if path_to_insert["rec"] > path["rec"]:
            return paths[:i] + [path_to_insert] + paths[i:]

    return [path_to_insert] + paths


def get_moves(grid, portals, path):
    moves = []
    for direction in [[-1, 0], [1, 0], [0, -1], [0, 1]]:
        adjacent_position = (path["position"][0] + direction[0], path["position"][1] + direction[1])
        if get_cell(grid, adjacent_position) == ".":
            moves.append({"destination": adjacent_position, "rec": 0})

        if portals.get(path["position"]) is not None:
            moves.append(portals[path["position"]])

    return moves


def find_portal(grid, position):
    if not is_upper_case(get_cell(grid, position)):
        return None

    for direction in [[0, 1], [1, 0]]:
        before = (position[0] - direction[0], position[1] - direction[1])
        after = (position[0] + direction[0], position[1] + direction[1])
        after_after = (position[0] + direction[0] * 2, position[1] + direction[1] * 2)

        if get_cell(grid, before) == "." and is_upper_case(get_cell(grid, after)):
            portal_label = get_cell(grid, position) + get_cell(grid, after)
            return (portal_label, before)

        if is_upper_case(get_cell(grid, after)) and get_cell(grid, after_after) == ".":
            portal_label = get_cell(grid, position) + get_cell(grid, after)
            return (portal_label, after_after)

    return None


def find_label(grid, label, exclude=None):
    for i in range(len(grid)):
        for j in range(len(grid[0])):
            portal = find_portal(grid, (i, j))
            if portal is not None and portal[0] == label and portal[1] != exclude:
                return portal[1]

    return None


def find_portals(grid):
    portals = {}
    found_portals = ["AA", "ZZ"]

    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if 2 < i < len(grid) - 2 and 2 < j < len(grid[0]) - 2:
                continue

            portal = find_portal(grid, (i, j))
            if portal is not None and portal[0] not in found_portals:
                start = portal[1]
                end = find_label(grid, portal[0], portal[1])

                portals[start] = {"rec": -1, "destination": end}
                portals[end] = {"rec": 1, "destination": start}
                found_portals.append(portal[0])

    return portals


def find_shortest_path(grid, recursive=False):
    MAX_RECURSION_DEPTH = 25
    shortest_path = inf
    visited = {}
    portals = find_portals(grid)
    end = find_label(grid, "ZZ")

    path = {"position": find_label(grid, "AA"), "steps": 0, "rec": 0}
    path_queue = [path]

    while len(path_queue) > 0:
        curr_path, path_queue = path_queue[0], path_queue[1:]

        visited_key = curr_path["position"] + tuple([curr_path["rec"]])
        if visited.get(visited_key, inf) <= curr_path["steps"]:
            continue
        visited[visited_key] = curr_path["steps"]

        if curr_path["position"] == end and curr_path["rec"] == 0:
            shortest_path = curr_path["steps"]
            continue

        if curr_path["rec"] > MAX_RECURSION_DEPTH:
            continue

        for move in get_moves(grid, portals, curr_path):
            next_path = {"position": move["destination"], "steps": curr_path["steps"] + 1, "rec": curr_path["rec"]}

            if recursive:
                next_path["rec"] += move["rec"]
                if next_path["rec"] < 0:
                    continue

            path_queue = insert_path(path_queue, next_path)

    return shortest_path


min_num_steps = find_shortest_path(initial_grid)
min_num_recursive_steps = find_shortest_path(initial_grid, recursive=True)

time_elapsed = round(time() - start, 5)

print(
    f"""The shortest path from AA to ZZ takes {min_num_steps} steps.
The shortest path when the maze is recursive is {min_num_recursive_steps} steps.
Solution generated in {time_elapsed}s."""
)
