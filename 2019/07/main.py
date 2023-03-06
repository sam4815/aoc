from time import time
from sys import path

path.append("../05")
from computer import Computer

start = time()

with open("input.txt", encoding="utf8") as file:
    memory = [int(x) for x in file.read().strip().split(",")]


def output_signal(phases):
    computer = Computer(memory, [int(phases[0]), 0])
    for phase in phases[1:]:
        computer = Computer(memory, [int(phase), computer.run()])

    return computer.run()


def output_signal_loop(phases):
    a = Computer(memory, [int(phases[0]), 0])
    b = Computer(memory, [int(phases[1]), a.run()])
    c = Computer(memory, [int(phases[2]), b.run()])
    d = Computer(memory, [int(phases[3]), c.run()])
    e = Computer(memory, [int(phases[4]), d.run()])

    computers = [a, b, c, d, e]
    active_computer, active_computer_index = e, 4

    while True:
        active_computer_output = active_computer.run()
        if active_computer.terminated:
            break

        active_computer_index += 1
        active_computer = computers[active_computer_index % 5]
        active_computer.inputs.append(active_computer_output)

    return e.output


def phase_permutations(possible):
    if len(possible) == 1:
        return possible

    permutations = []
    for i in possible:
        permutations += [i + permutation for permutation in phase_permutations([x for x in possible if x != i])]

    return permutations


largest_signal = 0
for phase_permuation in phase_permutations(["0", "1", "2", "3", "4"]):
    signal = output_signal(phase_permuation)
    if signal > largest_signal:
        largest_signal = signal

largest_loop_signal = 0
for phase_permuation in phase_permutations(["5", "6", "7", "8", "9"]):
    signal = output_signal_loop(phase_permuation)
    if signal > largest_loop_signal:
        largest_loop_signal = signal

time_elapsed = round(time() - start, 5)

print(
    f"""The highest possible thruster signal is {largest_signal}.
The highest possible thruster signal using loops is {largest_loop_signal}.
Solution generated in {time_elapsed}s."""
)
