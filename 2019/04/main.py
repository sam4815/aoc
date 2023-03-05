from time import time

start = time()

with open("input.txt", encoding="utf8") as file:
    password_range = file.read().strip().split("-")


def repeating_pattern(password: str):
    repeating_counts = []

    curr, count = password[0], 1
    for digit in password[1:]:
        if digit == curr:
            count += 1
        else:
            repeating_counts.append(count)
            count = 1
        curr = digit

    repeating_counts.append(count)
    return repeating_counts


def higher_passwords(password: str):
    if len(password) == 1:
        return [str(x) for x in range(int(password[0]), 10)]

    higher = []
    for x in higher_passwords(password[0]):
        for y in higher_passwords(x * len(password[1:])):
            higher.append(x + y)

    return higher


num_first_criteria, num_second_criteria, max_num = 0, 0, int(password_range[1])

for password in higher_passwords(password_range[0][0] * 6):
    if int(password) > max_num:
        continue

    repeating = repeating_pattern(password)

    if any(x >= 2 for x in repeating):
        num_first_criteria += 1

    if 2 in repeating:
        num_second_criteria += 1


print(higher_passwords("888888"))

time_elapsed = round(time() - start, 5)

print(
    f"""{num_first_criteria} passwords meet the first set of criteria.
{num_second_criteria} passwords meet the second set of criteria.
Solution generated in {time_elapsed}s."""
)
