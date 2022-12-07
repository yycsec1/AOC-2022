package day7

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solve1() {
	result := 0
	dirSizes := parseInput("./day7/input.txt")
	for _, size := range dirSizes {
		if size < 100000 {
			result += size
		}
	}
	fmt.Printf("The result is %d\n", result)
}

func Solve2() {
	result := math.MaxInt
	dirSizes := parseInput("./day7/input.txt")
	freeSpace := 70000000 - dirSizes["/"]
	for _, size := range dirSizes {
		if size+freeSpace > 30000000 {
			result = int(math.Min(float64(size), float64(result)))
		}
	}
	fmt.Printf("The result is %d\n", result)
}

func parseInput(fileName string) (dirSizes map[string]int) {
	byteData, _ := os.ReadFile(fileName)
	dirSizes = make(map[string]int)
	path := []string{}
	for _, line := range strings.Split(string(byteData), "\r\n") {
		temp := strings.Split(line, " ")
		if temp[1] == "cd" && temp[2] == ".." {
			path = path[:len(path)-1]
			continue
		}
		if temp[0] == "dir" || temp[1] == "ls" {
			continue
		}
		if temp[1] == "cd" {
			path = append(path, temp[2])
			continue
		}
		for i := 0; i < len(path); i++ {
			size, _ := strconv.Atoi(temp[0])
			dirSizes[strings.Join(path[:i+1], "/")] += size
		}
	}
	return
}
