from time import time
from re import match
from typing import List

start = time()

field_range_map, field_pos_map = {}, {}

with open("input.txt", encoding="utf8") as file:
    contents = file.read().strip().split("\n\n")
    for field in contents[0].splitlines():
        match_result = match(r"(.+?): (\d+-\d+).+?(\d+-\d+)", field)
        field_range_map[match_result[1]] = [
            [int(x) for x in match_result[2].split("-")],
            [int(x) for x in match_result[3].split("-")],
        ]

    my_ticket = [int(x) for x in contents[1].splitlines()[1].split(",")]
    nearby_tickets = [
        [int(x) for x in ticket.split(",")] for ticket in contents[2].splitlines()[1:]
    ]


def is_within_range(number: int, number_range: List[int]) -> bool:
    return number_range[0] <= number <= number_range[1]


def is_valid_number(number: int) -> bool:
    for field_ranges in field_range_map.values():
        for field_range in field_ranges:
            if is_within_range(number, field_range):
                return True
    return False


error_rate, valid_tickets = 0, []
for ticket in nearby_tickets:
    valid_ticket = True
    for num in ticket:
        if not is_valid_number(num):
            error_rate += num
            valid_ticket = False

    if valid_ticket:
        valid_tickets.append(ticket)

valid_ticket_field_values = {i: [] for i, _ in enumerate(valid_tickets[0])}
for ticket in valid_tickets:
    for i, num in enumerate(ticket):
        valid_ticket_field_values[i].append(num)


def values_satisfy_ranges(values: List[int], ranges: List[List[int]]):
    all_values_satisfy_range = True
    for value in values:
        if not any(is_within_range(value, field_range) for field_range in ranges):
            all_values_satisfy_range = False

    return all_values_satisfy_range


while len(field_pos_map) < len(field_range_map):
    for i, values in valid_ticket_field_values.items():
        num_valid_fields, last_possible_field = 0, ""
        for field_name, field_ranges in field_range_map.items():
            if field_pos_map.get(field_name) is not None:
                continue
            if values_satisfy_ranges(values, field_ranges):
                num_valid_fields += 1
                last_possible_field = field_name

        if num_valid_fields == 1:
            field_pos_map[last_possible_field] = i

departure_indices = [
    i for field, i in field_pos_map.items() if field.startswith("departure")
]

departure_fields_product = 1
for i in departure_indices:
    departure_fields_product *= my_ticket[i]

time_elapsed = round(time() - start, 5)

print(
    f"""The ticket scanner error rate is {error_rate}.
The product of fields beginning with "departure" is {departure_fields_product}.
Solution generated in {time_elapsed}s."""
)
