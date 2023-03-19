from time import time
from sys import path

path.append("../05")
from computer import Computer

start = time()

with open("input.txt", encoding="utf8") as file:
    memory = [int(x) for x in file.read().strip().split(",")]


def get_neighbours(grid, pos):
    neighbours = []
    for offset in [[-1, 0], [1, 0], [0, -1], [0, 1]]:
        offset_position = [pos[0] + offset[0], pos[1] + offset[1]]
        if offset_position[0] not in [-1, len(grid)] and offset_position[1] not in [-1, len(grid[0])]:
            neighbours.append(grid[offset_position[0]][offset_position[1]])

    return neighbours


ascii_program = Computer(memory, [])
scaffold_map, scaffold_row = [], []

while not ascii_program.terminated:
    output = ascii_program.run()
    if output == 10:
        scaffold_map.append(scaffold_row[:])
        scaffold_row = []
    else:
        scaffold_row.append(chr(output))

alignment_sum = 0
for i, row in enumerate(scaffold_map[:-3]):
    for j, cell in enumerate(row):
        if cell == "#" and get_neighbours(scaffold_map[:-3], [i, j]) == ["#"] * 4:
            scaffold_map[i][j] = "O"
            alignment_sum += i * j

MAIN_ROUTINE = ["B", "A", "A", "B", "C", "A", "C", "C", "A", "B"]
FUNCTION_A = ["R", "8", "L", "6", "L", "6"]
FUNCTION_B = ["R", "10", "R", "8", "L", "10", "L", "10"]
FUNCTION_C = ["L", "10", "R", "10", "L", "6"]
VIDEO_FEED = ["n"]

full_input = []
for line in [MAIN_ROUTINE, FUNCTION_A, FUNCTION_B, FUNCTION_C, VIDEO_FEED]:
    for instruction in line:
        for char in instruction:
            full_input.append(ord(char))
        full_input.append(ord(","))

    full_input[-1] = ord("\n")

vacuum_robot = Computer([2] + memory[1:], full_input)
while not vacuum_robot.terminated:
    dust_total = vacuum_robot.run()

time_elapsed = round(time() - start, 5)

print(
    f"""The sum of the alignment parameters is {alignment_sum}.
The vacuum robot reports collecting {dust_total} dust.
Solution generated in {time_elapsed}s."""
)
