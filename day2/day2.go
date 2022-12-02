package day2

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parseInput(fileName string) (opponentChoice []uint8, yourChoices []uint8) {
	byteDate, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Unable to open file %s\n", fileName)
	}

	stringData := strings.Split(string(byteDate), "\r\n")

	for _, line := range stringData {
		opponentChoice = append(opponentChoice, line[0])
		yourChoices = append(yourChoices, line[2])
	}

	return
}

/*
65 88 Rock (win 21, draw 23, lose 22)
66 89 Paper (win 24, draw 23, lose 22)
67 90 Scissor (win 24, draw 23, lose 25)
*/

func getGamePoints(opponentChoice uint8, yourChoice uint8) (points int) {
	diff := yourChoice - opponentChoice
	if diff == 23 {
		//Draw
		points = 3
		return
	}
	if diff == 21 || diff == 24 {
		//Win
		points = 6
		return
	}
	points = 0

	return
}

type Choice struct {
	opponentChoice uint8
	gameResult     uint8
}

func Solve1() {

	opponentChoice, yourChoice := parseInput("./day2/input.txt")
	iterations := len(opponentChoice)

	totalPoints := 0

	for i := 0; i < iterations; i++ {
		totalPoints += int(yourChoice[i]-87) + getGamePoints(opponentChoice[i], yourChoice[i])
	}

	fmt.Printf("The result is %d\n", totalPoints)
}

func Solve2() {
	gamePoints := map[uint8]int{
		88: 0,
		89: 3,
		90: 6,
	}

	choicePoints := map[Choice]int{
		Choice{65, 88}: 3,
		Choice{66, 88}: 1,
		Choice{67, 88}: 2,
		Choice{65, 89}: 1,
		Choice{66, 89}: 2,
		Choice{67, 89}: 3,
		Choice{65, 90}: 2,
		Choice{66, 90}: 3,
		Choice{67, 90}: 1,
	}

	opponentChoice, gameResult := parseInput("./day2/input.txt")
	iterations := len(opponentChoice)

	totalPoints := 0

	for i := 0; i < iterations; i++ {
		totalPoints += choicePoints[Choice{opponentChoice[i], gameResult[i]}] + gamePoints[gameResult[i]]
	}

	fmt.Printf("The result is %d\n", totalPoints)
}
