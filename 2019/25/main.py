from time import time
from re import findall
from sys import path

path.append("../05")
from computer import Computer

start = time()

with open("input.txt", encoding="utf8") as file:
    memory = [int(x) for x in file.read().strip().split(",")]

COMMANDS = [
    "north",
    "north",
    "east",
    "take antenna",
    "west",
    "south",
    "east",
    "take cake",
    "west",
    "south",
    "west",
    "west",
    "west",
    "take coin",
    "east",
    "east",
    "east",
    "east",
    "east",
    "east",
    "east",
    "take boulder",
    "north",
    "east",
]


def to_ascii(command):
    ascii_instruction = []
    for char in command:
        ascii_instruction.append(ord(char))
    ascii_instruction.append(ord("\n"))

    return ascii_instruction


def run_droid(mode="automatic"):
    droid = Computer(memory, [])
    ascii_output = ""

    if mode == "automatic":
        for command in COMMANDS:
            droid.inputs.extend(to_ascii(command))

    while not droid.terminated:
        ascii_output += chr(droid.run())

        if mode == "interactive" and ascii_output.endswith("Command?"):
            print(ascii_output)
            ascii_output = ""

            user_input = input("Enter command: ")
            droid.inputs = to_ascii(user_input)

    return findall(r"\d+", ascii_output)[-1]


password = run_droid(mode="automatic")

time_elapsed = round(time() - start, 5)

print(
    f"""The password for the main airlock is {password}.
Solution generated in {time_elapsed}s."""
)
