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

func flipHorizontally(b [7][7]byte) [7][7]byte {
	board := [7][7]byte{}
	for y := range b {
		for x := range b {
			board[y][x] = b[y][6-x]
		}
	}
	return board
}

func flipVertically(b [7][7]byte) [7][7]byte {
	board := [7][7]byte{}
	for y := range b {
		board[y] = b[6-y]
	}
	return board
}

func rotations(board [7][7]byte) [4][7][7]byte {
	r90 := rotate90(board)
	r180 := rotate90(r90)
	r270 := rotate90(r180)
	return [4][7][7]byte{
		board,
		r90,
		r180,
		r270,
	}
}

func Check(uniques map[[7][7]byte]bool, board [7][7]byte) bool {
	fh := flipHorizontally(board)
	fv := flipVertically(board)
	fvh := flipVertically(fh)

	options := [4][4][7][7]byte{
		rotations(board),
		rotations(fh),
		rotations(fv),
		rotations(fvh),
	}

	for _, arr := range options {
		for _, val := range arr {
			if uniques[val] {
				return false
			}
		}
	}

	return true
}
