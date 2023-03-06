from typing import List


class Computer:
    def __init__(self, memory, inputs):
        self.memory = memory[:]
        self.inputs = inputs
        self.output = 0
        self.pc = 0
        self.terminated = False

    def current_instruction(self):
        return self.memory[self.pc]

    def current_instruction_str(self):
        return str(self.current_instruction()).zfill(5)

    def params(self, num) -> List[int]:
        params = []
        for i in range(1, num + 1):
            param = self.memory[self.pc + i]
            mode = self.current_instruction_str()[-2 - i]
            params.append(self.memory[param] if mode == "0" else param)
        return params

    def set(self, position, value):
        self.memory[position] = value

    def add(self):
        a, b = self.params(2)
        self.set(self.memory[self.pc + 3], a + b)
        self.pc += 4

    def mul(self):
        a, b = self.params(2)
        self.set(self.memory[self.pc + 3], a * b)
        self.pc += 4

    def set_input(self):
        self.set(self.memory[self.pc + 1], self.inputs[0])
        self.inputs = self.inputs[1:]
        self.pc += 2

    def set_output(self):
        self.output = self.params(1)[0]
        self.pc += 2

    def jump_if_true(self):
        a, pointer = self.params(2)
        if a != 0:
            self.pc = pointer
        else:
            self.pc += 3

    def jump_if_false(self):
        a, pointer = self.params(2)
        if a == 0:
            self.pc = pointer
        else:
            self.pc += 3

    def less_than(self):
        a, b = self.params(2)
        self.set(self.memory[self.pc + 3], 1 if a < b else 0)
        self.pc += 4

    def equals(self):
        a, b = self.params(2)
        self.set(self.memory[self.pc + 3], 1 if a == b else 0)
        self.pc += 4

    def run(self):
        while self.current_instruction() != 99:
            opcode = self.current_instruction() % 100

            if opcode == 1:
                self.add()
            elif opcode == 2:
                self.mul()
            elif opcode == 3:
                self.set_input()
            elif opcode == 4:
                self.set_output()
                return self.output
            elif opcode == 5:
                self.jump_if_true()
            elif opcode == 6:
                self.jump_if_false()
            elif opcode == 7:
                self.less_than()
            elif opcode == 8:
                self.equals()

        self.terminated = True
        return self.output
