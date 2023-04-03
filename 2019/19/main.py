from time import time
from sys import path

path.append("../05")
from computer import Computer

start = time()

with open("input.txt", encoding="utf8") as file:
    memory = [int(x) for x in file.read().strip().split(",")]

num_affected = 0
grid = [[] for _ in range(50)]


def within_beam(x, y):
    drone_system = Computer(memory, [x, y])
    while not drone_system.terminated:
        drone_system.run()

    return drone_system.output == 1


def within_beam_square(x, y, size):
    points = [(x + size - 1, y), (x, y - size + 1), (x + size - 1, y - size + 1)]
    return all(within_beam(point[0], point[1]) for point in points)


for y in range(50):
    for x in range(50):
        is_point_within_beam = within_beam(x, y)

        num_affected += 1 if is_point_within_beam else 0
        grid[y].append("#" if is_point_within_beam else ".")

curr_x, curr_y = 0, 10
while not within_beam_square(curr_x, curr_y, 100):
    curr_y += 1
    while not within_beam(curr_x, curr_y):
        curr_x += 1

square_code = curr_x * 10000 + (curr_y - 99)

time_elapsed = round(time() - start, 5)

print(
    f"""The number of points affected by the tractor beam is {num_affected}.
The coordinates of the first 100x100 square give code {square_code}.
Solution generated in {time_elapsed}s."""
)
