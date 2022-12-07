package day6

import (
	"fmt"
	"os"
)

func Solve1() {
	buffer := parseInput("./day6/input.txt")
	for i := 0; i < len(buffer); i++ {
		if len(stringSet(buffer[i:i+4])) == 4 {
			fmt.Printf("The result is %d\n", i+4)
			break
		}
	}
}

func Solve2() {
	buffer := parseInput("./day6/input.txt")
	for i := 0; i < len(buffer); i++ {
		if len(stringSet(buffer[i:i+14])) == 14 {
			fmt.Printf("The result is %d\n", i+14)
			break
		}
	}
}

func parseInput(fileName string) (input string) {
	byteData, _ := os.ReadFile(fileName)
	input = string(byteData)
	return
}

func stringSet(str string) map[rune]bool {
	strSet := make(map[rune]bool)
	for _, ch := range str {
		strSet[ch] = true
	}
	return strSet
}
