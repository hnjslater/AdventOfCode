package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var input_file = flag.String("input", "input.txt", "Input file")
var debug = flag.Bool("debug", false, "Debug")

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	board := make([]byte, len(line))

	for i, r := range line {
		board[i] = byte(r) - '0'
	}

	cur := 0

	for move := 1; move <= 100; move++ {
		fmt.Printf("-- move (%v) --\n", move)
		fmt.Printf("cups: ")
		for i, v := range board {
			if i == cur {
				fmt.Printf("(%v)", v)
			} else {
				fmt.Printf(" %v ", v)
			}
		}
		fmt.Println()

		pickup := make([]byte, 3)
		for i := 0; i < 3; i++ {
			pickup[i] = board[(cur+1+i)%len(board)]
		}

		fmt.Printf("pick up:")
		for i, v := range pickup {
			fmt.Printf(" %v", v)
			if i != len(pickup)-1 {
				fmt.Print(",")
			}
		}
		fmt.Println()

		dest_label := board[cur] - 1
		if dest_label == 0 {
			dest_label = 9
		}
		for dest_label == pickup[0] || dest_label == pickup[1] || dest_label == pickup[2] {
			dest_label -= 1
			if dest_label == 0 {
				dest_label = 9
			}
		}

		fmt.Printf("destination: %v\n", dest_label)
		fmt.Println()

		board2 := make([]byte, 0, len(board))

		for _, v := range board {
			if v == pickup[0] || v == pickup[1] || v == pickup[2] {
				// do nothing
			} else if v == dest_label {
				board2 = append(board2, v)
				board2 = append(board2, pickup[0])
				board2 = append(board2, pickup[1])
				board2 = append(board2, pickup[2])
			} else {
				board2 = append(board2, v)
			}
		}
		cur_label := board[cur]
		board = board2

		for i, v := range board {
			if v == cur_label {
				cur = i + 1
			}
		}
		if cur == len(board) {
			cur = 0
		}
	}
	print := false
	for _, v := range board {
		if print {
			fmt.Print(v)
		}
		if v == 1 {
			print = true
		}
	}
	for _, v := range board {
		if v == 1 {
			break
		}
		fmt.Print(v)
	}
	fmt.Println()
}
