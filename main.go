package main

import (
	"encoding/json"
	"fmt"
	getmoves "main/get-moves"
	startingboard "main/starting-board"
	"main/unique"
	"os"
	"time"
)

func main() {
	start()
	// board := startingboard.Startingboard

	// board[2][2] = 1

	// printboard.PrintBoard(board)

	// ro90 := comparitors.Rotate90(board)
	// ro180 := comparitors.Rotate90(ro90)
	// ro270 := comparitors.Rotate90(ro180)

	// printboard.PrintBoard(ro90)
	// printboard.PrintBoard(ro180)
	// printboard.PrintBoard(ro270)
}

func start() {
	data := [][][7][7]byte{}
	main := make(chan [][7][7]byte, 31)
	now := time.Now()
	total := time.Now()

	go computeLayer([][7][7]byte{startingboard.Startingboard}, main)

	for cur := range main {
		fmt.Println(len(data)+1, "layers computed items:", len(cur), "time", time.Since(now))
		now = time.Now()
		data = append(data, cur)

		if len(data) == 31 {
			close(main)
		}
	}

	fmt.Println("done", len(data)+1, "layers computed, time taken:", time.Since(total))

	file, err := json.MarshalIndent(data, "", " ")

	fmt.Println(err)

	fmt.Println(file)

	os.WriteFile("data.json", file, 0644)
}

func computeLayer(boards [][7][7]byte, main chan<- [][7][7]byte) {
	jobs := make(chan [7][7]byte, len(boards))
	results := make(chan [][7][7]byte, len(boards))

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	go func() {
		for _, board := range boards {
			jobs <- board
		}
		close(jobs)
	}()

	uniques := map[[7][7]byte]bool{}
	moves := [][7][7]byte{}
	for i := 0; i < len(boards); i++ {
		res := <-results
		for _, board := range res {
			if unique.Check(uniques, board) {
				uniques[board] = true
				moves = append(moves, board)
			}
		}
	}
	close(results)
	if len(moves) == 0 {
		return
	}

	main <- moves

	go computeLayer(moves, main)
}

func worker(jobs <-chan [7][7]byte, results chan<- [][7][7]byte) {
	for job := range jobs {
		results <- getmoves.Getmoves(job)
	}
}
