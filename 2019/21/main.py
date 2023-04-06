from time import time
from sys import path

path.append("../05")
from computer import Computer

start = time()

with open("input.txt", encoding="utf8") as file:
    memory = [int(x) for x in file.read().strip().split(",")]

WALK_INSTRUCTIONS = [
    "NOT A J",
    "NOT B T",
    "OR T J",
    "NOT C T",
    "OR T J",
    "AND D J",
    "WALK",
]

RUN_INSTRUCTIONS = [
    "NOT A J",
    "NOT B T",
    "OR T J",
    "NOT C T",
    "OR T J",
    "AND D J",
    "NOT E T",
    "NOT T T",
    "OR H T",
    "AND T J",
    "RUN",
]


def to_ascii(instructions):
    ascii_instructions = []
    for instruction in instructions:
        for char in instruction:
            ascii_instructions.append(ord(char))
        ascii_instructions.append(ord("\n"))

    return ascii_instructions


def run_springdroid(instructions):
    springdroid = Computer(memory, to_ascii(instructions))
    hull_damage = 0
    ascii_output = ""

    while not springdroid.terminated:
        output = springdroid.run()
        if 0 <= output <= 127:
            ascii_output += chr(output)
        else:
            hull_damage = output

    return ascii_output, hull_damage


walk_output, walk_damage = run_springdroid(WALK_INSTRUCTIONS)
run_output, run_damage = run_springdroid(RUN_INSTRUCTIONS)

time_elapsed = round(time() - start, 5)

print(
    f"""The amount of hull damage reported when the droid is walking is {walk_damage}.
The amount of hull damage reported when the droid is running is {run_damage}.
Solution generated in {time_elapsed}s."""
)
