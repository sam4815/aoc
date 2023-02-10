from time import time
from dataclasses import dataclass
from enum import Enum


class Direction(Enum):
    EAST = 1
    SOUTH = 2
    WEST = 3
    NORTH = 4


ROTATE_LEFT_MAP = {
    Direction.EAST: Direction.NORTH,
    Direction.SOUTH: Direction.EAST,
    Direction.WEST: Direction.SOUTH,
    Direction.NORTH: Direction.WEST,
}

ROTATE_RIGHT_MAP = {
    Direction.NORTH: Direction.EAST,
    Direction.EAST: Direction.SOUTH,
    Direction.SOUTH: Direction.WEST,
    Direction.WEST: Direction.NORTH,
}

DIRECTION_MAP = {
    "E": Direction.EAST,
    "S": Direction.SOUTH,
    "W": Direction.WEST,
    "N": Direction.NORTH,
}


@dataclass
class Seabound:
    facing: Direction
    east_mag: int
    north_mag: int


def rotate_ship_left(ship: Seabound):
    ship["facing"] = ROTATE_LEFT_MAP[ship["facing"]]


def rotate_ship_right(ship: Seabound):
    ship["facing"] = ROTATE_RIGHT_MAP[ship["facing"]]


def rotate_waypoint_left(wpoint: Seabound):
    wpoint["east_mag"], wpoint["north_mag"] = -wpoint["north_mag"], wpoint["east_mag"]


def rotate_waypoint_right(wpoint: Seabound):
    wpoint["east_mag"], wpoint["north_mag"] = wpoint["north_mag"], -wpoint["east_mag"]


def move(sea_object: Seabound, direction: Direction, value: int):
    if direction == Direction.NORTH:
        sea_object["north_mag"] += value
    if direction == Direction.SOUTH:
        sea_object["north_mag"] -= value
    if direction == Direction.EAST:
        sea_object["east_mag"] += value
    if direction == Direction.WEST:
        sea_object["east_mag"] -= value


start = time()

first_ship = {"facing": Direction.EAST, "east_mag": 0, "north_mag": 0}
second_ship = {"facing": Direction.EAST, "east_mag": 0, "north_mag": 0}
waypoint = {"facing": Direction.EAST, "east_mag": 10, "north_mag": 1}

with open("input.txt", encoding="utf8") as file:
    instructions = file.read().splitlines()

for instruction in instructions:
    action, val = instruction[0], int(instruction[1:])

    if action in ["N", "S", "E", "W"]:
        move(first_ship, DIRECTION_MAP[action], val)
        move(waypoint, DIRECTION_MAP[action], val)

    if action == "L":
        for _ in range(int(val / 90)):
            rotate_ship_left(first_ship)
            rotate_waypoint_left(waypoint)
    if action == "R":
        for _ in range(int(val / 90)):
            rotate_ship_right(first_ship)
            rotate_waypoint_right(waypoint)

    if action == "F":
        move(first_ship, first_ship["facing"], val)
        second_ship["north_mag"] += waypoint["north_mag"] * val
        second_ship["east_mag"] += waypoint["east_mag"] * val


first_manhattan_distance = abs(first_ship["east_mag"]) + abs(first_ship["north_mag"])
second_manhattan_distance = abs(second_ship["east_mag"]) + abs(second_ship["north_mag"])

time_elapsed = round(time() - start, 5)

print(
    f"""The first Manhattan distance from the start is {first_manhattan_distance}.
The second Manhattan distance from the start is {second_manhattan_distance}.
Solution generated in {time_elapsed}s."""
)
