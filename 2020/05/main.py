import time

start = time.time()

highest_id, seat_dict = 0, {}

with open("input.txt", encoding="utf8") as file:
    boarding_passes = file.read().splitlines()


def binary_partition(space: int, search: str) -> int:
    while len(space) > 1:
        if search[0] in "FL":
            space = space[: int(len(space) / 2)]
        else:
            space = space[int(len(space) / 2) :]
        search = search[1:]

    return space[0]


for boarding_pass in boarding_passes:
    row = binary_partition(range(128), boarding_pass[:7])
    col = binary_partition(range(8), boarding_pass[7:])
    seat_id = row * 8 + col

    seat_dict[seat_id] = True

    if seat_id > highest_id:
        highest_id = seat_id

for row in range(11, 110):
    for col in range(8):
        seat_id = row * 8 + col
        if seat_dict.get(seat_id) is None:
            my_seat_id = seat_id

time_elapsed = round(time.time() - start, 5)

print(
    f"""The highest seat ID is {highest_id}.
The ID of my seat is {my_seat_id}.
Solution generated in {time_elapsed}s."""
)
