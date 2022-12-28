package day13

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Pair struct {
	left  []interface{}
	right []interface{}
}

func Solve1() {
	result := 0
	pairs, _ := parseInput("./day13/input.txt")
	for i, pair := range pairs {
		ans := isOrdered(pair.left, pair.right)
		//fmt.Println(ans)
		if ans == 1 {
			result += i + 1
		}
	}
	fmt.Printf("The result is %d.\n", result)
}

func Solve2() {
	dp1Index := 0
	dp2Index := 0
	_, signals := parseInput("./day13/input.txt")
	var dp1 []interface{}
	var dp2 []interface{}
	_ = json.Unmarshal([]byte("[[2]]"), &dp1)
	_ = json.Unmarshal([]byte("[[6]]"), &dp2)
	for _, signal := range signals {
		ans1, ans2 := isOrdered(signal, dp1), isOrdered(signal, dp2)
		if ans1 == 1 {
			dp1Index++
		}
		if ans2 == 1 {
			dp2Index++
		}
	}
	if dp1Index < dp2Index {
		fmt.Printf("The result is %d.\n", (dp1Index+1)*(dp2Index+2))
	} else {
		fmt.Printf("The result is %d.\n", (dp2Index+1)*(dp1Index+2))
	}
}

func parseInput(fileName string) (pairs []Pair, signals [][]interface{}) {
	byteData, _ := os.ReadFile(fileName)
	stringData := strings.ReplaceAll(string(byteData), "\r\n", "\n")
	for _, section := range strings.Split(stringData, "\n\n") {
		var left []interface{}
		var right []interface{}
		lines := strings.Split(section, "\n")
		_ = json.Unmarshal([]byte(lines[0]), &left)
		_ = json.Unmarshal([]byte(lines[1]), &right)
		pairs = append(pairs, Pair{left, right})
		signals = append(signals, left, right)
	}
	return
}

func isOrdered(left, right []interface{}) int {
	i := 0
	for {
		if i == len(left) || i == len(right) {
			break
		}
		v1, v2 := reflect.TypeOf(left[i]), reflect.TypeOf(right[i])
		valueLeft := reflect.ValueOf(left[i])
		valueRight := reflect.ValueOf(right[i])
		if v1.Kind() == reflect.Slice && v2.Kind() == reflect.Slice {

			var tempLeft []interface{}
			var tempRight []interface{}
			for a := 0; a < valueLeft.Len(); a++ {
				tempLeft = append(tempLeft, valueLeft.Index(a).Interface())
			}
			for b := 0; b < valueRight.Len(); b++ {
				tempRight = append(tempRight, valueRight.Index(b).Interface())
			}
			ans := isOrdered(tempLeft, tempRight)
			if ans == 0 {
				i++
				continue
			}
			return ans
		}
		if v1.Kind() == reflect.Float64 && v2.Kind() == reflect.Float64 {
			if valueLeft.Float() == valueRight.Float() {
				i++
				continue
			} else if valueLeft.Float() < valueRight.Float() {
				return 1
			} else {
				return -1
			}
		}
		if v1.Kind() == reflect.Float64 && v2.Kind() == reflect.Slice {
			var tempLeft []interface{}
			var tempRight []interface{}
			for b := 0; b < valueRight.Len(); b++ {
				tempRight = append(tempRight, valueRight.Index(b).Interface())
			}
			tempLeft = append(tempLeft, valueLeft.Float())
			ans := isOrdered(tempLeft, tempRight)
			if ans == 0 {
				i++
				continue
			}
			return ans
		}
		if v1.Kind() == reflect.Slice && v2.Kind() == reflect.Float64 {
			var tempLeft []interface{}
			var tempRight []interface{}
			for a := 0; a < valueLeft.Len(); a++ {
				tempLeft = append(tempLeft, valueLeft.Index(a).Interface())
			}
			tempRight = append(tempRight, valueRight.Float())
			ans := isOrdered(tempLeft, tempRight)
			if ans == 0 {
				i++
				continue
			}
			return ans
		}
		i++
	}
	if len(left) < len(right) {
		return 1
	}
	if len(right) < len(left) {
		return -1
	}
	return 0
}
