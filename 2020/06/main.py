import time
import functools

start = time.time()

with open("input.txt", encoding="utf8") as file:
    groups = file.read().strip().split("\n\n")

yes_count, all_yes_count = 0, 0

for group in groups:
    num_people = group.count("\n") + 1
    group = group.replace("\n", "")

    yes_questions = functools.reduce(
        lambda acc, x: acc.update({x: acc.get(x, 0) + 1}) or acc, group, {}
    )
    yes_count += len(yes_questions)
    all_yes_count += len(
        list(filter(lambda x: x == num_people, yes_questions.values()) or [])
    )


time_elapsed = round(time.time() - start, 5)

print(
    f"""The number of questions that were answered with "yes" is {yes_count}.
The number of questions that everyone answered with "yes" is {all_yes_count}.
Solution generated in {time_elapsed}s."""
)
