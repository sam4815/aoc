from time import time

start = time()

with open("input.txt", encoding="utf8") as file:
    masses = [int(x) for x in file.read().strip().splitlines()]


def fuel_requirement(x: int) -> int:
    fuel = additional_mass = x // 3 - 2
    while additional_mass > 0:
        additional_mass = additional_mass // 3 - 2
        if additional_mass > 0:
            fuel += additional_mass

    return fuel


naive_fuel_sum = sum(x // 3 - 2 for x in masses)
fuel_sum = sum(fuel_requirement(x) for x in masses)

time_elapsed = round(time() - start, 5)

print(
    f"""The total fuel required is {naive_fuel_sum}.
Taking the mass of the fuel into account, the total fuel required is {fuel_sum}.
Solution generated in {time_elapsed}s."""
)
