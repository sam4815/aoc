from time import time
from sys import path

path.append("../05")
from computer import Computer

start = time()

with open("input.txt", encoding="utf8") as file:
    memory = [int(x) for x in file.read().strip().split(",")]


class PaintingRobot:
    def __init__(self, cpu):
        self.cpu = cpu
        self.facing = "up"
        self.curr_position = (0, 0)
        self.hull = {}

    def get_colour(self):
        return self.hull.get(self.curr_position, "black")

    def get_input(self):
        return 1 if self.get_colour() == "white" else 0

    def set_colour(self, colour_code):
        self.hull[self.curr_position] = "white" if colour_code == 1 else "black"

    def set_direction(self, direction_code):
        if direction_code == 0:
            self.turn_left()
        if direction_code == 1:
            for _ in range(3):
                self.turn_left()

    def turn_left(self):
        self.facing = {
            "up": "left",
            "left": "down",
            "down": "right",
            "right": "up",
        }[self.facing]

    def move_forward(self):
        x, y = self.curr_position

        if self.facing == "up":
            self.curr_position = (x, y + 1)
        if self.facing == "down":
            self.curr_position = (x, y - 1)
        if self.facing == "right":
            self.curr_position = (x + 1, y)
        if self.facing == "left":
            self.curr_position = (x - 1, y)

    def run(self):
        while not self.cpu.terminated:
            self.cpu.inputs = [self.get_input()]
            self.set_colour(self.cpu.run())
            self.set_direction(self.cpu.run())
            self.move_forward()


black_robot = PaintingRobot(cpu=Computer(memory, []))
black_robot.run()
num_painted_panels = len(black_robot.hull)

white_robot = PaintingRobot(cpu=Computer(memory, []))
white_robot.hull[(0, 0)] = "white"
white_robot.run()

hull = ""
x_max = max(x for x, _ in white_robot.hull)
y_min = min(y for _, y in white_robot.hull)

for y in range(0, y_min - 1, -1):
    for x in range(0, x_max):
        hull += "■" if white_robot.hull.get((x, y), "black") == "black" else "□"
    hull += "\n"

time_elapsed = round(time() - start, 5)

print(
    f"""The robot paints {num_painted_panels} panels at least once.
After starting on a white panel, the robot paints:
{hull.strip()}
Solution generated in {time_elapsed}s."""
)
