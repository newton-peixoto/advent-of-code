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

	numbers := strings.Split(string(bytes), ",")
	nList := make([]int, len(numbers))

	for index, number := range numbers {
		n, _ := strconv.Atoi(number)
		nList[index] = n
	}

	part1 := part01(nList, 12, 2)

	fmt.Println("Part01 : ", part1[0])

	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			for index, number := range numbers {
				n, _ := strconv.Atoi(number)
				nList[index] = n
			}
			part02 := part01(nList, x, y)

			if part02[0] == 19690720 {
				fmt.Println("Part02 : ", 100*part02[1]+part02[2])
			}
		}
	}

}

func part01(nList []int, one int, two int) []int {
	nList[1] = one
	nList[2] = two

	opcode := 0

	for opcode < len(nList) {

		if nList[opcode] == 99 {
			break
		}

		if nList[opcode] == 1 {
			firstInput := nList[nList[opcode+1]]
			secondInput := nList[nList[opcode+2]]
			result := firstInput + secondInput
			nList[nList[opcode+3]] = result
			opcode += 4
			continue
		}

		if nList[opcode] == 2 {
			firstInput := nList[nList[opcode+1]]
			secondInput := nList[nList[opcode+2]]
			result := firstInput * secondInput
			nList[nList[opcode+3]] = result
			opcode += 4
			continue
		}
	}

	return nList
}
