from time import time
from typing import List

start = time()

with open("input.txt", encoding="utf8") as file:
    rules, words = [block.splitlines() for block in file.read().split("\n\n")]

rule_map = dict([rule.split(": ") for rule in rules])


def match_rule_list(word: str, rule_list: str) -> List[int]:
    matched_indices = []
    queue = [{"index": 0, "rule": rule} for rule in rule_list]

    while len(queue) > 0:
        queue, curr = queue[1:], queue[0]
        rule, index = curr["rule"], curr["index"]

        if len(rule) == 0:
            matched_indices.append(index)
            continue

        sub_rule = rule[0]

        if isinstance(sub_rule, str):
            if word[index] == sub_rule:
                queue.append({"index": index + 1, "rule": rule[1:]})

        elif isinstance(sub_rule, list):
            for valid_index in match_rule_list(word[index:], sub_rule):
                queue.append({"index": index + valid_index, "rule": rule[1:]})

        elif sub_rule.get("pattern") is not None:
            for perm in sub_rule["pattern"]:
                if word[index:].startswith(perm):
                    queue.append({"index": index + len(perm), "rule": rule[1:]})

        elif sub_rule.get("recursive_pattern") is not None:
            for perm in sub_rule["recursive_pattern"]:
                if word[index:].startswith(perm):
                    queue.append({"index": index + len(perm), "rule": rule[1:]})
                    queue.append({"index": index + len(perm), "rule": rule})

        else:
            for perm in sub_rule["start_pattern"]:
                if word[index:].startswith(perm):
                    simple = [{"pattern": sub_rule["end_pattern"]}] + rule[1:]
                    rec = [sub_rule] + [{"pattern": sub_rule["end_pattern"]}] + rule[1:]

                    queue.append({"index": index + len(perm), "rule": simple})
                    queue.append({"index": index + len(perm), "rule": rec})

    return matched_indices


def permutations(rule_list: str) -> List[str]:
    permutations = []
    for i in range(2**8):
        binary = [*format(i, "b").rjust(8, "0")]
        word = "".join(["a" if x == "0" else "b" for x in binary])

        if len(match_rule_list(word, rule_list)) > 0:
            permutations.append(word)

    return permutations


def define_rule(rule_list: str, allow_loops: bool = False):
    if rule_list == "42" and allow_loops:
        return {"recursive_pattern": permutations(define_rule(rule_map["42"]))}

    if rule_list == "42 31" and allow_loops:
        return {
            "start_pattern": permutations(define_rule(rule_map["42"])),
            "end_pattern": permutations(define_rule(rule_map["31"])),
        }

    return [
        [
            define_rule(rule_map[sub], allow_loops) if sub.isdigit() else sub.strip('"')
            for sub in rule.split(" ")
        ]
        for rule in rule_list.split(" | ")
    ]


num_match_zero, num_match_loopable_zero = 0, 0

rule_zero = define_rule(rule_map["0"], allow_loops=False)
looping_rule_zero = define_rule(rule_map["0"], allow_loops=True)

for word in words:
    if len(word) in match_rule_list(word, rule_zero):
        num_match_zero += 1
    if len(word) in match_rule_list(word, looping_rule_zero):
        num_match_loopable_zero += 1

time_elapsed = round(time() - start, 5)

print(
    f"""{num_match_zero} messages completely match rule 0.
With loops, {num_match_loopable_zero} messages completely match rule 0.
Solution generated in {time_elapsed}s."""
)
