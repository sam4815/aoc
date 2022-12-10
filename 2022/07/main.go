package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Dir struct {
	name     string
	size     int
	parent   *Dir
	children []interface{}
}

type File struct {
	name   string
	size   int
	parent *Dir
}

func (parent *Dir) AppendFile(name string, size int) {
	child := &File{name: name, size: size, parent: parent}
	parent.children = append(parent.children, child)
}

func (parent *Dir) AppendDirectory(name string) {
	child := &Dir{name: name, parent: parent}
	parent.children = append(parent.children, child)
}

func (parent *Dir) CalculateDirectorySizes() int {
	size := 0

	for _, child := range parent.children {
		if _, ok := child.(*Dir); ok {
			size += child.(*Dir).CalculateDirectorySizes()
		} else {
			size += child.(*File).size
		}
	}

	parent.size = size
	return size
}

func (parent *Dir) SumDirectoriesSmallerThan(size int) int {
	sum := 0

	if parent.size < size {
		sum += parent.size
	}

	for _, child := range parent.children {
		if _, ok := child.(*Dir); ok {
			sum += child.(*Dir).SumDirectoriesSmallerThan(size)
		}
	}

	return sum
}

func (parent *Dir) SmallestDirectoryLargerThan(size int) int {
	smallest := 0

	if parent.size > size {
		smallest = parent.size
	} else {
		return 0
	}

	for _, child := range parent.children {
		if _, ok := child.(*Dir); ok {
			smallest_from_dir := child.(*Dir).SmallestDirectoryLargerThan(size)
			if smallest_from_dir != 0 && smallest_from_dir < smallest {
				smallest = smallest_from_dir
			}
		}
	}

	return smallest
}

func (parent *Dir) Print(indent_level int) {
	log.Printf(
		"%s\x1b[32m- %s\033[0m (dir, size=%d)\n",
		strings.Repeat(" ", indent_level*2),
		parent.name,
		parent.size,
	)

	for _, child := range parent.children {
		if _, ok := child.(*Dir); ok {
			child.(*Dir).Print(indent_level + 1)
		} else {
			log.Printf(
				"%s  \033[31m- %s\033[0m (file, size=%d)\n",
				strings.Repeat(" ", indent_level*2),
				child.(*File).name,
				child.(*File).size,
			)
		}
	}
}

func (current_directory *Dir) ChangeDirectory(directory_name string) *Dir {
	if directory_name == "/" {
		for current_directory.parent != nil {
			current_directory = current_directory.parent
		}
	} else if directory_name == ".." {
		current_directory = current_directory.parent
	} else {
		for _, child := range current_directory.children {
			if dir, ok := child.(*Dir); ok {
				if dir.name == directory_name {
					current_directory = child.(*Dir)
				}
			}
		}
	}

	return current_directory
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	file_system := &Dir{name: "/"}

	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")

		if command[0] == "$" {
			if command[1] == "cd" {
				file_system = file_system.ChangeDirectory(command[2])
			}
			continue
		}

		if command[0] == "dir" {
			file_system.AppendDirectory(command[1])
			continue
		}

		size, err := strconv.Atoi(command[0])
		if err != nil {
			log.Fatal(err)
		}

		file_system.AppendFile(command[1], size)
	}

	file_system = file_system.ChangeDirectory("/")
	file_system.CalculateDirectorySizes()
	// file_system.Print(0)

	sum_small_directories := file_system.SumDirectoriesSmallerThan(100001)
	required_space := 70000000 - 30000000
	deletable_directory_size := file_system.SmallestDirectoryLargerThan(file_system.size - required_space)

	time_elapsed := time.Since(start)

	log.Printf(`
The sum of the total sizes of directories smaller than 100001 is %d.
The size of the directory that should be deleted is %d.
Solution generated in %s.`,
		sum_small_directories,
		deletable_directory_size,
		time_elapsed,
	)
}
