from time import time
from re import search

start = time()

with open("input.txt", encoding="utf8") as file:
    rules, words = [block.splitlines() for block in file.read().split("\n\n")]

rule_map = dict([rule.split(": ") for rule in rules])

while search(r"\d+", rule_map["0"]) is not None:
    for rule_num, rule in rule_map.items():
        rule_map["0"] = rule_map["0"].replace(rule_num, f"({rule})")
    print(rule_map["0"])


def match_rule(word: str, rule: str) -> bool:
    index = 0
    return True


num_match_zero = 0
for word in words:
    if match_rule(word, rule_map["0"]):
        num_match_zero += 1


time_elapsed = round(time() - start, 5)

print(
    f"""{num_match_zero} messages completely match rule 0.
Solution generated in {time_elapsed}s."""
)
