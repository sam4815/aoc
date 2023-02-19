from time import time
from re import search, findall

start = time()

with open("input.txt", encoding="utf8") as file:
    expressions = file.read().splitlines()


def evaluate_leftmost_operation(expression: str) -> str:
    leftmost = search(r"^\d+ [\+|\*] \d+", expression)

    while leftmost is not None:
        expression = expression.replace(leftmost.group(), str(eval(leftmost.group())), 1)
        leftmost = search(r"\d+ [\+|\*] \d+", expression)

    return expression


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


def evaluate_addition_first(expression: str) -> str:
    while search(r"^\d+$", expression) is None:
        expression = evaluate_addition(expression)
        expression = evaluate_multiplication(expression)

    return expression


def evaluate(expression: str, evaluation_function) -> str:
    subexpressions = findall(r"\([^\(]+?\)", expression)

    while subexpressions is not None and len(subexpressions) > 0:
        for subexpression in subexpressions:
            evaluated = evaluation_function(subexpression[1:-1])
            expression = expression.replace(subexpression, evaluated)

        subexpressions = findall(r"\([^\(]+?\)", expression)

    return int(evaluation_function(expression))


left_right_expression_sum, addition_first_expression_sum = 0, 0
for expression in expressions:
    left_right_expression_sum += evaluate(expression, evaluate_leftmost_operation)
    addition_first_expression_sum += evaluate(expression, evaluate_addition_first)

time_elapsed = round(time() - start, 5)

print(
    f"""The sum of all expressions evaluated left to right is {left_right_expression_sum}.
The sum of all expressions evaluated addition-first is {addition_first_expression_sum}.
Solution generated in {time_elapsed}s."""
)
