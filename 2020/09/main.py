from time import time
from typing import List

start = time()

with open("input.txt", encoding="utf8") as file:
    nums = [int(x) for x in file.read().splitlines()]


def has_summable_pair(target: int, numbers: List[int]) -> bool:
    for i, x in enumerate(numbers):
        for y in numbers[i + 1 :]:
            if x + y == target:
                return True

    return False


fails_rule = 0
for idx, num in enumerate(nums[25:]):
    if not has_summable_pair(num, nums[idx : idx + 25]):
        fails_rule = num
        break


def find_contiguous_region(target: int, numbers: List[int]) -> List[int]:
    for i, x in enumerate(numbers):
        total, contiguous = x, [x]
        while total < target:
            for y in numbers[i + 1 :]:
                total += y
                contiguous.append(y)

                if total == target:
                    return contiguous


contiguous_region = find_contiguous_region(fails_rule, nums)
contiguous_region.sort()
contiguous_sum = contiguous_region[0] + contiguous_region[len(contiguous_region) - 1]

time_elapsed = round(time() - start, 5)

print(
    f"""The first number that isn't a sum of the previous 25 is {fails_rule}.
The encryption weakness is {contiguous_sum}.
Solution generated in {time_elapsed}s."""
)
