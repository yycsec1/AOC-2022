package day15

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Y = 2000000
const L = 4000000

func Solve1() {
	result := make(map[Point]bool)
	mapData := parseInput("./day15/input.txt")
	for _, data := range mapData {
		dist := distance(data.sensor, data.beacon)
		xDist := dist - abs(data.sensor.y-Y)
		if xDist > 0 {
			for x := data.sensor.x - xDist; x <= data.sensor.x+xDist; x++ {
				p := Point{x, Y}
				if !isBeacon(&mapData, &p) {
					result[p] = true
				}
			}
		}
	}
	fmt.Printf("The result is %d.\n", len(result))
}

func Solve2() {
	var distances []int
	mapData := parseInput("./day15/input.txt")
	for _, data := range mapData {
		distances = append(distances, distance(data.sensor, data.beacon))
	}
OUTER:
	for y := 0; y <= L; y++ {
		x := 0
	INNER:
		for x <= L {
			for i, data := range mapData {
				d := distance(data.sensor, Point{x, y})
				if d <= distances[i] {
					x += (distances[i] - d) + 1
					continue INNER
				}
			}
			fmt.Printf("The result is %d.\n", x*L+y)
			break OUTER
		}
	}
}

func parseInput(fileName string) (mapData []Data) {
	byteData, _ := os.ReadFile(fileName)
	stringData := strings.ReplaceAll(string(byteData), "\r\n", "\n")
	for _, line := range strings.Split(stringData, "\n") {
		sections := strings.Split(line, " ")
		t := strings.Split(sections[2], "=")[1]
		x1, _ := strconv.Atoi(t[:len(t)-1])
		t = strings.Split(sections[3], "=")[1]
		y1, _ := strconv.Atoi(t[:len(t)-1])
		t = strings.Split(sections[8], "=")[1]
		x2, _ := strconv.Atoi(t[:len(t)-1])
		y2, _ := strconv.Atoi(strings.Split(sections[9], "=")[1])
		mapData = append(mapData, Data{
			sensor: Point{x: x1, y: y1},
			beacon: Point{x: x2, y: y2},
		})
	}
	return
}

type Point struct {
	x int
	y int
}

type Data struct {
	sensor Point
	beacon Point
}

func abs(num int) (result int) {
	result = num
	if result < 0 {
		return -result
	}
	return result
}

func distance(a, b Point) (dist int) {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func isBeacon(mapData *[]Data, point *Point) bool {
	for _, data := range *mapData {
		if data.beacon == *point {
			return true
		}
	}
	return false
}
