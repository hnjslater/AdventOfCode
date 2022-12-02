package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var input = flag.String("input", "input02.txt", "")
var part = flag.Int("part", 0, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var scores = map[byte]int{'R': 1, 'P': 2, 'S': 3}
	var elf_moves = map[byte]byte{'A': 'R', 'B': 'P', 'C': 'S'}
	var my_moves = map[byte]byte{'X': 'R', 'Y': 'P', 'Z': 'S'}
	var winner = map[byte]byte{'R': 'P', 'P': 'S', 'S': 'R'}

	loser := make(map[byte]byte)
	for k, v := range winner {
		loser[v] = k
	}

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		elf_move := elf_moves[line[0]]
		my_move := line[2]

		if *part == 1 {
			if my_move == 'X' {
				my_move = loser[elf_move]
			} else if my_move == 'Y' {
				my_move = elf_move
			} else {
				my_move = winner[elf_move]
			}
		} else {
			my_move = my_moves[my_move]
		}

		if winner[elf_move] == my_move {
			total += 6
		}
		if elf_move == my_move {
			total += 3
		}
		total += scores[my_move]

		fmt.Println(string(elf_move), string(my_move), total)

	}
	fmt.Println(total)
}
