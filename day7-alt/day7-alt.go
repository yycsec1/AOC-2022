package day7_alt

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

type Node struct {
	tag      string
	name     string
	size     int
	children []*Node
}

func Solve1() {
	root := parseInput("./day7/input.txt")
	root.getSize()
	fmt.Println(root.solve1())
}

func Solve2() {
	parseInput("./day7/sample.txt")
}

func parseInput(fileName string) (root Node) {
	byteData, _ := os.ReadFile(fileName)
	root = Node{"dir", "/", 0, []*Node{}}
	var path []*Node
	var cwd *Node
	for _, line := range strings.Split(string(byteData), "\r\n") {
		if line[0] == '$' {
			if line[2] == 'c' {
				temp := strings.Split(line, " ")
				if temp[2] == "/" {
					path = append(path, &root)
					cwd = &root
				} else if temp[2] == ".." {
					cwd = path[len(path)-1]
					path = path[:len(path)-1]
				} else {
					if !cwd.isChildOf(temp[2]) {
						cwd.children = append(cwd.children, &Node{"dir", temp[2], 0, []*Node{}})
					}
					path = append(path, cwd)
					cwd = cwd.getChildAddress(temp[2])
				}
			}
		} else {
			temp := strings.Split(line, " ")
			if temp[0] == "dir" {
				if !cwd.isChildOf(temp[1]) {
					cwd.children = append(cwd.children, &Node{"dir", temp[1], 0, []*Node{}})
				}
			} else {
				size, _ := strconv.Atoi(temp[0])
				cwd.children = append(cwd.children, &Node{"file", temp[1], size, nil})
			}
		}
	}
	return
}

func (node *Node) isChildOf(name string) bool {
	for _, child := range node.children {
		if child.name == name {
			return true
		}
	}
	return false
}

func (node *Node) getChildAddress(name string) *Node {
	for i := 0; i < len(node.children); i++ {
		if node.children[i].name == name {
			return node.children[i]
		}
	}
	return nil
}

func (node *Node) getSize() int {
	if node.tag == "file" {
		return node.size
	}
	if node.size != 0 {
		return node.size
	}
	sum := 0
	for i := 0; i < len(node.children); i++ {
		sum += node.children[i].getSize()
	}
	node.size = sum
	return node.size
}

func (node *Node) solve1() (size int) {
	stack := stack.New()
	stack.Push()

	for stack.Len() > 0 {
		cur := stack.Pop()
		if cur.tag == "dir" {
			for i := 0; i < len(cur.children); i++ {
				stack.Push(cur.children[i])
			}

			if cur.size >= 100000 {
				size += cur.size
			}
		}
	}
	return
}
