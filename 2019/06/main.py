from time import time

start = time()

with open("input.txt", encoding="utf8") as file:
    orbits = [x.split(")") for x in file.read().strip().splitlines()]

planets = {"COM": {"name": "COM"}}

for orbited, orbiting in orbits:
    planets[orbiting] = {"name": orbiting, "orbiting": orbited}

for planet in planets.values():
    if planet.get("orbiting") is not None:
        planet["orbiting"] = planets[planet["orbiting"]]


def count_orbits(planet):
    count = 0
    while planet.get("orbiting") is not None:
        planet = planet.get("orbiting")
        count += 1
    return count


def planet_transfers(planet):
    costs, count = {}, 0
    while planet.get("orbiting") is not None:
        planet = planet.get("orbiting")
        costs[planet.get("name")] = count
        count += 1
    return costs


orbits_sum = sum(count_orbits(planet) for planet in planets.values())

source_transfers = planet_transfers(planets["YOU"])
target_planet = planets["SAN"]["orbiting"]
orbital_transfers = 0

while source_transfers.get(target_planet["name"]) is None:
    target_planet = target_planet["orbiting"]
    orbital_transfers += 1

orbital_transfers = orbital_transfers + source_transfers[target_planet["name"]]

time_elapsed = round(time() - start, 5)

print(
    f"""The total number of orbits is {orbits_sum}.
The minimum number of orbital transfers required is {orbital_transfers}.
Solution generated in {time_elapsed}s."""
)
