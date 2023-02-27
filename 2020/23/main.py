from time import time

start = time()

with open("input.txt", encoding="utf8") as file:
    small_cups = [int(x) for x in [*file.read().strip()]]
    large_cups = small_cups + list(range(10, 1000001))


def link(cups):
    init = {"val": cups[0]}
    curr_cup = init
    cup_map = {cups[0]: curr_cup}

    for cup in cups[1:]:
        curr_cup["next"] = {"val": cup}
        curr_cup = curr_cup["next"]
        cup_map[cup] = curr_cup

    curr_cup["next"] = init
    curr_cup = curr_cup["next"]

    return curr_cup, cup_map


def find(cups, val):
    found_cup = cups

    while found_cup["val"] != val and found_cup["next"] is not None:
        found_cup = found_cup["next"]

    if found_cup["val"] != val:
        return None

    return found_cup


def insert(cups, elements):
    next_cup = cups["next"]
    cups["next"] = elements
    count = 0
    while cups["next"] is not None:
        count += 1
        cups = cups["next"]
    cups["next"] = next_cup


def pop(cups, num):
    ref, removed = cups, cups["next"]
    for _ in range(num):
        ref = ref["next"]
    cups["next"] = ref["next"]
    ref["next"] = None
    return removed


def get_label(cups):
    label = ""
    while cups["val"] != 1:
        label += str(cups["val"])
        cups = cups["next"]
    return label


def length(cups):
    initial, length = cups["val"], 1
    cups = cups["next"]
    while cups["val"] != initial:
        length += 1
        cups = cups["next"]
    return length


def max_val(cups):
    initial, highest = cups["val"], cups["val"]
    cups = cups["next"]
    while cups["val"] != initial:
        if cups["val"] > highest:
            highest = cups["val"]
        cups = cups["next"]
    return highest


def step(cups, cup_map):
    removed = pop(cups, 3)
    destination_val = cups["val"] - 1

    while destination_val <= 0 or find(removed, destination_val) is not None:
        destination_val -= 1
        if destination_val <= 0:
            destination_val = max_val(cups)

    insert_cup = cup_map[destination_val]

    insert(insert_cup, removed)
    return cups["next"]


small_cups, small_cup_map = link(small_cups)
for _ in range(100):
    small_cups = step(small_cups, small_cup_map)

large_cups, large_cup_map = link(large_cups)
for i in range(10000000):
    large_cups = step(large_cups, large_cup_map)

cup_labels = get_label(small_cup_map[1]["next"])
star_cup_product = large_cup_map[1]["next"]["val"] * large_cup_map[1]["next"]["next"]["val"]

time_elapsed = round(time() - start, 5)

print(
    f"""After 100 crab moves, the cup labels are {cup_labels}.
The product of the two cups hiding stars is {star_cup_product}.
Solution generated in {time_elapsed}s."""
)
