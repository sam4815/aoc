import time
from typing import Tuple

start = time.time()

with open("input.txt", encoding="utf8") as file:
    tree_map = file.read().splitlines()


def count_trees(slope: Tuple[int, int]) -> int:
    num_trees, curr_pos = 0, (0, 0)

    while curr_pos[0] < len(tree_map):
        if tree_map[curr_pos[0]][curr_pos[1] % len(tree_map[0])] == "#":
            num_trees += 1

        curr_pos = (curr_pos[0] + slope[0], curr_pos[1] + slope[1])

    return num_trees


first_slope = count_trees((1, 3))

tree_product = count_trees((1, 1)) * first_slope
tree_product *= count_trees((1, 5))
tree_product *= count_trees((1, 7))
tree_product *= count_trees((2, 1))

time_elapsed = round(time.time() - start, 5)

print(
    f"""
Following the first slope, {first_slope} trees would be encountered.
The product of all the trees encountered on all of the slopes is {tree_product}.
Solution generated in {time_elapsed}s.
  """
)
