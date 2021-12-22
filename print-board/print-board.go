package printboard

import "fmt"

func PrintBoard(board [7][7]byte) {
	str := ""

	for _, row := range board {
		for _, val := range row {
			if val == 0 {
				str += "  "
			} else if val == 2 {
				str += "● "
			} else {
				str += "○ "
			}
		}
		str += "\n"
	}
	fmt.Println(str)
}
