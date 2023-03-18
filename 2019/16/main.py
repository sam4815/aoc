from time import time
from math import floor

start = time()

with open("input.txt", encoding="utf8") as file:
    init_nums = [*[int(x) for x in file.read().strip()]]
    full_nums = init_nums * 10000
    offset = int("".join(str(x) for x in full_nums[:7]))


def multiply_by_pattern(nums, pos):
    pattern_sum = 0
    pattern_length = 4 * (pos + 1)

    for i, num in enumerate(nums[pos:]):
        quadrant = floor((((i + pos + 1) % pattern_length) / pattern_length) * 4)
        if quadrant == 1:
            pattern_sum += num
        if quadrant == 3:
            pattern_sum -= num

    return abs(pattern_sum) % 10


def phase(nums):
    next_nums = []
    for i, _ in enumerate(nums):
        next_nums.append(multiply_by_pattern(nums, i))

    return next_nums


def phase_by_offset(nums, off):
    next_nums = nums[:off]

    pattern_sum = abs(sum(nums[off:])) % 10
    next_nums.append(pattern_sum)

    for i in range(off, len(nums)):
        pattern_sum = (pattern_sum - nums[i]) % 10
        next_nums.append(pattern_sum)

    return next_nums


for _ in range(100):
    init_nums = phase(init_nums)
    full_nums = phase_by_offset(full_nums, offset)

first_8 = "".join(str(x) for x in init_nums[:8])
message = "".join(str(x) for x in full_nums[offset : offset + 8])

time_elapsed = round(time() - start, 5)

print(
    f"""After 100 phases of FFT, the first 8 digits are {first_8}.
The eight-digit message embedded in the final output list is {message}.
Solution generated in {time_elapsed}s."""
)
