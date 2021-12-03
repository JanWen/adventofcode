package main

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

func get_measurement(linenr int, lines []string) int {
	measurement, err := strconv.Atoi(strings.TrimSpace(lines[linenr]))
	check(err)
	return measurement
}

func main() {
	data, err := ioutil.ReadFile("./part1.txt")
	check(err)
	lines := strings.Split(string(data), "\n")
	var count = 0

	for index, _ := range lines[:len(lines)-3] {
		// first, err := '
		window1 := get_measurement(index, lines) + get_measurement(index+1, lines) + get_measurement(index+2, lines)
		check(err)
		window2 := get_measurement(index+1, lines) + get_measurement(index+2, lines) + get_measurement(index+3, lines)
		check(err)
		if window1 < window2 {
			count++
		}
	}
	fmt.Println(count)

}
