package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve1() {
	result := ""
	stacks, moves := parseInput("./day5/input.txt")
	for _, move := range moves {
		for i := 0; i < move[0]; i++ {
			push(&stacks[move[2]-1], pop(&stacks[move[1]-1]))
		}
	}
	for i := 0; i < len(stacks); i++ {
		result += stacks[i][0]
	}
	fmt.Printf("The result is %s\n", result)
}

func Solve2() {
	result := ""
	stacks, moves := parseInput("./day5/input.txt")
	for _, move := range moves {
		var temp []string
		for i := 0; i < move[0]; i++ {
			temp = append(temp, pop(&stacks[move[1]-1]))
		}
		for i := len(temp); i > 0; i-- {
			push(&stacks[move[2]-1], temp[i-1])
		}
	}
	for i := 0; i < len(stacks); i++ {
		result += stacks[i][0]
	}
	fmt.Printf("The result is %s\n", result)
}

func parseInput(fileName string) (stacks [][]string, moves [][]int) {
	byteData, _ := os.ReadFile(fileName)
	stringData := strings.Split(string(byteData), "\r\n")

	index := indexOf(stringData, "")
	stackCount := len(strings.Fields(stringData[index-1]))
	for i := 0; i < stackCount; i++ {
		stacks = append(stacks, []string{})
	}

	for _, line := range stringData[:index-1] {
		for i := 1; i < len(line); i += 4 {
			if line[i] != 32 {
				stacks[(i-1)/4] = append(stacks[(i-1)/4], string(line[i]))
			}
		}
	}

	for _, line := range stringData[index+1:] {
		temp := strings.Split(line, " ")
		a, _ := strconv.Atoi(temp[1])
		b, _ := strconv.Atoi(temp[3])
		c, _ := strconv.Atoi(temp[5])
		moves = append(moves, []int{a, b, c})
	}

	return
}

func push(list *[]string, ele string) {
	*list = append([]string{ele}, *list...)
}

func pop(list *[]string) (ele string) {
	ele = (*list)[0]
	*list = (*list)[1:]
	return
}

func indexOf(array []string, item string) (index int) {
	index = -1
	if len(array) == 0 {
		return
	}
	for i, ele := range array {
		if ele == item {
			index = i
			return
		}
	}
	return
}
