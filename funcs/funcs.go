package funcs

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func GetSide() rune {
	return []rune{'X', 'O'}[rand.Intn(1+1)]
}

func GenerateBoard() [][]string {
	return [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
}

func DisplayBoard(board [][]string) {
	fmt.Printf("\n")
	for i := 0; i < len(board); i++ {
		fmt.Printf("  %s\n", strings.Join(board[i], "  "))
	}
	fmt.Printf("\n")
}

func InterpretMoves(cmd string) (int8, int8, error) {
	command := strings.Split(cmd, " ") // Slice

	/* var possibleMoves = map[string]int8{
		"keski":    1,
		"ylä":      0,
		"ala":      2,
		"vasen":    0,
		"keskellä": 1,
		"oikea":    2,
	} */
	var x, y int8
	for _, v := range command {
		switch v {
		case "ylä":
			y = 0
		case "keski":
			y = 1
		case "ala":
			y = 2
		case "vasen":
			x = 0
		case "keskellä":
			x = 1
		case "oikea":
			x = 2
		default:
			x = -1
			y = -1
		}
	}
	if x > 2 || y > 2 {
		return -1, -1, errors.New("Faulty command for location!")
	} else if x == -1 && y == -1 {
		return -1, -1, errors.New("Faulty command for location!")
	} else {
		return x, y, nil
	}
}

func AssignMoves(board *[][]string, side string, x int8, y int8) error {
	if (*board)[y][x] != "_" {
		return errors.New("You can't place over existing letter!")
	} else {
		(*board)[y][x] = side
		return nil
	}
}

func ComputerMove(board *[][]string, side string) {
	for i := 0; i < 9; i++ {
		randX, randY := rand.Intn(3), rand.Intn(3)
		if (*board)[randY][randX] == "_" { // Jos indeksissä ei ole mitään, aseta symboli sinne
			(*board)[randY][randX] = side
			break
		} else {
			continue
		}
	}
}

func HasWon(board *[][]string, side string) bool {
	// Horizontal
	if (*board)[0][0] == side && (*board)[0][1] == side && (*board)[0][2] == side {
		return true
	}
	if (*board)[1][0] == side && (*board)[1][1] == side && (*board)[1][2] == side {
		return true
	}
	if (*board)[2][0] == side && (*board)[2][1] == side && (*board)[2][2] == side {
		return true
	}

	// Vertical
	if (*board)[0][0] == side && (*board)[1][0] == side && (*board)[2][0] == side {
		return true
	}
	if (*board)[0][1] == side && (*board)[1][1] == side && (*board)[2][1] == side {
		return true
	}
	if (*board)[0][2] == side && (*board)[1][2] == side && (*board)[2][2] == side {
		return true
	}

	// Diagonals
	if (*board)[0][0] == side && (*board)[1][1] == side && (*board)[2][2] == side {
		return true
	}
	if (*board)[2][0] == side && (*board)[1][1] == side && (*board)[0][2] == side {
		return true
	}
	return false
}

func IsDraw(board *[][]string) bool {
	var count int
	for _, v := range *board {
		for _, v := range v {
			if v != "_" {
				count++
			}
		}
	}

	if count == 9 { // 9 is the number of tiles in game
		return true
	}
	return false
}
