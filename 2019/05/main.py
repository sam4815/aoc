from time import time
from computer import Computer

start = time()

with open("input.txt", encoding="utf8") as file:
    memory = [int(x) for x in file.read().strip().split(",")]


diagnostic_program = Computer(memory, [1])
while not diagnostic_program.terminated:
    diagnostic_program.run()
diagnostic_result = diagnostic_program.output

thermal_radiator_program = Computer(memory, [5])
thermal_radiator_result = thermal_radiator_program.run()


time_elapsed = round(time() - start, 5)

print(
    f"""The result of the diagnostic program is {diagnostic_result}.
The result of the thermal radiator program is {thermal_radiator_result}.
Solution generated in {time_elapsed}s."""
)
