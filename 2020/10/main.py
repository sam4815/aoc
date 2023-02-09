from time import time

start = time()

with open("input.txt", encoding="utf8") as file:
    nums = [int(x) for x in file.read().splitlines()]

nums.sort()
nums = [0] + nums + [(nums[-1] + 3)]

jolt_differences = {1: 0, 3: 0}
for idx, num in enumerate(nums[1:]):
    jolt_differences[num - nums[idx]] += 1

jolt_product = jolt_differences[1] * jolt_differences[3]

visited = {num: 1 if num == 0 else 0 for num in nums}
for idx, num in enumerate(nums[:-1]):
    for jump in range(1, 4):
        if jump + idx < len(nums) and nums[idx + jump] - num <= 3:
            visited[nums[idx + jump]] += visited[num]

num_distinct_paths = visited[nums[-1]]

time_elapsed = round(time() - start, 5)

print(
    f"""The product of 1-jolt and 3-jolt differences is {jolt_product}.
The number of distinct adapter arrangements is {num_distinct_paths}.
Solution generated in {time_elapsed}s."""
)
