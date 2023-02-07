import time

start = time.time()

with open("input.txt", encoding="utf8") as file:
    nums = [int(line) for line in file.readlines()]

for x in range(len(nums) - 2):
    for y in range(x + 1, len(nums) - 1):
        if nums[x] + nums[y] == 2020:
            two_sum = nums[x] * nums[y]
        for z in range(y + 1, len(nums)):
            if nums[x] + nums[y] + nums[z] == 2020:
                three_sum = nums[x] * nums[y] * nums[z]

time_elapsed = round(time.time() - start, 5)

print(
    f"""The product of the two entries that sum to 2020 are {two_sum}.
The product of the three entries that sum to 2020 are {three_sum}.
Solution generated in {time_elapsed}s."""
)
