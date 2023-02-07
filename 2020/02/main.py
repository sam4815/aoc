import time

start = time.time()

with open("input.txt", encoding="utf8") as file:
    lines = file.readlines()

policy_one_valid, policy_two_valid = 0, 0

for line in lines:
    range_str, letter, password = line.split(" ")
    min_range, max_range = [int(num) for num in range_str.split("-")]
    letter = letter[0]

    occurrences = password.count(letter)
    if min_range <= occurrences <= max_range:
        policy_one_valid += 1

    letters = password[min_range - 1 : min_range] + password[max_range - 1 : max_range]
    if (letters).count(letter) == 1:
        policy_two_valid += 1


time_elapsed = round(time.time() - start, 5)

print(
    f"""There are {policy_one_valid} valid passwords according to the first policy.
There are {policy_two_valid} valid passwords according to the second policy.
Solution generated in {time_elapsed}s."""
)
