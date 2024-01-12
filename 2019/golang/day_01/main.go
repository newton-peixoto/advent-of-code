package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args

	filename := args[1]

	bytes, err := os.ReadFile(filename)

	if err != nil {
		panic("no such file found!")
	}

	numbers := strings.Split(string(bytes), "\n")

	result := part01(numbers)

	fmt.Println("Part 01 : ", result)
	fmt.Println("Part 02 : ", part02(numbers))
}

func part01(numbers []string) int {
	total := 0

	for _, number := range numbers {
		n, _ := strconv.Atoi(number)

		total += n/3 - 2
	}

	return total
}

func part02(numbers []string) int {
	total := 0

	for _, number := range numbers {
		n, _ := strconv.Atoi(number)

		total += calculateFuel(n)
	}

	return int(total)
}

func calculateFuel(mass int) int {
	if mass/3-2 <= 0 {
		return 0
	}
	return mass/3 - 2 + calculateFuel(mass/3-2)
}
