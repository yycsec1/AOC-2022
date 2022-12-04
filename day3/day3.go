package day3

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func parseInput(fileName string) (rucksacks, groups [][]string) {
	byteData, _ := os.ReadFile(fileName)

	stringData := strings.Split(string(byteData), "\r\n")
	for i := 0; i < len(stringData); i += 3 {

		rucksacks = append(rucksacks,
			[]string{stringData[i][:len(stringData[i])/2], stringData[i][len(stringData[i])/2:]},
			[]string{stringData[i+1][:len(stringData[i+1])/2], stringData[i+1][len(stringData[i+1])/2:]},
			[]string{stringData[i+2][:len(stringData[i+2])/2], stringData[i+2][len(stringData[i+2])/2:]},
		)
		groups = append(groups, []string{stringData[i], stringData[i+1], stringData[i+2]})
	}
	return
}

func score(ch rune) int {
	if unicode.IsLower(ch) {
		return int(ch) - 96
	}
	return int(ch) - 38
}

func Solve1() {
	priorities := 0
	rucksacks, _ := parseInput("./day3/input.txt")

	for _, rucksack := range rucksacks {
		for _, ch := range rucksack[0] {
			if !strings.ContainsRune(rucksack[1], ch) {
				continue
			}

			priorities += score(ch)
			break

		}
	}

	fmt.Printf("The result is %d\n", priorities)
}

func Solve2() {
	priorities := 0
	_, groups := parseInput("./day3/input.txt")

	for _, group := range groups {
		for _, ch := range group[0] {
			if !(strings.ContainsRune(group[1], ch) && strings.ContainsRune(group[2], ch)) {
				continue
			}

			priorities += score(ch)
			break
		}
	}

	fmt.Printf("The result is %d\n", priorities)
}
