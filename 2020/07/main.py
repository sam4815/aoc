import time
import re
from typing import Dict
from dataclasses import dataclass


@dataclass
class Bag:
    contents: Dict[str, int]
    type: str


start = time.time()

with open("input.txt", encoding="utf8") as file:
    bag_descriptions = file.readlines()

BAGS = {}
for bag_description in bag_descriptions:
    type_of_bag = re.match(r"(\w+ \w+)", bag_description)[0]
    contents = re.findall(r"(\d+)\s(\w+ \w+)\sbag", bag_description)

    contents_map = {}
    for content in contents:
        contents_map[content[1]] = int(content[0])

    BAGS[type_of_bag] = Bag(contents_map, type_of_bag)


def has_gold(bag: Bag) -> bool:
    if bag.type == "shiny gold":
        return True

    return any(has_gold(BAGS[bag_type]) for bag_type in bag.contents.keys())


def count_bags(bag: Bag) -> int:
    return sum(
        num + num * count_bags(BAGS[bag_type]) for bag_type, num in bag.contents.items()
    )


# Subtract one to discard the shiny gold bag itself
num_bags_containing_gold = len([bag for bag in BAGS.values() if has_gold(bag)]) - 1
num_bags_inside_gold = count_bags(BAGS["shiny gold"])

time_elapsed = round(time.time() - start, 5)

print(
    f"""The number of bags containing a gold bag is {num_bags_containing_gold}.
The number of bags required inside a gold bag is {num_bags_inside_gold}.
Solution generated in {time_elapsed}s."""
)
