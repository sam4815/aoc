from time import time
from sys import path
from math import inf
from copy import deepcopy

path.append("../05")
from computer import Computer

start = time()

with open("input.txt", encoding="utf8") as file:
    memory = [int(x) for x in file.read().strip().split(",")]

MOVEMENTS = {
    1: (-1, 0),
    2: (1, 0),
    3: (0, -1),
    4: (0, 1),
}


class RepairDroid:
    def __init__(self, cpu):
        self.cpu = cpu
        self.steps = 0
        self.position = (0, 0)
        self.has_reached_oxygen = False

    def possible_moves(self, area):
        possible = []
        for direction, offset in MOVEMENTS.items():
            move = (self.position[0] + offset[0], self.position[1] + offset[1])
            if area.get(move) is None:
                possible.append(direction)
            elif area.get(move) < self.steps - 1:
                possible.append(direction)

        return possible

    def move(self, direction, grid):
        self.cpu.inputs = [direction]
        status = self.cpu.run()

        offset = MOVEMENTS[direction]
        next_position = (self.position[0] + offset[0], self.position[1] + offset[1])

        if status == 0:
            grid[next_position] = inf
        if status in (1, 2):
            self.steps += 1
            self.position = next_position
            grid[next_position] = self.steps
        if status == 2:
            self.has_reached_oxygen = True


shared_grid, oxygen_location = {}, (0, 0)
droid_queue = [RepairDroid(cpu=Computer(memory, []))]

while len(droid_queue) > 0:
    curr_droid, droid_queue = droid_queue[0], droid_queue[1:]

    if curr_droid.has_reached_oxygen:
        oxygen_location = curr_droid.position
        continue

    for direction in curr_droid.possible_moves(shared_grid):
        next_droid = deepcopy(curr_droid)
        next_droid.move(direction, shared_grid)
        droid_queue.append(next_droid)

min_steps = shared_grid[oxygen_location]

shared_grid[oxygen_location] = "O"
oxygen_minutes = 0


def step_oxygen(grid):
    oxygen_positions = [position for position, cell in grid.items() if cell == "O"]
    oxygen_moved = False

    for oxygen_position in oxygen_positions:
        for offset in MOVEMENTS.values():
            adjacent_position = (oxygen_position[0] + offset[0], oxygen_position[1] + offset[1])
            if isinstance(grid.get(adjacent_position), int):
                oxygen_moved = True
                grid[adjacent_position] = "O"

    return oxygen_moved


while step_oxygen(shared_grid):
    oxygen_minutes += 1

time_elapsed = round(time() - start, 5)

print(
    f"""The fewest number of steps required to get the droid to the oxygen system is {min_steps}.
It takes {oxygen_minutes} minutes for the area to fill with oxygen.
Solution generated in {time_elapsed}s."""
)
