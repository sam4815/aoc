from time import time

start = time()

with open("input.txt", encoding="utf8") as file:
    memory = [int(x) for x in file.read().strip().split(",")]


class Computer:
    def __init__(self, memory):
        self.memory = memory[:]
        self.pc = 0

    def inputs(self):
        return [self.memory[pos] for pos in [self.memory[self.pc + 1], self.memory[self.pc + 2]]]

    def output_pos(self):
        return self.memory[self.pc + 3]

    def set(self, position, value):
        self.memory[position] = value

    def add(self):
        self.set(self.output_pos(), sum(self.inputs()))
        self.pc += 4

    def mul(self):
        inputs = self.inputs()
        self.set(self.output_pos(), inputs[0] * inputs[1])
        self.pc += 4

    def run(self):
        while self.memory[self.pc] != 99:
            if self.memory[self.pc] == 1:
                self.add()

            elif self.memory[self.pc] == 2:
                self.mul()

        return self.memory[0]


gravity_assist_program = Computer(memory)
gravity_assist_program.set(1, 12)
gravity_assist_program.set(2, 2)
gravity_assist_result = gravity_assist_program.run()

for x in range(0, 100):
    for y in range(0, 100):
        program = Computer(memory)
        program.set(1, x)
        program.set(2, y)

        if program.run() == 19690720:
            noun_verb_combo = 100 * x + y


time_elapsed = round(time() - start, 5)

print(
    f"""The result of the gravity assist program is {gravity_assist_result}.
The noun-verb combo that produces the desired output is {noun_verb_combo}.
Solution generated in {time_elapsed}s."""
)
