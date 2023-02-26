from time import time
from re import search

start = time()

with open("input.txt", encoding="utf8") as file:
    tiles = [
        {
            "id": search(r"\d+", tile).group(),
            "pattern": [[*row] for row in tile.splitlines()[1:]],
        }
        for tile in file.read().strip().split("\n\n")
    ]

    grid_length = int(len(tiles) ** (1 / 2))

SEA_MONSTER = [
    [" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "#", " "],
    ["#", " ", " ", " ", " ", "#", "#", " ", " ", " ", " ", "#", "#", " ", " ", " ", " ", "#", "#", "#"],
    [" ", "#", " ", " ", "#", " ", " ", "#", " ", " ", "#", " ", " ", "#", " ", " ", "#", " ", " ", " "],
]


def flipVertical(tile):
    return tile[::-1]


def flipHorizontal(tile):
    return [row[::-1] for row in tile]


def rotate90(tile):
    return flipHorizontal([[row[i] for row in tile] for i, _ in enumerate(tile)])


def permutations(tile):
    return [
        tile,
        rotate90(tile),
        rotate90(rotate90(tile)),
        rotate90(rotate90(rotate90(tile))),
        flipVertical(tile),
        rotate90(flipVertical(tile)),
        flipHorizontal(tile),
        rotate90(flipHorizontal(tile)),
    ]


def v_aligned(a, b) -> bool:
    for i, _ in enumerate(a):
        if a[i][-1] != b[i][0]:
            return False
    return True


def h_aligned(a, b) -> bool:
    for i, _ in enumerate(a):
        if a[-1][i] != b[0][i]:
            return False
    return True


def find_matches(grid):
    used_tiles = [tile["id"] for tile in grid]
    pos = len(grid)

    matches = []
    for tile in tiles:
        if tile["id"] in used_tiles:
            continue

        for perm in tile["permutations"]:
            if pos < grid_length:
                if v_aligned(grid[pos - 1]["pattern"], perm):
                    matches.append(grid + [{"id": tile["id"], "pattern": perm}])

            elif pos % grid_length == 0:
                if h_aligned(grid[pos - grid_length]["pattern"], perm):
                    matches.append(grid + [{"id": tile["id"], "pattern": perm}])

            elif v_aligned(grid[pos - 1]["pattern"], perm) and h_aligned(grid[pos - grid_length]["pattern"], perm):
                matches.append(grid + [{"id": tile["id"], "pattern": perm}])

    return matches


def compose_image(grid):
    tile_size = len(tiles[0]["pattern"]) - 2
    image = [[] for _ in range(grid_length * tile_size)]

    for y in range(grid_length):
        for x in range(grid_length):
            tile = grid[y * grid_length + x]
            for i in range(tile_size):
                image[y * tile_size + i] += tile["pattern"][i + 1][1:-1]

    return image


def cell_has_sea_monster(image, x, y):
    has_sea_monster = True
    for i, row in enumerate(SEA_MONSTER):
        for j, cell in enumerate(row):
            if cell == "#" and image[y + i][x + j] == ".":
                has_sea_monster = False

    return has_sea_monster


def add_sea_monster(image, x, y):
    for i, row in enumerate(SEA_MONSTER):
        for j, cell in enumerate(row):
            if cell == "#":
                image[y + i][x + j] = "O"


def find_sea_monsters(image):
    perms = permutations(image)
    y_range = len(image) - len(SEA_MONSTER)
    x_range = len(image[0]) - len(SEA_MONSTER[0])

    for perm in perms:
        for y in range(y_range):
            for x in range(x_range):
                if cell_has_sea_monster(perm, x, y):
                    matched_perm = perm
                    add_sea_monster(perm, x, y)

    return matched_perm


def calculate_roughness(image):
    return sum(sum(1 for x in row if x == "#") for row in image)


queue = []
for tile in tiles:
    tile["permutations"] = permutations(tile["pattern"])
    queue.append([{"id": tile["id"], "pattern": tile["pattern"]}])


while len(queue) > 0:
    grid, queue = queue[0], queue[1:]

    if len(grid) == len(tiles):
        completed_grid = grid
        break

    queue = find_matches(grid) + queue


corner_product = (
    int(completed_grid[0]["id"])
    * int(completed_grid[grid_length - 1]["id"])
    * int(completed_grid[-grid_length]["id"])
    * int(completed_grid[-1]["id"])
)

image = compose_image(completed_grid)
image_with_monsters = find_sea_monsters(image)
water_roughness = calculate_roughness(image_with_monsters)

time_elapsed = round(time() - start, 5)

print(
    f"""The product of the corner IDs is {corner_product}.
The water roughness is {water_roughness}.
Solution generated in {time_elapsed}s."""
)
