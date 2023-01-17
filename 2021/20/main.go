package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type Image [][]string
type Point struct {
	x, y int
}
type Binary string

func (image Image) PaddedEmpty() Image {
	padded := make(Image, len(image)+2)

	for i := 0; i < len(padded); i++ {
		padded[i] = make([]string, len(image[0])+2)
		for j := 0; j < len(padded[i]); j++ {
			padded[i][j] = "."
		}
	}

	return padded
}

func (image Image) Print() {
	for _, line := range image {
		log.Print(line)
	}
}

func (image Image) Get(point Point) string {
	return image[point.x][point.y]
}

func (image *Image) Set(point Point, val string) {
	(*image)[point.x][point.y] = val
}

func (image Image) ForEach(f func(Point)) {
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			f(Point{i, j})
		}
	}
}

func FindInfinitePixel(algo string, step int) string {
	first_pixel, last_pixel := string(algo[0]), string(algo[len(algo)-1])

	if first_pixel == "." {
		return "."
	}
	if first_pixel == "#" && last_pixel == "#" && step > 0 {
		return "#"
	}
	if first_pixel == "#" && last_pixel == "." && step%2 == 1 {
		return "#"
	}
	return "."
}

func (image Image) Binary(point Point, algo string, step int) Binary {
	neighbours_string := Binary("")
	for i := point.x - 1; i <= point.x+1; i++ {
		for j := point.y - 1; j <= point.y+1; j++ {
			var pixel string

			if i < 0 || j < 0 || i >= len(image) || j >= len(image[i]) {
				pixel = FindInfinitePixel(algo, step)
			} else {
				pixel = image.Get(Point{i, j})
			}

			if pixel == "#" {
				neighbours_string += "1"
			} else {
				neighbours_string += "0"
			}
		}
	}
	return neighbours_string
}

func (binary Binary) ToInt() int64 {
	num, _ := strconv.ParseInt(string(binary), 2, 64)
	return num
}

func (image Image) PixelCount() int {
	num_pixels := 0
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			if image[i][j] == "#" {
				num_pixels += 1
			}
		}
	}
	return num_pixels
}

func (image Image) Enhance(algo string, step int) Image {
	enhanced := image.PaddedEmpty()

	enhanced.ForEach(func(point Point) {
		binary := image.Binary(Point{point.x - 1, point.y - 1}, algo, step)
		num := binary.ToInt()
		pixel := string(algo[num])
		enhanced.Set(point, pixel)
	})

	return enhanced
}

func main() {
	start := time.Now()
	f, _ := ioutil.ReadFile("input.txt")
	file_str := strings.Split(strings.TrimSpace(string(f)), "\n\n")
	enhancement_algo, image_str := file_str[0], file_str[1]
	image, pixel_count_2 := make(Image, 0), 0

	for _, image_line := range strings.Split(image_str, "\n") {
		image = append(image, strings.Split(image_line, ""))
	}

	for i := 0; i < 50; i++ {
		if i == 2 {
			pixel_count_2 = image.PixelCount()
		}

		image = image.Enhance(enhancement_algo, i)
	}

	pixel_count_50 := image.PixelCount()
	time_elapsed := time.Since(start)

	log.Printf(`
After 2 enhancements, there are %d lit pixels.
After 50 enhancements, there are %d lit pixels.
Solution generated in %s.`,
		pixel_count_2,
		pixel_count_50,
		time_elapsed,
	)
}
