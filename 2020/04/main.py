import time
import re

start = time.time()

with open("input.txt", encoding="utf8") as file:
    passports = file.read().strip().split("\n\n")

num_seemingly_valid_passports, num_valid_passports = 0, 0

validity_dict = {
    "byr": lambda x: 1920 <= int(x) <= 2002,
    "iyr": lambda x: 2010 <= int(x) <= 2020,
    "eyr": lambda x: 2010 <= int(x) <= 2030,
    "hcl": lambda x: bool(re.match(r"#[0-9a-f]{6}", x)),
    "ecl": lambda x: x in ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"],
    "pid": lambda x: len(x) == 9,
    "cid": lambda _: True,
    "hgt": lambda x: (150 <= int(x[:-2]) <= 193)
    if "cm" in x
    else ("in" in x and (59 <= int(x[:-2]) <= 76)),
}

for passport_str in passports:
    fields = {}
    for entry in re.split(r"\s", passport_str):
        key, value = entry.split(":")
        fields[key] = value

    if len(fields) == 8 or (len(fields) == 7 and fields.get("cid") is None):
        num_seemingly_valid_passports += 1

        if all(map(lambda x: validity_dict[x[0]](x[1]), fields.items())):
            num_valid_passports += 1

time_elapsed = round(time.time() - start, 5)

print(
    f"""
There are {num_seemingly_valid_passports} passports with the right number of fields.
There are {num_valid_passports} passports that are valid.
Solution generated in {time_elapsed}s."""
)
