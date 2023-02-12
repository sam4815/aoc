from time import time

start = time()

with open("input.txt", encoding="utf8") as file:
    starting_nums = [int(x) for x in file.read().strip().split(",")]

turn_map = {}
for i, num in enumerate(starting_nums):
    turn_map[num] = i + 1

turn, diff = len(starting_nums) + 1, len(starting_nums)
turn_2020, turn_30000000 = 0, 0

while True:
    spoken = 0 if diff == turn - 1 else diff
    diff = turn - turn_map.get(spoken, 0)
    turn_map[spoken] = turn

    if turn == 2020:
        turn_2020 = spoken

    if turn == 30000000:
        turn_30000000 = spoken
        break

    turn += 1

time_elapsed = round(time() - start, 5)

print(
    f"""The 2020th number spoken is {turn_2020}.
The 30000000th number spoken is {turn_30000000}.
Solution generated in {time_elapsed}s."""
)
