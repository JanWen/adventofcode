package parts

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Part1(input_file string) int {
	data, err := ioutil.ReadFile(input_file)
	check(err)
	lines := strings.Split(string(data), "\n")
	var count = 0
	var prev = 0
	for _, val := range lines {
		i, err := strconv.Atoi(strings.TrimSpace(val))
		check(err)
		if prev > 0 && prev < i {
			count++
		}
		prev = i
	}
	return count
}
