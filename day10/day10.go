package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func Solve1() {
	instructions := parseInput("./day10/input.txt")
	signalStrength := 0
	x := 1
	cycles := 0
	i := 0
	for i < len(instructions) {
		switch instructions[i] {
		case "noop":
			doCycle(&cycles, &x, &signalStrength)
			i++
		case "addx":
			doCycle(&cycles, &x, &signalStrength)
			doCycle(&cycles, &x, &signalStrength)
			value, _ := strconv.Atoi(instructions[i+1])
			x += value
			i += 2
		}
	}
	fmt.Printf("The result is %d\n", signalStrength)
}

func Solve2() {
	instructions := parseInput("./day10/input.txt")
	sprite := []int{1, 2, 3}
	cycles := 0
	screenText := ""
	i := 0
	for i < len(instructions) {
		switch instructions[i] {
		case "noop":
			renderCycle(&cycles, &screenText, &sprite)
			i++
		case "addx":
			renderCycle(&cycles, &screenText, &sprite)
			renderCycle(&cycles, &screenText, &sprite)
			value, _ := strconv.Atoi(instructions[i+1])
			for j := 0; j < len(sprite); j++ {
				sprite[j] += value
			}
			i += 2
		}
	}
	fmt.Println(screenText)
}

func parseInput(fileName string) (instructions []string) {
	byteData, _ := os.ReadFile(fileName)
	return strings.Fields(string(byteData))
}

func doCycle(cycle *int, x *int, signalStrength *int) {
	*cycle++
	if *cycle == 20 || *cycle == 60 || *cycle == 100 || *cycle == 140 || *cycle == 180 || *cycle == 220 {
		*signalStrength += *cycle * *x
	}
}

func renderCycle(cycle *int, screenText *string, sprite *[]int) {
	*cycle++
	pixel := "."
	mod := *cycle % 40
	if slices.Contains(*sprite, mod) {
		pixel = "#"
	}
	*screenText += pixel
	if mod == 0 {
		*screenText += "\n"
	}
}
