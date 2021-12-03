package parts

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func greaterThen(a int, b int) bool { return a > b }

func lessThen(a int, b int) bool { return a < b }

func recurse(numbers []string, position int, operator func(a int, b int) bool, default_val string) string {
	if len(numbers) == 1 {
		return numbers[0]
	}
	zeros := 0
	ones := 0
	for _, number := range numbers {
		if number[position] == '1' {
			ones += 1
		} else {
			zeros += 1
		}
	}

	var selector string
	if zeros == ones {
		selector = default_val
	} else if operator(zeros, ones) {
		selector = "0"
	} else {
		selector = "1"
	}

	new_array := []string{}
	for _, number := range numbers {
		if string(number[position]) == selector {
			new_array = append(new_array, number)
		}
	}

	return recurse(new_array, position+1, operator, default_val)
}

func Part2(input_file string) {
	data, err := ioutil.ReadFile(input_file)
	check(err)
	lines := strings.Split(string(data), "\n")
	trimmed := []string{}
	for _, line := range lines {
		val := strings.TrimSpace(line)
		trimmed = append(trimmed, val)
	}
	oxygen_generator_rating, err := strconv.ParseInt(recurse(trimmed, 0, greaterThen, "1"), 2, 0)
	check(err)
	co2_scrubber_rating, err := strconv.ParseInt(recurse(trimmed, 0, lessThen, "0"), 2, 0)
	check(err)
	fmt.Printf("oxygen generator rating: %d\n", oxygen_generator_rating)
	fmt.Printf("CO2 scrubber rating: %d\n", co2_scrubber_rating)
	fmt.Printf("Product: %d\n", int(co2_scrubber_rating)*int(oxygen_generator_rating))
}
