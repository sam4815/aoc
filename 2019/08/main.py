from time import time
from math import inf

start = time()

with open("input.txt", encoding="utf8") as file:
    image_data = [int(x) for x in [*file.read().strip()]]

image_width, image_height = 25, 6
layer_length = image_width * image_height

layers = []
for i, data in enumerate(image_data):
    if i % layer_length == 0:
        layers.append([])
    layers[i // layer_length].append(data)


layer_product, fewest_zeros = 0, inf
image = [[2] * image_width for _ in range(image_height)]

for layer in layers:
    num_zeros = sum(1 for x in layer if x == 0)

    if num_zeros < fewest_zeros:
        fewest_zeros = num_zeros
        layer_product = sum(1 for x in layer if x == 1) * sum(1 for x in layer if x == 2)

    for i, pixel in enumerate(layer):
        x, y = i % image_width, i // image_width
        if image[y][x] == 2:
            image[y][x] = pixel

output = ""
for row in image:
    for cell in row:
        output += "■" if cell == 0 else "□"
    output += "\n"

time_elapsed = round(time() - start, 5)

print(
    f"""The product of 1 and 2 digits for the layer with the fewest 0 digits is {layer_product}.
The final image is:
{output.strip()}
Solution generated in {time_elapsed}s."""
)
