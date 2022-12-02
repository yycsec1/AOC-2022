package day1

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve1() {
	result := 0
	sum := 0

	parsedData := parseInput("./day1/sample.txt")
	for _, line := range parsedData {
		if line == "" {
			if sum > result {
				result = sum
			}
			sum = 0
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Unable to parse number %s", line)
		}

		sum += num
	}
	if sum > result {
		result = sum
	}

	fmt.Printf("The result is %d\n", result)
}

func Solve2() {
	var calories []int
	sum := 0

	parsedData := parseInput("./day1/sample.txt")
	for _, line := range parsedData {
		if line == "" {
			calories = append(calories, sum)
			sum = 0
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Unable to parse number %s\n", line)
		}

		sum += num
	}
	calories = append(calories, sum)

	len := len(calories)
	sort.Ints(calories)
	fmt.Printf("The result is %d", calories[len-1]+calories[len-2]+calories[len-3])
}

func parseInput(fileName string) []string {

	byteData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Unable to open file %s", fileName)
	}

	return strings.Split(string(byteData), "\r\n")
}
