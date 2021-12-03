package parts

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Part1(input_file string) {
	data, err := ioutil.ReadFile(input_file)
	check(err)
	lines := strings.Split(string(data), "\n")
	var horizontal = 0
	var depth = 0
	for _, line := range lines {
		split := strings.Split(string(line), " ")
		cmd := split[0]
		val, err := strconv.Atoi(strings.TrimSpace(split[1]))
		check(err)
		switch cmd {
		case "up":
			depth -= val
		case "down":
			depth += val
		case "forward":
			horizontal += val
		}
	}
	fmt.Printf("\tDepth: %d\n", depth)
	fmt.Printf("\tHorizontal Position: %d\n", horizontal)
	fmt.Printf("\tProduct: %d\n", horizontal*depth)
}
