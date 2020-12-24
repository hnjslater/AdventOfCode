package main

import (
	"bufio"
	"container/ring"
	"flag"
	"fmt"
	"log"
	"os"
)

var input_file = flag.String("input", "input.txt", "Input file")
var debug = flag.Bool("debug", false, "Debug")
var part2 = flag.Bool("part2", false, "Part2?")

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
	cache := make(map[int]*ring.Ring)

	var max_value int
	if !*part2 {
		max_value = 9
	} else {
		max_value = 1_000_000
	}
	board := ring.New(max_value)

	for _, r := range line {
		v := int(byte(r) - '0')
		board.Value = v
		cache[v] = board
		board = board.Next()
	}
	if *part2 {
		for i := 10; i <= max_value; i++ {
			board.Value = i
			cache[i] = board
			board = board.Next()
		}
	}

	final_move := 100
	if *part2 {
		final_move = 10_000_000
	}

	for move := 1; move <= final_move; move++ {
		if *debug {
			fmt.Printf("-- move (%v) --\n", move)
			fmt.Printf("cups: ")
			for e := board; ; e = e.Next() {
				if e == board {
					fmt.Printf("(%v)", e.Value)
				} else {
					fmt.Printf(" %v ", e.Value)
				}
				if e.Next() == board {
					break
				}
			}
			fmt.Println()
		}

		pickup := board.Unlink(3)

		if *debug {
			fmt.Printf("pick up:")
			for e := pickup; ; e = e.Next() {
				fmt.Printf(" %v", e.Value)
				if e.Next() != pickup {
					fmt.Print(",")
				} else {
					break
				}
			}
			fmt.Println()
		}

		dest_label := board.Value.(int) - 1
		if dest_label == 0 {
			dest_label = max_value
		}
		for dest_label == pickup.Value || dest_label == pickup.Next().Value || dest_label == pickup.Next().Next().Value {
			dest_label -= 1
			if dest_label == 0 {
				dest_label = max_value
			}
		}

		if *debug {
			fmt.Printf("destination: %v\n", dest_label)
			fmt.Println()
		}

		cache[dest_label].Link(pickup)

		board = board.Next()

	}
	if !*part2 {
		print := false
		for e := board; ; e = e.Next() {
			if print {
				fmt.Print(e.Value)
			}
			if e.Value == 1 {
				print = true
			}
			if e.Next() == board {
				break
			}
		}
		for e := board; e.Next() != board; e = e.Next() {
			if e.Value == 1 {
				break
			}
			fmt.Print(e.Value)
		}
		fmt.Println()
	} else {
		fmt.Println(cache[1].Next().Value.(int) * cache[1].Next().Next().Value.(int))
	}
}
