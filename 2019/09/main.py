from time import time
from sys import path

path.append("../05")
from computer import Computer

start = time()

with open("input.txt", encoding="utf8") as file:
    memory = [int(x) for x in file.read().strip().split(",")]

boost_program_test_mode = Computer(memory, [1])
boost_keycode = boost_program_test_mode.run()

boost_program_sensor_mode = Computer(memory, [2])
distress_cooordinates = boost_program_sensor_mode.run()

time_elapsed = round(time() - start, 5)

print(
    f"""The BOOST keycode is {boost_keycode}.
The coordinates of the distress signal are {distress_cooordinates}.
Solution generated in {time_elapsed}s."""
)
