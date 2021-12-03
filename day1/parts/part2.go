package parts

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func get_measurement(linenr int, lines []string) int {
	measurement, err := strconv.Atoi(strings.TrimSpace(lines[linenr]))
	check(err)
	return measurement
}

func Part2(input_file string) int {
	data, err := ioutil.ReadFile(input_file)
	check(err)
	lines := strings.Split(string(data), "\n")
	var count = 0

	for index := range lines[:len(lines)-3] {
		// first, err := '
		window1 := get_measurement(index, lines) + get_measurement(index+1, lines) + get_measurement(index+2, lines)
		check(err)
		window2 := get_measurement(index+1, lines) + get_measurement(index+2, lines) + get_measurement(index+3, lines)
		check(err)
		if window1 < window2 {
			count++
		}
	}
	return count
}
