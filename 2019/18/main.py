from time import time
from math import inf

start = time()

DATA_CORRECTION = [
    ["@", "#", "@"],
    ["#", "#", "#"],
    ["@", "#", "@"],
]

with open("input.txt", encoding="utf8") as file:
    initial_grid = [[*row] for row in file.read().strip().splitlines()]


def is_key(cell):
    return 97 <= ord(cell) <= 122


def is_door(cell):
    return 65 <= ord(cell) <= 90


def curr_cell(grid, position):
    return grid[position[0]][position[1]]


def num_keys(grid):
    return sum(sum(1 for cell in row if is_key(cell)) for row in grid)


def find_starting_positions(grid):
    starting_positions = []
    for i, row in enumerate(grid):
        for j, cell in enumerate(row):
            if cell == "@":
                starting_positions.append((i, j))
    return starting_positions


def find_key_positions(grid):
    keys = {}
    for i, row in enumerate(grid):
        for j, cell in enumerate(row):
            if is_key(cell):
                keys[cell] = (i, j)
    return keys


def has_required_keys(full_path, key_path):
    return all(key in full_path["collected_keys"] for key in key_path["required_keys"])


def insert_path(paths, path_to_insert):
    for i, path in enumerate(paths):
        if len(path["collected_keys"]) < len(path_to_insert["collected_keys"]):
            return paths[:i] + [path_to_insert] + paths[i:]

    return [path_to_insert] + paths


def insert_key(keys, key_to_insert):
    for i, key in enumerate(keys):
        if key > key_to_insert:
            return keys[:i] + [key_to_insert] + keys[i:]

    return keys + [key_to_insert]


def get_moves(grid, path):
    moves = []
    for offset in [[-1, 0], [1, 0], [0, -1], [0, 1]]:
        offset_position = (path["position"][0] + offset[0], path["position"][1] + offset[1])
        if offset_position[0] in [-1, len(grid)] or offset_position[1] in [-1, len(grid[0])]:
            continue
        if curr_cell(grid, offset_position) == "#":
            continue

        moves.append(offset_position)

    return moves


def find_shortest_paths_to_keys(grid, position):
    shortest_paths = {}
    visited = {}

    path = {
        "position": position,
        "required_keys": [],
        "steps": 0,
    }
    path_queue = [path]

    while len(path_queue) > 0:
        curr_path, path_queue = path_queue[0], path_queue[1:]

        if visited.get(curr_path["position"], inf) < curr_path["steps"]:
            continue

        visited[curr_path["position"]] = curr_path["steps"]

        if is_key(curr_cell(grid, curr_path["position"])):
            key = curr_cell(grid, curr_path["position"])
            if shortest_paths.get(key) is None or shortest_paths[key]["steps"] > curr_path["steps"]:
                shortest_paths[key] = {
                    "steps": curr_path["steps"],
                    "required_keys": curr_path["required_keys"][:],
                }

        if is_door(curr_cell(grid, curr_path["position"])):
            curr_path["required_keys"].append(curr_cell(grid, curr_path["position"]).lower())

        for move in get_moves(grid, curr_path):
            next_path = {
                "position": move,
                "required_keys": curr_path["required_keys"][:],
                "steps": curr_path["steps"] + 1,
            }

            path_queue.insert(0, next_path)

    return shortest_paths


def find_shortest_path(grid):
    shortest_path = inf
    visited = {}

    key_positions = find_key_positions(grid)
    key_paths = {}

    for key, position in key_positions.items():
        key_paths[key] = find_shortest_paths_to_keys(grid, position)

    for i, start_pos in enumerate(find_starting_positions(grid)):
        key_paths[i] = find_shortest_paths_to_keys(grid, start_pos)

    path_queue = [
        {
            "curr_positions": list(range(len(find_starting_positions(grid)))),
            "collected_keys": [],
            "steps": 0,
        }
    ]

    while len(path_queue) > 0:
        curr_path, path_queue = path_queue[0], path_queue[1:]

        if len(curr_path["collected_keys"]) == len(key_positions) and curr_path["steps"] < shortest_path:
            shortest_path = curr_path["steps"]
            continue

        if curr_path["steps"] > shortest_path:
            continue

        visited_key = tuple(curr_path["curr_positions"]) + (0, 0) + tuple(curr_path["collected_keys"])
        if visited.get(visited_key, inf) <= curr_path["steps"]:
            continue
        visited[visited_key] = curr_path["steps"]

        for i, curr_position in enumerate(curr_path["curr_positions"]):
            for key, key_path in key_paths[curr_position].items():
                if key in curr_path["collected_keys"]:
                    continue
                if not has_required_keys(curr_path, key_path):
                    continue

                next_keys = curr_path["curr_positions"][:]
                next_keys[i] = key

                next_path = {
                    "curr_positions": next_keys,
                    "collected_keys": insert_key(curr_path["collected_keys"], key),
                    "steps": curr_path["steps"] + key_path["steps"],
                }
                path_queue = insert_path(path_queue, next_path)

    return shortest_path


min_num_steps = find_shortest_path(initial_grid)

divided_grid = [row[:] for row in initial_grid]
starting_position = find_starting_positions(initial_grid)[0]

for x_offset, row in enumerate(DATA_CORRECTION):
    for y_offset, cell in enumerate(row):
        x = starting_position[0] - 1 + x_offset
        y = starting_position[1] - 1 + y_offset
        divided_grid[x][y] = cell


min_num_robot_steps = find_shortest_path(divided_grid)

time_elapsed = round(time() - start, 5)

print(
    f"""The shortest path to collect all keys individually takes {min_num_steps} steps.
The shortest path to collect all keys using the robots takes {min_num_robot_steps} steps.
Solution generated in {time_elapsed}s."""
)
