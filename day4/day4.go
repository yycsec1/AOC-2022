package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve1() {
	result := 0

	pairsList := parseInput("./day4/input.txt")
	for _, pair := range pairsList {
		if (pair[0] <= pair[2] && pair[3] <= pair[1]) || (pair[2] <= pair[0] && pair[1] <= pair[3]) {
			result++
		}
	}

	fmt.Printf("The result is %d\n", result)
}

func Solve2() {
	result := 0

	pairsList := parseInput("./day4/input.txt")
	for _, pair := range pairsList {
		if (pair[0] <= pair[2] && pair[2] <= pair[1]) || (pair[0] <= pair[3] && pair[3] <= pair[1]) || (pair[2] <= pair[0] && pair[0] <= pair[3]) || (pair[2] <= pair[1] && pair[1] <= pair[3]) {
			result++

		}
	}

	fmt.Printf("The result is %d\n", result)
}

func parseInput(fileName string) (pairsList [][]int) {
	byteData, _ := os.ReadFile(fileName)

	for _, line := range strings.Split(string(byteData), "\r\n") {
		temp1 := strings.Split(line, ",")
		temp2 := strings.Split(temp1[0], "-")
		temp3 := strings.Split(temp1[1], "-")
		x1, _ := strconv.Atoi(temp2[0])
		y1, _ := strconv.Atoi(temp2[1])
		x2, _ := strconv.Atoi(temp3[0])
		y2, _ := strconv.Atoi(temp3[1])
		pairsList = append(pairsList, []int{x1, y1, x2, y2})
	}

	return
}
