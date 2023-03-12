from time import time
from sys import path

path.append("../05")
from computer import Computer

start = time()

with open("input.txt", encoding="utf8") as file:
    memory = [int(x) for x in file.read().strip().split(",")]


class Game:
    def __init__(self, cpu):
        self.cpu = cpu
        self.grid = {}
        self.score = 0

    def set_tile(self, position, tile_id):
        self.grid[position] = tile_id

    def paddle_position(self):
        for coords, tile in self.grid.items():
            if tile == 3:
                return coords
        return None

    def ball_position(self):
        for coords, tile in self.grid.items():
            if tile == 4:
                return coords
        return None

    def move_paddle(self):
        paddle_pos, ball_pos = self.paddle_position(), self.ball_position()
        if paddle_pos is None or ball_pos is None:
            return

        x_paddle, x_ball = paddle_pos[0], ball_pos[0]
        self.cpu.inputs = [0]

        if x_paddle != x_ball:
            self.cpu.inputs = [1 if x_paddle < x_ball else -1]

    def run(self):
        while not self.cpu.terminated:
            position = (self.cpu.run(), self.cpu.run())
            if position == (-1, 0):
                self.score = self.cpu.run()
            else:
                self.set_tile(position, self.cpu.run())

            self.move_paddle()


game = Game(cpu=Computer(memory, []))
game.run()
num_block_tiles = sum(1 for tile in game.grid.values() if tile == 2)

cracked_game = Game(cpu=Computer([2] + memory[1:], []))
cracked_game.run()

time_elapsed = round(time() - start, 5)

print(
    f"""The game contains {num_block_tiles} block tiles.
After the last block is broken, the score is {cracked_game.score}.
Solution generated in {time_elapsed}s."""
)
