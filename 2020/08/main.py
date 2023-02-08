from time import time
from typing import List, Dict, Tuple
from copy import deepcopy

start = time()

with open("input.txt", encoding="utf8") as file:
    instructions = [
        {"op": line.split(" ")[0], "value": int(line.split(" ")[1])}
        for line in file.read().splitlines()
    ]
    instructions.append({"op": "fin", "value": 0})


def run_program(instructs: List[Dict[str, str]]) -> Tuple[int, bool]:
    curr_index, acc, visited = 0, 0, {}
    while visited.get(curr_index) is None:
        visited[curr_index] = True
        ins = instructs[curr_index]

        if ins["op"] == "acc":
            acc += ins["value"]
            curr_index += 1
        elif ins["op"] == "jmp":
            curr_index += ins["value"]
        elif ins["op"] == "nop":
            curr_index += 1
        elif ins["op"] == "fin":
            return acc, True

    return acc, False


broken_accumulator = run_program(instructions)[0]
fixed_accumulator = 0

for index, instruction in enumerate(instructions):
    if instruction["op"] == "acc":
        continue

    new_instruction_set = deepcopy(instructions)
    new_instruction_set[index]["op"] = "nop" if instruction["op"] == "jmp" else "jmp"
    (accumulator, program_finished) = run_program(new_instruction_set)

    if program_finished:
        fixed_accumulator = accumulator
        break


time_elapsed = round(time() - start, 5)

print(
    f"""Before the infinite loop begins the accumulator is {broken_accumulator}.
After the program is fixed, the accumulator is {fixed_accumulator}.
Solution generated in {time_elapsed}s."""
)
