package main

import (
	day1 "advent/day1/parts"
	day2 "advent/day2/parts"
	"fmt"
)

func main() {
	day1part1 := day1.Part1("./day1/input/part1.txt")
	if day1part1 != 1393 {
		panic("Day 1 part 1 result wrong!")
	}
	fmt.Printf("Day1Part1: %d\n", day1part1)

	day1part2 := day1.Part2("./day1/input/part1.txt")
	if day1part2 != 1359 {
		panic("Day 1 part 2 result wrong!")
	}
	fmt.Printf("Day1Part2: %d\n", day1part2)

	day2.Part1("./day2/input/part1.txt")

}
