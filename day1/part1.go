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

func main() {
	data, err := ioutil.ReadFile("./part1.txt")
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
	fmt.Println(count)

}
