package day14

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solve1() {
	paths, minX, maxX, maxY := parseInput("./day14/input.txt")
	result := 0

	//Initialize cave
	cave := make([][]bool, maxY+1)
	for i := 0; i < maxY+1; i++ {
		cave[i] = make([]bool, maxX-minX+1)
	}

	//Draw paths in cave
	for _, path := range paths {
		for i := 1; i < len(path); i++ {
			if path[i-1].x == path[i].x {
				dy := abs(path[i-1].y, path[i].y)
				sig := sign(path[i].y - path[i-1].y)
				x := path[i].x - minX
				for a := 0; a <= dy; a++ {
					y := path[i-1].y + a*sig
					cave[y][x] = true
				}
			}
			if path[i-1].y == path[i].y {
				dx := abs(path[i-1].x, path[i].x)
				sig := sign(path[i].x - path[i-1].x)
				for a := 0; a <= dx; a++ {
					x := path[i-1].x + a*sig - minX
					cave[path[i].y][x] = true
				}
			}
		}
	}
	for simSand(&cave, Point{500 - minX, 0}) {
		result++
	}
	fmt.Printf("The result is %d.\n", result)
	//printCave(cave)
}

func Solve2() {
	paths, _, _, maxY := parseInput("./day14/input.txt")
	height, width := maxY+3, 2*(maxY+2)+3
	entryPoint := Point{height + 1, 0}
	result := 0

	//Initialize cave
	cave := make([][]bool, height)
	for i := 0; i < height; i++ {
		cave[i] = make([]bool, width)
	}
	for i := 0; i < width; i++ {
		cave[height-1][i] = true
	}

	//Draw paths in cave
	for _, path := range paths {
		for i := 1; i < len(path); i++ {
			if path[i-1].x == path[i].x {
				dy := abs(path[i-1].y, path[i].y)
				sig := sign(path[i].y - path[i-1].y)
				x := path[i].x - (500 - entryPoint.x)
				for a := 0; a <= dy; a++ {
					y := path[i-1].y + a*sig
					cave[y][x] = true
				}
			}
			if path[i-1].y == path[i].y {
				dx := abs(path[i-1].x, path[i].x)
				sig := sign(path[i].x - path[i-1].x)
				for a := 0; a <= dx; a++ {
					x := path[i-1].x + a*sig - (500 - entryPoint.x)
					cave[path[i].y][x] = true
				}
			}
		}
	}

	for simSand(&cave, Point{entryPoint.x, 0}) {
		result++
	}
	fmt.Printf("The result is %d.\n", result)
	//printCave(cave)
}

func parseInput(fileName string) (paths [][]Point, minX, maxX, maxY int) {
	minX = math.MaxInt
	maxX, maxY = 0, 0
	byteData, _ := os.ReadFile(fileName)
	stringData := strings.ReplaceAll(string(byteData), "\r\n", "\n")
	for _, line := range strings.Split(stringData, "\n") {
		var points []Point
		for _, p := range strings.Split(line, " -> ") {
			nums := strings.Split(p, ",")
			x, _ := strconv.Atoi(nums[0])
			y, _ := strconv.Atoi(nums[1])
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
			if x < minX {
				minX = x
			}
			points = append(points, Point{x, y})
		}
		paths = append(paths, points)
	}
	return
}

type Point struct {
	x int
	y int
}

func abs(x, y int) (result int) {
	result = x - y
	if result < 0 {
		return -result
	}
	return
}

func sign(num int) (result int) {
	if num < 0 {
		result = -1
	} else if num > 0 {
		result = 1
	} else {
		result = 0
	}
	return
}

func printCave(cave [][]bool) {
	for _, row := range cave {
		line := ""
		for _, ch := range row {
			if ch {
				line += "#"
			} else {
				line += "."
			}

		}
		fmt.Println(line)
	}
}

func simSand(cave *[][]bool, point Point) bool {
	for {
		if (*cave)[point.y][point.x] {
			return false
		}
		if point.y == len(*cave) || point.x == 0 || point.x == len((*cave)[0]) {
			return false
		}
		if !(*cave)[point.y+1][point.x] {
			point.y += 1
			continue
		}
		if !(*cave)[point.y+1][point.x-1] {
			point.x -= 1
			point.y += 1
			continue
		}
		if !(*cave)[point.y+1][point.x+1] {
			point.x += 1
			point.y += 1
			continue
		}
		(*cave)[point.y][point.x] = true
		break
	}
	return true
}
