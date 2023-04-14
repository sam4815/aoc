from time import time

start = time()

with open("input.txt", encoding="utf8") as file:
    initial_grid = [[[*row] for row in file.read().splitlines()]]


LAYER_MAPPING = {
    (2, 1): [(0, 0), (1, 0), (2, 0), (3, 0), (4, 0)],
    (1, 2): [(0, 0), (0, 1), (0, 2), (0, 3), (0, 4)],
    (2, 3): [(0, 4), (1, 4), (2, 4), (3, 4), (4, 4)],
    (3, 2): [(4, 0), (4, 1), (4, 2), (4, 3), (4, 4)],
}


def count_bugs(grid, position):
    count = 0

    layer_position = (position[1], position[2])
    for inner, outer in LAYER_MAPPING.items():
        if inner == layer_position and position[0] + 1 < len(grid):
            for inner_position in outer:
                count += 1 if grid[position[0] + 1][inner_position[0]][inner_position[1]] == "#" else 0

        if layer_position in outer and position[0] - 1 >= 0:
            count += 1 if grid[position[0] - 1][inner[0]][inner[1]] == "#" else 0

    for offset in [[0, 1], [1, 0], [0, -1], [-1, 0]]:
        offset_position = [layer_position[0] + offset[0], layer_position[1] + offset[1]]
        if 0 <= offset_position[0] < len(grid[0]) and 0 <= offset_position[1] < len(grid[0]):
            count += 1 if grid[position[0]][offset_position[0]][offset_position[1]] == "#" else 0

    return count


def new_layer():
    return [["." for _ in range(5)] for _ in range(5)]


def layer_has_bug(layer):
    return any(any(cell == "#" for cell in row) for row in layer)


def tick(grid, recursive=False):
    if recursive and layer_has_bug(grid[0]):
        grid.insert(0, new_layer())
    if recursive and layer_has_bug(grid[-1]):
        grid.append(new_layer())

    next_grid = [[[] for _ in layer] for layer in grid]

    for z, layer in enumerate(grid):
        for y, row in enumerate(layer):
            for x, cell in enumerate(row):
                if recursive and (x, y) == (2, 2):
                    next_grid[z][y].append(".")
                    continue

                num_adjacent_bugs = count_bugs(grid, [z, y, x])
                if cell == "#" and num_adjacent_bugs != 1:
                    next_grid[z][y].append(".")
                elif cell == "." and 1 <= num_adjacent_bugs <= 2:
                    next_grid[z][y].append("#")
                else:
                    next_grid[z][y].append(cell)

    return next_grid


def get_biodiversity(grid):
    return sum(sum(2 ** (x + y * len(grid)) for x, cell in enumerate(row) if cell == "#") for y, row in enumerate(grid))


def count_all_bugs(grid):
    return sum(sum(sum(1 for cell in row if cell == "#") for row in level) for level in grid)


def find_repeating_pattern(grid):
    grid_history = [grid]

    while True:
        next_grid = tick(grid_history[-1])
        if next_grid in grid_history:
            return next_grid

        grid_history.append(next_grid)


def find_recursive_bugs(grid, n):
    for _ in range(n):
        grid = tick(grid, recursive=True)
    return grid


bio_rating = get_biodiversity(find_repeating_pattern(initial_grid)[0])
recursive_bug_count = count_all_bugs(find_recursive_bugs(initial_grid, 200))

time_elapsed = round(time() - start, 5)


print(
    f"""The biodiversity rating for the first layout that appears twice is {bio_rating}.
After 200 minutes, {recursive_bug_count} bugs are present.
Solution generated in {time_elapsed}s."""
)
