package getmoves

func opp(num byte) byte {
	if num == 1 {
		return 2
	}
	return 1
}

func Getmoves(board [7][7]byte) [][7][7]byte {
	moves := [][7][7]byte{}
	for y, row := range board {
		for x, val := range row {
			if val != 2 {
				continue
			}

			if y != 0 && y != 6 && board[y+1][x] != board[y-1][x] && board[y+1][x] != 0 && board[y-1][x] != 0 {
				copy := board
				copy[y][x] = 1
				copy[y+1][x] = opp(copy[y+1][x])
				copy[y-1][x] = opp(copy[y-1][x])
				moves = append(moves, copy)
			}

			if x != 0 && x != 6 && board[y][x+1] != board[y][x-1] && board[y][x+1] != 0 && board[y][x-1] != 0 {
				copy := board
				copy[y][x] = 1
				copy[y][x+1] = opp(copy[y][x+1])
				copy[y][x-1] = opp(copy[y][x-1])
				moves = append(moves, copy)
			}
		}
	}
	return moves
}
