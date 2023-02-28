from time import time
from typing import List

start = time()

directions = {
    "e": [2, 0],
    "se": [1, -1],
    "sw": [-1, -1],
    "w": [-2, 0],
    "nw": [-1, 1],
    "ne": [1, 1],
}


def tokenize(steps: str) -> List[str]:
    tokens = []

    while len(steps) > 0:
        for direction in directions:
            if steps.startswith(direction):
                tokens.append(direction)
                steps = steps[len(direction) :]

    return tokens


def get_tile(position_map, pos: List[int]):
    return position_map.get(pos[0], {}).get(pos[1], "white")


def set_tile(position_map, pos: List[int], colour: str):
    if position_map.get(pos[0]) is None:
        position_map[pos[0]] = {}
    position_map[pos[0]][pos[1]] = colour


def neighbours(pos: List[int]) -> List[List[int]]:
    return [[pos[0] + diff[0], pos[1] + diff[1]] for diff in directions.values()]


def flip(tiles):
    flipped = {}

    for x, column in tiles.items():
        for y in column:
            for tile_pos in neighbours([x, y]):
                if flipped.get(tile_pos[0], {}).get(tile_pos[1]) is not None:
                    continue

                tile_colour = get_tile(tiles, tile_pos)
                neighbour_colours = [get_tile(tiles, neighbour_pos) for neighbour_pos in neighbours(tile_pos)]
                num_black_neighbours = sum(1 for colour in neighbour_colours if colour == "black")

                if tile_colour == "white" and num_black_neighbours == 2:
                    set_tile(flipped, tile_pos, "black")
                elif tile_colour == "black" and (num_black_neighbours == 0 or num_black_neighbours > 2):
                    set_tile(flipped, tile_pos, "white")
                else:
                    set_tile(flipped, tile_pos, tile_colour)

    return flipped


with open("input.txt", encoding="utf8") as file:
    tiles = [tokenize(steps) for steps in file.read().strip().splitlines()]


tile_map = {}

for tile in tiles:
    position = [0, 0]
    for step in tile:
        position[0] += directions[step][0]
        position[1] += directions[step][1]

    tile_colour = get_tile(tile_map, position)
    set_tile(tile_map, position, "black" if tile_colour == "white" else "white")

num_black_tiles = sum(sum(1 for colour in column.values() if colour == "black") for column in tile_map.values())

for i in range(100):
    tile_map = flip(tile_map)

num_black_tiles_100 = sum(sum(1 for colour in column.values() if colour == "black") for column in tile_map.values())

time_elapsed = round(time() - start, 5)

print(
    f"""After following the instructions, {num_black_tiles} black tiles are left.
After 100 days, {num_black_tiles_100} black tiles are left.
Solution generated in {time_elapsed}s."""
)
