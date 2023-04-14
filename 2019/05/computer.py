from typing import List


class Computer:
    def __init__(self, memory, inputs):
        self.memory = memory[:] + [0] * 4096
        self.inputs = inputs
        self.output, self.pc, self.base = 0, 0, 0
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

            if mode == "0":
                params.append(self.memory[param])
            if mode == "1":
                params.append(param)
            if mode == "2":
                params.append(self.memory[param + self.base])

        return params

    def memory_address(self, offset):
        address = self.memory[self.pc + offset]
        mode = self.current_instruction_str()[-2 - offset]

        if mode == "2":
            return address + self.base

        return address

    def set(self, position, value):
        self.memory[position] = value

    def add(self):
        a, b = self.params(2)
        self.set(self.memory_address(3), a + b)
        self.pc += 4

    def mul(self):
        a, b = self.params(2)
        self.set(self.memory_address(3), a * b)
        self.pc += 4

    def set_input(self):
        self.set(self.memory_address(1), self.inputs.pop(0))
        self.pc += 2

    def set_output(self):
        self.output = self.params(1)[0]
        self.pc += 2

    def set_base(self):
        self.base += self.params(1)[0]
        self.pc += 2

    def jump_if_true(self):
        a, address = self.params(2)
        if a != 0:
            self.pc = address
        else:
            self.pc += 3

    def jump_if_false(self):
        a, address = self.params(2)
        if a == 0:
            self.pc = address
        else:
            self.pc += 3

    def less_than(self):
        a, b = self.params(2)
        self.set(self.memory_address(3), 1 if a < b else 0)
        self.pc += 4

    def equals(self):
        a, b = self.params(2)
        self.set(self.memory_address(3), 1 if a == b else 0)
        self.pc += 4

    def run(self):
        while self.current_instruction() != 99 and not self.terminated:
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
            elif opcode == 9:
                self.set_base()

        self.terminated = True
        return self.output
