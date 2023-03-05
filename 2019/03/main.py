from time import time
from typing import List
from math import inf

start = time()


def get_step(position_map, pos: List[int]):
    return position_map.get(pos[0], {}).get(pos[1], "")


def set_step(position_map, pos: List[int], step: str):
    if position_map.get(pos[0]) is None:
        position_map[pos[0]] = {}
    if position_map.get(pos[0]).get(pos[1]) is None:
        position_map[pos[0]][pos[1]] = step


def move(position: List[int], direction: str):
    if direction == "L":
        position[0] -= 1
    if direction == "R":
        position[0] += 1
    if direction == "U":
        position[1] += 1
    if direction == "D":
        position[1] -= 1


def manhattan_distance(position: List[int]):
    return abs(position[0]) + abs(position[1])


with open("input.txt", encoding="utf8") as file:
    wires = [
        [{"direction": instruction[0], "magnitude": int(instruction[1:])} for instruction in x.split(",")]
        for x in file.read().strip().splitlines()
    ]

positions, curr_pos, curr_distance = {}, [0, 0], 0
for instruction in wires[0]:
    for _ in range(instruction["magnitude"]):
        move(curr_pos, instruction["direction"])
        curr_distance += 1
        set_step(positions, curr_pos, "A" + str(curr_distance))

curr_pos, curr_distance, min_manhattan, min_steps = [0, 0], 0, inf, inf
for instruction in wires[1]:
    for _ in range(instruction["magnitude"]):
        move(curr_pos, instruction["direction"])
        curr_distance += 1

        if get_step(positions, curr_pos).startswith("A"):
            if manhattan_distance(curr_pos) < min_manhattan:
                min_manhattan = manhattan_distance(curr_pos)

            if curr_distance + int(get_step(positions, curr_pos)[1:]) < min_steps:
                min_steps = curr_distance + int(get_step(positions, curr_pos)[1:])

time_elapsed = round(time() - start, 5)

print(
    f"""The Manhattan distance to the nearest intersection is {min_manhattan}.
The fewest combined steps taken to reach an intersection is {min_steps}.
Solution generated in {time_elapsed}s."""
)
