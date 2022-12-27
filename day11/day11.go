package day11

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve1() {
	monkeys := parseInput("./day11/input.txt")
	for a := 0; a < 20; a++ {
		for i := 0; i < len(monkeys); i++ {
			monkeys[i].inspectedItems += len(monkeys[i].itemsWorryLevel)
			for _, itemWorryLevel := range monkeys[i].itemsWorryLevel {
				newItemWorryLevel := eval(monkeys[i].operation, itemWorryLevel)
				newItemWorryLevel = newItemWorryLevel / 3
				if newItemWorryLevel%monkeys[i].test[0] == 0 {
					monkeys[monkeys[i].test[1]].itemsWorryLevel = append(monkeys[monkeys[i].test[1]].itemsWorryLevel, newItemWorryLevel)
				} else {
					monkeys[monkeys[i].test[2]].itemsWorryLevel = append(monkeys[monkeys[i].test[2]].itemsWorryLevel, newItemWorryLevel)
				}
			}
			monkeys[i].itemsWorryLevel = []int{}
		}
	}

	var itemsInspected []int
	for _, monkey := range monkeys {
		itemsInspected = append(itemsInspected, monkey.inspectedItems)
	}
	sort.Slice(itemsInspected, func(i, j int) bool {
		return itemsInspected[i] > itemsInspected[j]
	})
	fmt.Printf("The result is %d.\n", itemsInspected[0]*itemsInspected[1])
}

func Solve2() {
	monkeys := parseInput("./day11/input.txt")
	mod := getMod(monkeys)
	for a := 0; a < 10000; a++ {
		for i := 0; i < len(monkeys); i++ {
			monkeys[i].inspectedItems += len(monkeys[i].itemsWorryLevel)
			for _, itemWorryLevel := range monkeys[i].itemsWorryLevel {
				newItemWorryLevel := eval(monkeys[i].operation, itemWorryLevel)
				newItemWorryLevel %= mod
				if newItemWorryLevel%monkeys[i].test[0] == 0 {
					monkeys[monkeys[i].test[1]].itemsWorryLevel = append(monkeys[monkeys[i].test[1]].itemsWorryLevel, newItemWorryLevel)
				} else {
					monkeys[monkeys[i].test[2]].itemsWorryLevel = append(monkeys[monkeys[i].test[2]].itemsWorryLevel, newItemWorryLevel)
				}
			}
			monkeys[i].itemsWorryLevel = []int{}
		}
	}

	var itemsInspected []int
	for _, monkey := range monkeys {
		itemsInspected = append(itemsInspected, monkey.inspectedItems)
	}
	sort.Slice(itemsInspected, func(i, j int) bool {
		return itemsInspected[i] > itemsInspected[j]
	})
	fmt.Printf("The result is %d.\n", itemsInspected[0]*itemsInspected[1])
}

func parseInput(filename string) (monkeys []Monkey) {
	byteData, _ := os.ReadFile(filename)
	stringData := strings.ReplaceAll(string(byteData), "\r\n", "\n")
	for _, section := range strings.Split(stringData, "\n\n") {
		sectionLines := strings.Split(section, "\n")
		monkeys = append(monkeys, Monkey{
			parseItemsWorryLevel(sectionLines[1]),
			parseOperation(sectionLines[2]),
			parseTest([]string{sectionLines[3], sectionLines[4], sectionLines[5]}),
			0,
		})
	}
	return
}

func parseItemsWorryLevel(line string) (worryLevels []int) {
	line = strings.Split(line, ":")[1]
	for _, n := range strings.Split(line, ",") {
		num, _ := strconv.Atoi(strings.TrimSpace(n))
		worryLevels = append(worryLevels, num)
	}
	return
}

func parseOperation(line string) (operation []string) {
	line = strings.TrimSpace(strings.Split(line, "=")[1])
	operation = strings.Split(line, " ")
	return
}

func parseTest(lines []string) (test []int) {
	for _, line := range lines {
		n := strings.Split(line, " ")
		num, _ := strconv.Atoi(n[len(n)-1])
		test = append(test, num)
	}
	return
}

func eval(operation []string, itemWorryLevel int) (newItemWorryLevel int) {
	if operation[1] == "+" {
		if operation[2] != "old" {
			num, _ := strconv.Atoi(operation[2])
			newItemWorryLevel = itemWorryLevel + num
		}

		if operation[2] == "old" {
			newItemWorryLevel = itemWorryLevel + itemWorryLevel
		}
	}

	if operation[1] == "*" {
		if operation[2] != "old" {
			num, _ := strconv.Atoi(operation[2])
			newItemWorryLevel = itemWorryLevel * num
		}

		if operation[2] == "old" {
			newItemWorryLevel = itemWorryLevel * itemWorryLevel
		}
	}
	return
}

func getMod(monkeys []Monkey) (mod int) {
	mod = 1
	for _, monkey := range monkeys {
		mod *= monkey.test[0]
	}
	return
}

type Monkey struct {
	itemsWorryLevel []int
	operation       []string
	test            []int
	inspectedItems  int
}
