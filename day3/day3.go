package day3

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Rucksack struct {
	compOne, compTwo string
}

type Group struct {
	rucksacks [3]string
	indexes   [3]int
}

func parseInputOne(fileName string) (rucksacks []Rucksack) {
	byteData, _ := os.ReadFile(fileName)

	for _, line := range strings.Split(string(byteData), "\r\n") {
		half := len(line) / 2
		rucksacks = append(rucksacks, Rucksack{line[:half], line[half:]})
	}

	return
}

func parseInputTwo(fileName string) (groups []Group) {
	byteData, _ := os.ReadFile(fileName)

	stringData := strings.Split(string(byteData), "\r\n")
	for i := 0; i < len(stringData); i += 3 {
		indexes := [3]int{0, 1, 2}

		if len(stringData[i+1]) < len(stringData[i+2]) {
			if len(stringData[i+1]) < len(stringData[i]) {
				indexes = [3]int{1, 0, 2}
			}
		} else {
			if len(stringData[i+2]) < len(stringData[i]) {
				indexes = [3]int{2, 0, 1}
			}
		}

		groups = append(groups, Group{[3]string{stringData[i], stringData[i+1], stringData[i+2]}, indexes})
	}

	return
}

func Solve1() {
	priorities := 0
	rucksacks := parseInputOne("./day3/input.txt")

	for _, rucksack := range rucksacks {
		for _, ch := range rucksack.compOne {
			if !strings.ContainsRune(rucksack.compTwo, ch) {
				continue
			}

			if unicode.IsLower(ch) {
				priorities += int(ch) - 96
				break
			}
			priorities += int(ch) - 38
			break

		}
	}

	fmt.Printf("The result is %d\n", priorities)
}

func Solve2() {
	priorities := 0
	groups := parseInputTwo("./day3/input.txt")

	for _, group := range groups {
		for _, ch := range group.rucksacks[group.indexes[0]] {
			if !(strings.ContainsRune(group.rucksacks[group.indexes[1]], ch) && strings.ContainsRune(group.rucksacks[group.indexes[2]], ch)) {
				continue
			}

			if unicode.IsLower(ch) {
				priorities += int(ch) - 96
				break
			}
			priorities += int(ch) - 38
			break

		}
	}

	fmt.Printf("The result is %d\n", priorities)
}
