// Name	: Muhammad Ziad Rahmatullah
// Date	: Oct 10, 21:19

package main

import (
	"fmt"
	"regexp"
)

func instructions(count int, move string) string {
	if move == "R" {
		return fmt.Sprintf("Move right %d times", count)
	} else if move == "L" {
		return fmt.Sprintf("Move left %d times", count)
	} else { //if move == "A"
		return fmt.Sprintf("Move advance %d times", count)
	}
}

func calculateMove(moves string) (result []string, err error) {
	// Validation
	valid := regexp.MustCompile(`^[RLA]+$`)
	if !valid.MatchString(moves) {
		return nil, fmt.Errorf("Invalid input: %s", moves)
	}

	// Calculate
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
	if tempMove != "" {
		result = append(result, instructions(count, tempMove))
	}

	return result, nil
}

func display(input []string) {
	for _, move := range input {
		fmt.Println(move)
	}
}

func main() {
	var input string
	fmt.Scan(&input)
	movement, err := calculateMove(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	display(movement)
}
