package day5

import (
	"fmt"
	"golang.org/x/exp/slices"
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
		//With append() is it always better to use fill slice syntax (e.g. slice[min:max:capacity]) to avoid unexpected behaviour
		stacks[move[2]-1] = append(stacks[move[1]-1][0:move[0]:move[0]], stacks[move[2]-1]...)
		stacks[move[1]-1] = stacks[move[1]-1][move[0]:]
	}
	for i := 0; i < len(stacks); i++ {
		result += stacks[i][0]
	}
	fmt.Printf("The result is %s\n", result)
}

func parseInput(fileName string) (stacks [][]string, moves [][]int) {
	byteData, _ := os.ReadFile(fileName)
	stringData := strings.Split(string(byteData), "\r\n")

	index := slices.Index(stringData, "")
	//index := indexOf(stringData, "")
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
