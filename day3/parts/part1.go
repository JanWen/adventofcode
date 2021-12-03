package parts

import (
	"fmt"
	"io/ioutil"
	"math"
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
	var gamma_values []int
	for range strings.TrimSpace(lines[0]) {
		gamma_values = append(gamma_values, 0)
	}
	for _, line := range lines {
		val := strings.TrimSpace(line)
		for i, bit := range val {
			switch bit {
			case '1':
				gamma_values[i] += 1
			case '0':
				gamma_values[i] -= 1
			}
		}
	}

	gamma := 0
	for _, val := range gamma_values {
		if val > 0 {
			gamma = (gamma << 1) | 1
		} else {
			gamma = gamma << 1
		}
	}

	epsilon := int((math.Pow(2, float64(len(gamma_values))))-1) ^ gamma
	fmt.Printf("\tGamma: %d\n", gamma)
	fmt.Printf("\tEpsilon: %d\n", epsilon)
	fmt.Printf("\tProduct: %d\n", gamma*epsilon)
}
