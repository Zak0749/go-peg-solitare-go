package unique

func rotate90(arr [7][7]byte) [7][7]byte {
	board := arr
	for y := range arr {
		for x := range arr[y] {
			temp := arr[y][x]
			board[y][x] = arr[7-1-x][y]
			board[7-1-x][y] = arr[7-1-y][7-1-x]
			board[7-1-y][7-1-x] = arr[x][7-1-y]
			board[x][7-1-y] = temp
		}
	}
	return board
}

func Check(uniques map[[7][7]byte]bool, board [7][7]byte) bool {
	r90 := rotate90(board)
	r180 := rotate90(r90)
	r270 := rotate90(r180)

	return !(uniques[board] || uniques[r90] || uniques[r180] || uniques[r270])
}
