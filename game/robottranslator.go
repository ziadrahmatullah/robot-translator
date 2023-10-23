// Name	: Muhammad Ziad Rahmatullah
// Date	: Oct 10, 21:19

package game

import (
	"fmt"
	"regexp"
)

func instructions(count int, move string) string {
	var mapTime = map[bool]string{
		true:  "times",
		false: "time",
	}

	var mapMove = map[string]string{
		"R": "right",
		"L": "left",
		"A": "advance",
	}

	timeWord := false
	if count > 1 {
		timeWord = true
	}

	return fmt.Sprintf("Move %s %d % s", mapMove[move], count, mapTime[timeWord])
}

func isMovesValid(moves string) bool {
	valid := regexp.MustCompile(`^[RLA]+$`)
	return valid.MatchString(moves)
}

func calculateMove(moves string) (result []string) {
	if isMovesValid(moves) {
		count := 0
		tempMove := ""
		for _, move := range moves {
			move := string(move)
			if tempMove == "" {
				tempMove = move
				count = 1
			} else if tempMove == move {
				count++
			} else { //if tempMove != move
				result = append(result, instructions(count, tempMove))
				tempMove = move
				count = 1
			}
		}
		result = append(result, instructions(count, tempMove))
		return
	}
	return
}

func display(input []string) {
	for _, move := range input {
		fmt.Println(move)
	}
}

func Run() {
	var input string
	fmt.Scan(&input)
	movement := calculateMove(input)
	if movement == nil {
		fmt.Println("Input is Invalid")
		return
	}
	display(movement)
}
