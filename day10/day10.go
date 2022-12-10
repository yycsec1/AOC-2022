package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
	x := 1
	cycles := 0
	screenText := ""
	i := 0
	for i < len(instructions) {
		switch instructions[i] {
		case "noop":
			renderCycle(&cycles, &screenText, &x)
			i++
		case "addx":
			renderCycle(&cycles, &screenText, &x)
			renderCycle(&cycles, &screenText, &x)
			value, _ := strconv.Atoi(instructions[i+1])
			x += value
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

func renderCycle(cycle *int, screenText *string, x *int) {
	*cycle++
	pixel := " "
	mod := *cycle % 40
	if 0 <= (mod-*x) && (mod-*x) <= 2 {
		pixel = "#"
	}
	*screenText += pixel
	if mod == 0 {
		*screenText += "\n"
	}
}
