from time import time
from typing import List, Tuple, Callable

start = time()

with open("input.txt", encoding="utf8") as file:
    old_seats = [list(row) for row in file.read().splitlines()]
    new_seats = [list(row) for row in old_seats]


directions = [
    [-1, -1],
    [-1, 0],
    [-1, +1],
    [0, -1],
    [0, +1],
    [+1, -1],
    [+1, 0],
    [+1, +1],
]


def count_occupied_seats(seats: List[List[str]]) -> int:
    return sum(len(["#" for seat in row if seat == "#"]) for row in seats)


def valid_position(x: int, y: int) -> bool:
    return 0 <= x < len(old_seats) and 0 <= y < len(old_seats[0])


def count_neighbours(seats: List[List[str]], position: Tuple[int, int]) -> int:
    num_occupied = 0
    for direction in directions:
        curr_x, curr_y = position[0] + direction[0], position[1] + direction[1]
        if valid_position(curr_x, curr_y) and seats[curr_x][curr_y] == "#":
            num_occupied += 1

    return num_occupied


def count_visible(seats: List[List[str]], position: Tuple[int, int]) -> int:
    num_occupied = 0
    for direction in directions:
        curr_x, curr_y = position[0] + direction[0], position[1] + direction[1]
        while valid_position(curr_x, curr_y) and seats[curr_x][curr_y] == ".":
            curr_x, curr_y = curr_x + direction[0], curr_y + direction[1]

        if valid_position(curr_x, curr_y) and seats[curr_x][curr_y] == "#":
            num_occupied += 1

    return num_occupied


def step(condition: Callable, occupied_num: int) -> bool:
    moved = False

    for x, old_row in enumerate(old_seats):
        for y, old_seat in enumerate(old_row):
            if old_seat == "L" and condition(old_seats, [x, y]) == 0:
                new_seats[x][y] = "#"
                moved = True
            elif old_seat == "#" and condition(old_seats, [x, y]) >= occupied_num:
                new_seats[x][y] = "L"
                moved = True
            else:
                new_seats[x][y] = old_seats[x][y]

    return moved


while step(count_neighbours, 4):
    old_seats, new_seats = new_seats, old_seats
first_pass_occupied = count_occupied_seats(new_seats)

old_seats = [["." if seat == "." else "L" for seat in row] for row in old_seats]

while step(count_visible, 5):
    old_seats, new_seats = new_seats, old_seats
second_pass_occupied = count_occupied_seats(new_seats)

time_elapsed = round(time() - start, 5)

print(
    f"""There are {first_pass_occupied} occupied seats after the first simulation.
There are {second_pass_occupied} occupied seats after the second simulation.
Solution generated in {time_elapsed}s."""
)
