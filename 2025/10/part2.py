import re
from ortools.sat.python import cp_model

with open("input.txt", encoding="utf8") as file:
   machines = [line.split(" ") for line in file.read().strip().splitlines()]
   machines = [[[int(x) for x in re.findall(r"\d+", chunk)] for chunk in line] for line in machines]

total = 0

for machine in machines:
    model = cp_model.CpModel()

    variables = {}
    for i in range(len(machine) - 2):
        variables[i] = model.NewIntVar(0, 10**10, 'x_{i}')

    model.Minimize(sum(variables.values()))

    for (i, constraint) in enumerate(machine[-1]):
        buttons = [j for (j, indices) in enumerate(machine[1:-1]) if i in indices]
        model.Add(sum(variables[j] for j in buttons) == constraint)

    solver = cp_model.CpSolver()
    solver.Solve(model)

    solution = [solver.value(x) for x in variables.values()]
    total += sum(solution)

print(total)

