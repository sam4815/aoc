from time import time
from re import match
from typing import List

start = time()


def apply_mask(num: int, mask: str) -> int:
    binary = [*format(num, "b").rjust(36, "0")]
    for i, char in enumerate(mask):
        if char != "X":
            binary[i] = char
    return int("".join(binary), 2)


def get_addresses(num: int, mask: str) -> List[int]:
    x_occurrences = mask.count("X")
    addresses = []

    for x in range(2**x_occurrences):
        x_values = format(x, "b").rjust(x_occurrences, "0")
        binary = [*format(num, "b").rjust(36, "0")]
        curr_x = 0

        for i, char in enumerate(mask):
            if char == "1":
                binary[i] = "1"
            elif char == "X":
                binary[i] = x_values[curr_x]
                curr_x += 1

        addresses.append(int("".join(binary), 2))

    return addresses


version_one_memory, version_two_memory = {}, {}
active_mask = ""

with open("input.txt", encoding="utf8") as file:
    for line in file.read().splitlines():
        if line.startswith("mask"):
            active_mask = line.split(" = ")[1]
        else:
            match_result = match(r"mem\[(\d+)] = (\d+)", line)
            mem_pos, value = int(match_result[1]), int(match_result[2])

            version_one_memory[mem_pos] = apply_mask(value, active_mask)

            for mem_addresses in get_addresses(mem_pos, active_mask):
                version_two_memory[mem_addresses] = value

version_one_memory_sum = sum(version_one_memory.values())
version_two_memory_sum = sum(version_two_memory.values())

time_elapsed = round(time() - start, 5)

print(
    f"""The sum of the remaining values in version one is {version_one_memory_sum}.
The sum of the remaining values in version two is {version_two_memory_sum}.
Solution generated in {time_elapsed}s."""
)
