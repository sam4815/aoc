from time import time

start = time()

with open("input.txt", encoding="utf8") as file:
    keys = [int(x) for x in file.read().strip().splitlines()]


def step(val: int, subject: int) -> int:
    return (val * subject) % 20201227


def find_loop_size(public_key: int) -> int:
    val, loop = 1, 0
    while val != public_key:
        val = step(val, 7)
        loop += 1
    return loop


def decode(public_key: int, loop_size: int) -> int:
    val = 1
    for _ in range(loop_size):
        val = step(val, public_key)
    return val


card_public_key, door_public_key = keys[0], keys[1]
encryption_key = decode(card_public_key, find_loop_size(door_public_key))

time_elapsed = round(time() - start, 5)

print(
    f"""The encryption key is {encryption_key}.
Solution generated in {time_elapsed}s."""
)
