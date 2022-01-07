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
}

func fileboard(board [7][7]byte) string {
	str := ""
	for y := range board {
		for _, val := range board[y] {
			str += fmt.Sprint(val)
		}
	}
	str += ","
	return str
}

func start() {
	main := make(chan [][7][7]byte, 31)
	now := time.Now()

	go computeLayer([][7][7]byte{startingboard.Startingboard}, main)

	for i := 0; i < 31; i++ {
		cur := <-main
		now = time.Now()

		str := ""

		for _, board := range cur {
			str += fmt.Sprint(board) + ","
		}

		str += ";"

		file, _ := json.Marshal(cur)
		os.WriteFile("data/"+fmt.Sprint(i)+".json", file, 0644)

		fmt.Println(i+1, "layers computed items:", len(cur), "time", time.Since(now))
	}

	close(main)
}

func computeLayer(boards [][7][7]byte, main chan<- [][7][7]byte) {
	jobs := make(chan [7][7]byte, len(boards))
	results := make(chan [][7][7]byte, len(boards))

	go worker(jobs, results)
	go worker(jobs, results)
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
