package hw5

import (
	"fmt"

	tm "github.com/buger/goterm"
)

var field [3][3]int

func TicTacToe() {
	side := 1
	turn := 0
	for checkField() && turn < 9 {
		tm.Clear()
		tm.Println("Use NumPad to place your symbol")
		tm.Println("X is first")
		outputField()
		tm.Flush()
		if placeValue(side) {
			side *= -1
			turn += 1
		}
	}
	tm.Clear()
	tm.Flush()
	if side < 0 {
		fmt.Println("X is a winner!")
	} else {
		fmt.Println("O is a winner!")
	}
	if turn == 9 {
		fmt.Println("Its a draw!")
	}
}

func outputField() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tm.MoveCursor(i+3, j+2)
			if field[i][j] == 1 {
				tm.Print("X")
			}
			if field[i][j] == -1 {
				tm.Print("O")
			}
			if field[i][j] == 0 {
				tm.Print(" ")
			}
		}
	}
	tm.Println("")
}

func placeValue(value int) bool {
	var input string
	fmt.Scanf("%s\n", &input)
	switch input {
	case "7":
		if field[0][0] == 0 {
			field[0][0] = value
			return true
		}
	case "8":
		if field[1][0] == 0 {
			field[1][0] = value
			return true
		}
	case "9":
		if field[2][0] == 0 {
			field[2][0] = value
			return true
		}
	case "4":
		if field[0][1] == 0 {
			field[0][1] = value
			return true
		}
	case "5":
		if field[1][1] == 0 {
			field[1][1] = value
			return true
		}
	case "6":
		if field[2][1] == 0 {
			field[2][1] = value
			return true
		}
	case "1":
		if field[0][2] == 0 {
			field[0][2] = value
			return true
		}
	case "2":
		if field[1][2] == 0 {
			field[1][2] = value
			return true
		}
	case "3":
		if field[2][2] == 0 {
			field[2][2] = value
			return true
		}
	default:
		return false
	}
	return false
}

func checkDiagonals() bool {
	if field[0][0] == field[1][1] && field[1][1] == field[2][2] && field[2][2] != 0 {
		return false
	}
	if field[2][0] == field[1][1] && field[1][1] == field[0][2] && field[0][2] != 0 {
		return false
	}
	return true
}
func checkSides() bool {
	for i := 0; i < 3; i++ {
		if field[0][i] == field[1][i] && field[1][i] == field[2][i] && field[2][i] != 0 {
			return false
		}
		if field[i][0] == field[i][1] && field[i][1] == field[i][2] && field[i][2] != 0 {
			return false
		}
	}

	return true
}

func checkField() bool {
	return checkDiagonals() && checkSides()
}
