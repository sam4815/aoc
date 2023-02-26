from time import time
from re import findall

start = time()

with open("input.txt", encoding="utf8") as file:
    all_foods = file.read().strip().splitlines()
    all_foods = [[findall(r"\w+", part) for part in food.split("contains")] for food in all_foods]

unknown_allergens, known_allergens = {}, {}
for food in all_foods:
    for allergen in food[1]:
        if unknown_allergens.get(allergen) is None:
            unknown_allergens[allergen] = set(food[0])
        else:
            unknown_allergens[allergen] &= set(food[0])

while len(known_allergens) != len(unknown_allergens):
    for allergen, common_ingredients in unknown_allergens.items():
        common_ingredients -= set(known_allergens.values())
        if len(common_ingredients) == 1:
            known_allergens[allergen] = list(common_ingredients)[0]

num_allergen_free = 0
for food in all_foods:
    for ingredient in food[0]:
        if ingredient not in known_allergens.values():
            num_allergen_free += 1

ingredient_list = ",".join([known_allergens[allergen] for allergen in sorted(known_allergens.keys())])

time_elapsed = round(time() - start, 5)

print(
    f"""Ingredients without allergens appear {num_allergen_free} times.
The canonical dangerous ingredient list is {ingredient_list}.
Solution generated in {time_elapsed}s."""
)
