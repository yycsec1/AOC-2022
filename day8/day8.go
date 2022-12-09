package day8

import (
	"fmt"
	"os"
	"strings"
)

func Solve1() {
	count := 0
	grid := parseInput("./day8/input.txt")
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if isVisible(i, j, &grid) {
				count++
			}
		}
	}
	fmt.Printf("The result is %d\n", count+2*len(grid)+2*len(grid[0])-4)
}

func Solve2() {
	score := 0
	grid := parseInput("./day8/input.txt")

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			curScore := getScenicScore(i, j, &grid)
			if curScore > score {
				score = curScore
			}
		}
	}
	fmt.Printf("The result is %d\n", score)
}

func parseInput(fileName string) (grid [][]rune) {
	byteData, _ := os.ReadFile(fileName)
	for _, row := range strings.Split(string(byteData), "\r\n") {
		var temp []rune
		for _, ch := range row {
			temp = append(temp, ch)
		}
		grid = append(grid, temp)
	}
	return
}

func isVisible(i, j int, grid *[][]rune) bool {
	visibleLeft, visibleRight, visibleUp, visibleDown := true, true, true, true
	//check left
	for a := 0; a < j; a++ {
		if (*grid)[i][a] >= (*grid)[i][j] {
			visibleLeft = false
			break
		}
	}
	//check right
	for a := len((*grid)[i]) - 1; a > j; a-- {
		if (*grid)[i][a] >= (*grid)[i][j] {
			visibleRight = false
			break
		}
	}
	//check up
	for a := 0; a < i; a++ {
		if (*grid)[a][j] >= (*grid)[i][j] {
			visibleUp = false
			break
		}
	}
	//check right
	for a := len(*grid) - 1; a > i; a-- {
		if (*grid)[a][j] >= (*grid)[i][j] {
			visibleDown = false
			break
		}
	}
	if visibleLeft || visibleRight || visibleUp || visibleDown {
		return true
	}
	return false
}

func getScenicScore(i, j int, grid *[][]rune) (score int) {
	//check left
	score = 1
	countLeft, countRight, countUp, countDown := 0, 0, 0, 0
	for a := j - 1; a >= 0; a-- {
		countLeft++
		if (*grid)[i][a] >= (*grid)[i][j] {
			break
		}
	}
	//check right
	for a := j + 1; a < len((*grid)[i]); a++ {
		countRight++
		if (*grid)[i][a] >= (*grid)[i][j] {
			break
		}
	}
	//check up
	for a := i - 1; a >= 0; a-- {
		countUp++
		if (*grid)[a][j] >= (*grid)[i][j] {
			break
		}

	}
	//check down
	for a := i + 1; a < len(*grid); a++ {
		countDown++
		if (*grid)[a][j] >= (*grid)[i][j] {
			break
		}

	}
	score = countLeft * countRight * countUp * countDown
	return
}
