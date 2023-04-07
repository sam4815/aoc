# https://en.wikipedia.org/wiki/Modular_exponentiation
# https://en.wikipedia.org/wiki/Modular_multiplicative_inverse
from time import time

start = time()


def deal_with_increment(equation, inc):
    equation["x_coefficient"] *= inc
    equation["constant"] *= inc


def deal_into_new(equation):
    equation["x_coefficient"] = -equation["x_coefficient"]
    equation["constant"] = -equation["constant"] - 1


def cut(equation, offset):
    equation["constant"] -= offset


def decode_instruction(instruction):
    if "deal into" in instruction:
        return deal_into_new
    if "deal with increment" in instruction:
        increment = int(instruction.split(" ")[-1])
        return lambda equation: deal_with_increment(equation, increment)
    if "cut" in instruction:
        offset = int(instruction.split(" ")[-1])
        return lambda equation: cut(equation, offset)
    return None


def apply_equation(equation, x):
    return x * equation["x_coefficient"] + equation["constant"]


def nth_term(equation, mod, n):
    m = equation["x_coefficient"]
    c = equation["constant"]

    x_co = pow(m, n, mod)
    constant = c * (pow(m, n, mod) - 1) * pow(m - 1, -1, mod) % mod

    return {"x_coefficient": x_co, "constant": constant}


def solve_equation(equation, target, mod):
    inv = pow(equation["x_coefficient"], -1, mod)
    rhs = (target - equation["constant"]) % mod

    return (rhs * inv) % mod


with open("input.txt", encoding="utf8") as file:
    instructions = [decode_instruction(str) for str in file.read().strip().splitlines()]

equation = {"x_coefficient": 1, "constant": 0}
for instruction in instructions:
    instruction(equation)

index_2019 = apply_equation(equation, 2019) % 10007

iterated_equation = nth_term(equation, 119315717514047, 101741582076661)
index_2020 = solve_equation(iterated_equation, 2020, 119315717514047)

time_elapsed = round(time() - start, 5)

print(
    f"""After shuffling the deck once, the position of card 2019 is {index_2019}.
After shuffling the deck 101741582076661 times, the card in position 2020 is {index_2020}.
Solution generated in {time_elapsed}s."""
)
