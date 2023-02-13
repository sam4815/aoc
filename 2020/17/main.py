from time import time

start = time()

with open("input.txt", encoding="utf8") as file:
    three_dimensional_grid = {0: {0: {}}}
    four_dimensional_grid = {0: {0: {}}}

    grid_lines = file.read().strip().splitlines()
    for y, row in enumerate(grid_lines):
        y_midpoint = int((len(grid_lines) / 2) - 0.5)
        three_dimensional_grid[0][0][y - y_midpoint] = {}
        four_dimensional_grid[0][0][y - y_midpoint] = {}

        for x, cell in enumerate([*row]):
            x_midpoint = int((len(row) / 2) - 0.5)
            three_dimensional_grid[0][0][y - y_midpoint][x - x_midpoint] = cell
            four_dimensional_grid[0][0][y - y_midpoint][x - x_midpoint] = cell


def expand_grid(grid, expand_4th_dimension):
    if expand_4th_dimension:
        next_w_len, w_midpoint = len(grid) + 2, int((len(grid) / 2) + 0.5)
        next_grid = {i - w_midpoint: {} for i in range(next_w_len)}
    else:
        next_grid = {0: {}}

    for w in next_grid.keys():
        next_z_len, z_midpoint = len(grid[0]) + 2, int((len(grid[0]) / 2) + 0.5)
        next_grid[w] = {i - z_midpoint: {} for i in range(next_z_len)}

        for z in next_grid[w].keys():
            next_y_len = len(grid[0][0]) + 2
            y_midpoint = int((len(grid[0][0]) / 2) + 0.5)
            next_grid[w][z] = {i - y_midpoint: {} for i in range(next_y_len)}

            for y in next_grid[w][z].keys():
                next_x_len = len(grid[0][0][0]) + 2
                x_midpoint = int((len(grid[0][0][0]) / 2) + 0.5)
                next_grid[w][z][y] = {i - x_midpoint: {} for i in range(next_x_len)}

    return next_grid


def get_cube(x, y, z, w, grid):
    return grid.get(w, {}).get(z, {}).get(y, {}).get(x, ".")


def num_active_neighbours(x, y, z, w, grid):
    num_active = 0

    for m in range(w - 1, w + 2):
        for k in range(z - 1, z + 2):
            for j in range(y - 1, y + 2):
                for i in range(x - 1, x + 2):
                    if i == x and j == y and k == z and m == w:
                        continue

                    if get_cube(i, j, k, m, grid) == "#":
                        num_active += 1

    return num_active


def num_active_cubes(grid):
    num_active = 0
    for dimension in grid.values():
        for plane in dimension.values():
            for row in plane.values():
                num_active += sum(1 for cell in row.values() if cell == "#")
    return num_active


def step(grid, expand_4th_dimension):
    next_grid = expand_grid(grid, expand_4th_dimension)

    for w in next_grid.keys():
        for z in next_grid[w].keys():
            for y in next_grid[w][z].keys():
                for x in next_grid[w][z][y].keys():
                    is_active_cube = get_cube(x, y, z, w, grid) == "#"
                    num_neighbours = num_active_neighbours(x, y, z, w, grid)

                    if is_active_cube:
                        next_grid[w][z][y][x] = "#" if 2 <= num_neighbours <= 3 else "."
                    else:
                        next_grid[w][z][y][x] = "#" if num_neighbours == 3 else "."

    return next_grid


for _ in range(6):
    three_dimensional_grid = step(three_dimensional_grid, expand_4th_dimension=False)
    four_dimensional_grid = step(four_dimensional_grid, expand_4th_dimension=True)

num_3d_active = num_active_cubes(three_dimensional_grid)
num_4d_active = num_active_cubes(four_dimensional_grid)

time_elapsed = round(time() - start, 5)

print(
    f"""After six cycles there are {num_3d_active} cubes in 3-dimensional space.
After six cycles there are {num_4d_active} cubes in 4-dimensional space.
Solution generated in {time_elapsed}s."""
)
