from time import time

start = time()

with open("input.txt", encoding="utf8") as file:
    asteroid_map = [[*row] for row in file.read().strip().splitlines()]

asteroids = {}
for y, row in enumerate(asteroid_map):
    for x, cell in enumerate(row):
        if cell == "#":
            asteroids[(x, y)] = True


def gcd(a, b):
    if b == 0:
        return a
    return gcd(b, a % b)


def is_visible(location, asteroid, all_asteroids):
    if location == asteroid:
        return False

    x_offset, y_offset = asteroid[0] - location[0], asteroid[1] - location[1]
    factor = abs(gcd(x_offset, y_offset))
    x_step, y_step = x_offset // factor, y_offset // factor

    for i in range(1, abs(factor)):
        possible_coordinates = (location[0] + x_step * i, location[1] + y_step * i)
        if all_asteroids.get(possible_coordinates) is not None:
            return False

    return True


max_num_asteroids, station_location = 0, (0, 0)
for location in asteroids:
    num_visible = sum(1 for asteroid in asteroids if is_visible(location, asteroid, asteroids))
    if num_visible > max_num_asteroids:
        max_num_asteroids = num_visible
        station_location = location


def clockwise_priority(location):
    x, y = station_location[0] - location[0], location[1] - station_location[1]
    if x <= 0 and y < 0:
        return abs(x) / abs(y)
    if x < 0 and y >= 0:
        return abs(y) / abs(x) + 1000
    if x >= 0 and y > 0:
        return abs(x) / abs(y) + 2000
    return abs(y) / abs(x) + 3000


currently_visible = [ast for ast in asteroids if is_visible(station_location, ast, asteroids)]
currently_visible.sort(key=clockwise_priority)

asteroid_200 = currently_visible[199]
asteroid_sum = asteroid_200[0] * 100 + asteroid_200[1]

time_elapsed = round(time() - start, 5)

print(
    f"""{max_num_asteroids} asteroids can be detected from the best position for a monitoring station.
The sum of the coordinates of the 200th asteroid to be vaporized is {asteroid_sum}.
Solution generated in {time_elapsed}s."""
)
