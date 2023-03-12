from time import time
from re import findall

start = time()

with open("input.txt", encoding="utf8") as file:
    planet_positions = [[int(x) for x in findall(r"-?\d+", planet)] for planet in file.read().strip().splitlines()]
    planet_velocities = [[0] * 3 for _ in range(0, len(planet_positions))]


def gcd(a, b):
    if b == 0:
        return a
    return gcd(b, a % b)


def lcm(a, b):
    return (a * b) // gcd(a, b)


def apply_velocity(positions, velocities):
    for i, velocity in enumerate(velocities):
        for j, coordinate in enumerate(velocity):
            positions[i][j] += coordinate


def adjust_velocity(positions, velocities):
    for i, position_a in enumerate(positions):
        for j, position_b in enumerate(positions[i + 1 :]):
            for k, coordinate_a in enumerate(position_a):
                if coordinate_a < position_b[k]:
                    velocities[i][k] += 1
                    velocities[i + j + 1][k] -= 1
                if coordinate_a > position_b[k]:
                    velocities[i][k] -= 1
                    velocities[i + j + 1][k] += 1


def step(positions, velocities):
    adjust_velocity(positions, velocities)
    apply_velocity(positions, velocities)


def energy(position, velocity):
    return sum(abs(x) for x in position) * sum(abs(x) for x in velocity)


def sample(positions, velocities):
    coordinate_positions = [[], [], []]
    for _ in range(200000):
        coordinate_positions[0].append(positions[0][0])
        coordinate_positions[1].append(positions[0][1])
        coordinate_positions[2].append(positions[0][2])
        step(positions, velocities)

    return coordinate_positions


def find_cycle_length(sample):
    for possible_cycle in range(1, len(sample)):
        for i in range(100):
            if sample[i] != sample[i + possible_cycle]:
                break
            if i == 99:
                return possible_cycle
    return 0


for _ in range(1000):
    step(planet_positions, planet_velocities)

total_energy = sum(energy(pos, planet_velocities[i]) for i, pos in enumerate(planet_positions))

cycle_lengths = [find_cycle_length(coordinate) for coordinate in sample(planet_positions, planet_velocities)]
lowest_multiple = lcm(cycle_lengths[0], lcm(cycle_lengths[1], cycle_lengths[2]))

time_elapsed = round(time() - start, 5)

print(
    f"""The total energy in the system after 1000 steps is {total_energy}.
It takes {lowest_multiple} steps for the universe to repeat.
Solution generated in {time_elapsed}s."""
)
