package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	x int
	y int
}

type Rope struct {
	nodes []Node
}

type Motion struct {
	dir   uint8
	steps int
}

type Index struct {
	x int
	y int
}

func Solve1() {
	rope := Rope{[]Node{}}
	for i := 0; i < 2; i++ {
		rope.nodes = append(rope.nodes, Node{0, 0})
	}

	visited := make(map[Index]bool)
	motions := parseInput("./day9/input.txt")

	for _, motion := range motions {
		newlyVisited := (&rope).move(&motion)
		for k, _ := range newlyVisited {
			visited[k] = true
		}
	}
	fmt.Println(len(visited))
}

func Solve2() {
	rope := Rope{[]Node{}}
	for i := 0; i < 10; i++ {
		rope.nodes = append(rope.nodes, Node{0, 0})
	}

	visited := make(map[Index]bool)
	motions := parseInput("./day9/input.txt")

	for _, motion := range motions {
		newlyVisited := (&rope).move(&motion)
		for k, _ := range newlyVisited {
			visited[k] = true
		}
	}
	fmt.Println(len(visited))
}

func parseInput(fileName string) (motions []Motion) {
	byteData, _ := os.ReadFile(fileName)
	moves := strings.Fields(string(byteData))
	for i := 0; i < len(moves); i += 2 {
		steps, _ := strconv.Atoi(moves[i+1])
		motions = append(motions, Motion{moves[i][0], steps})
	}
	return
}
func (rope *Rope) move(motion *Motion) (visited map[Index]bool) {
	visited = make(map[Index]bool)
	l := len(rope.nodes)
	for i := 0; i < motion.steps; i++ {
		switch motion.dir {
		case 82: //right
			rope.nodes[0].x++
		case 76: //left
			rope.nodes[0].x--
		case 85: //up
			rope.nodes[0].y++
		case 68: //down
			rope.nodes[0].y--
		}
		for j := 1; j < len(rope.nodes); j++ {
			if abs(rope.nodes[j].x-rope.nodes[j-1].x) == 2 && abs(rope.nodes[j].y-rope.nodes[j-1].y) == 2 {
				rope.nodes[j].x -= sign(rope.nodes[j].x - rope.nodes[j-1].x)
				rope.nodes[j].y -= sign(rope.nodes[j].y - rope.nodes[j-1].y)
				continue
			}
			if abs(rope.nodes[j].x-rope.nodes[j-1].x) == 2 {
				rope.nodes[j].x -= sign(rope.nodes[j].x - rope.nodes[j-1].x)
				rope.nodes[j].y = rope.nodes[j-1].y
				continue
			}
			if abs(rope.nodes[j].y-rope.nodes[j-1].y) == 2 {
				rope.nodes[j].y -= sign(rope.nodes[j].y - rope.nodes[j-1].y)
				rope.nodes[j].x = rope.nodes[j-1].x
				continue
			}
		}
		visited[Index{rope.nodes[l-1].x, rope.nodes[l-1].y}] = true
	}
	return
}

func sign(num int) int {
	if num < 0 {
		return -1
	}
	if num > 0 {
		return 1
	}
	return 0
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
