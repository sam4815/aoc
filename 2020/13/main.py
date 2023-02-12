from time import time
from math import inf
from typing import List, Tuple

start = time()

with open("input.txt", encoding="utf8") as file:
    earliest_possible_time = int(file.readline().strip())
    buses = [(int(x), i) for i, x in enumerate(file.readline().split(",")) if x != "x"]

smallest_wait_time, bus_id = inf, 0
for bus in buses:
    bus_cadence = bus[0]
    wait_time = bus_cadence - (earliest_possible_time % bus_cadence)

    if wait_time < smallest_wait_time:
        smallest_wait_time, bus_id = wait_time, bus_cadence

id_wait_product = bus_id * smallest_wait_time


def valid_order(timestamp: int, buses: List[Tuple[int, int]]):
    for bus in buses:
        if (timestamp + bus[1]) % bus[0] != 0:
            return False
    return True


step = buses[0][0]
matched, earliest_timestamp = 0, 0
while True:
    if valid_order(earliest_timestamp, buses[: (matched + 2)]):
        matched += 1
        step *= buses[matched][0]

    if matched == len(buses) - 1:
        break

    earliest_timestamp += step

time_elapsed = round(time() - start, 5)

print(
    f"""The ID-wait product is {id_wait_product}.
The earliest possible departure for all buses is {earliest_timestamp}.
Solution generated in {time_elapsed}s."""
)
