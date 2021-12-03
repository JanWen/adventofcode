package parts

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Part2(input_file string) {
	data, err := ioutil.ReadFile(input_file)
	check(err)
	lines := strings.Split(string(data), "\n")
	var horizontal = 0
	var depth = 0
	var aim = 0
	for _, line := range lines {
		split := strings.Split(string(line), " ")
		cmd := split[0]
		val, err := strconv.Atoi(strings.TrimSpace(split[1]))
		check(err)
		switch cmd {
		case "up":
			aim -= val
		case "down":
			aim += val
		case "forward":
			horizontal += val
			depth += aim * val
		}
	}
	fmt.Printf("\tDepth: %d\n", depth)
	fmt.Printf("\tHorizontal Position: %d\n", horizontal)
	fmt.Printf("\tProduct: %d\n", horizontal*depth)
}
