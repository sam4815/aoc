from time import time
from typing import List, Tuple
from re import match, findall

start = time()

with open("input.txt", encoding="utf8") as file:
    expressions = file.read().splitlines()


def operate(operand_a: int, operand_b: int, operator: str) -> int:
    if operator == "+":
        return operand_a + operand_b
    if operator == "*":
        return operand_a * operand_b


def evaluate_left_right(expression: List[str], index: int = 0) -> Tuple[int, int]:
    accumulator, active_operation = 0, "+"

    while index < len(expression):
        if expression[index] == "(":
            result, index = evaluate_left_right(expression, index + 1)
            accumulator = operate(accumulator, result, active_operation)

        if expression[index] == ")":
            return accumulator, min(index + 1, len(expression) - 1)

        if expression[index] in ["+", "*"]:
            active_operation = expression[index]

        if match(r"\d", expression[index]):
            curr_num = expression[index]
            while index + 1 < len(expression) and match(r"\d", expression[index + 1]):
                index += 1
                curr_num += expression[index]
            accumulator = operate(accumulator, int(curr_num), active_operation)

        index += 1

    return accumulator, index


def evaluate_addition(expression: str) -> str:
    additions = findall(r"\d+ \+ \d+", expression)

    while additions is not None and len(additions) > 0:
        for addition in additions[:1]:
            expression = expression.replace(addition, str(eval(addition)), 1)

        additions = findall(r"\d+ \+ \d+", expression)

    return expression


def evaluate_multiplication(expression: str) -> str:
    for multiplication in findall(r"\d+ \* \d+", expression):
        expression = expression.replace(multiplication, str(eval(multiplication)), 1)

    return expression


def evaluate_expression_without_subexpression(expression: str):
    while match(r"^\d+$", expression) is None:
        expression = evaluate_addition(expression)
        expression = evaluate_multiplication(expression)

    return expression


def evaluate_addition_first(expression: str):
    subexpressions = findall(r"\([^\(]+?\)", expression)

    while subexpressions is not None and len(subexpressions) > 0:
        for subexpression in subexpressions:
            evaluated = evaluate_expression_without_subexpression(subexpression[1:-1])
            expression = expression.replace(subexpression, evaluated)

        subexpressions = findall(r"\([^\(]+?\)", expression)

    return int(evaluate_expression_without_subexpression(expression))


left_right_expression_sum, addition_first_expression_sum = 0, 0
for expression in expressions:
    left_right_expression_sum += evaluate_left_right([*expression])[0]
    addition_first_expression_sum += evaluate_addition_first(expression)

time_elapsed = round(time() - start, 5)

print(
    f"""The sum of all expressions evaluated left to right is {left_right_expression_sum}.
The sum of all expressions evaluated addition-first is {addition_first_expression_sum}.
Solution generated in {time_elapsed}s."""
)
