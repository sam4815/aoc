from time import time
from math import ceil

start = time()

with open("input.txt", encoding="utf8") as file:
    reactions = [reaction.split(" => ") for reaction in file.read().strip().splitlines()]

chemical_map = {"ORE": {"quantity": 1, "inputs": {"ORE": 1}}}
for inputs, output in reactions:
    output_quantity, output_chemical = output.split(" ")
    chemical_map[output_chemical] = {"quantity": int(output_quantity), "inputs": {}}

    for inp in inputs.split(", "):
        input_quantity, input_chemical = inp.split(" ")
        chemical_map[output_chemical]["inputs"][input_chemical] = int(input_quantity)


def calculate_required_ore(reaction_map, leftovers):
    required = dict(reaction_map["FUEL"]["inputs"])
    while len(required) != 1:
        next_required = {}

        for input_chemical, input_quantity in required.items():
            if leftovers[input_chemical] > 0:
                leftover_used = min(input_quantity, leftovers[input_chemical])
                input_quantity -= leftover_used
                leftovers[input_chemical] -= leftover_used

            chemical = reaction_map[input_chemical]
            required_multiple = ceil(input_quantity / chemical["quantity"])

            for required_chemical, required_quantity in chemical["inputs"].items():
                required = next_required.get(required_chemical, 0) + required_quantity * required_multiple
                next_required[required_chemical] = required

            leftovers[input_chemical] += (chemical["quantity"] * required_multiple) - input_quantity

        required = next_required

    return required["ORE"]


leftover = {chem: 0 for chem in chemical_map}
min_ore = calculate_required_ore(chemical_map, leftover)

remaining_ore = 1000000000000 - min_ore

fuel_produced = 0
while remaining_ore > 0:
    fuel_produced += 1
    remaining_ore -= calculate_required_ore(chemical_map, leftover)
    if fuel_produced % 10000 == 0:
        print(remaining_ore)

time_elapsed = round(time() - start, 5)

print(
    f"""The minimum amount of ORE required to produe 1 FUEL is {min_ore}.
The maximum amount of FUEL that can be produced is {fuel_produced}.
Solution generated in {time_elapsed}s."""
)
