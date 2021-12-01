package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	hands := make([][]int, 2)
	player := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Player ") {
			player, err = strconv.Atoi(line[7:8])
			if err != nil {
				log.Fatal(err)
			}
		} else if line == "" {
			// do nothing
		} else {
			x, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			hands[player-1] = append(hands[player-1], x)
		}
	}
	i := 1
	for len(hands[0]) > 0 && len(hands[1]) > 0 {
		fmt.Printf("-- Round %v --\n", i)
		fmt.Printf("Player 1's deck: %v\n", hands[0])
		fmt.Printf("Player 2's deck: %v\n", hands[1])
		fmt.Printf("Player 1 plays: %v\n", hands[0][0])
		fmt.Printf("Player 2 plays: %v\n", hands[1][0])

		winner := 0
		loser := 1
		if hands[0][0] > hands[1][0] {
			winner = 1
			loser = 0
		}
		fmt.Printf("Player %v wins ther round!\n\n", winner+1)
		winning_card := hands[winner][0]
		losing_card := hands[loser][0]

		hands[loser] = hands[loser][1:len(hands[loser])]
		hands[winner] = hands[winner][1:len(hands[winner])]

		hands[loser] = append(hands[loser], losing_card)
		hands[loser] = append(hands[loser], winning_card)

		i++
	}
	total := 0

	loser := 1
	if len(hands[0]) > len(hands[1]) {
		loser = 0
	}

	for i, x := range hands[loser] {
		total += (len(hands[loser]) - i) * x
	}

	fmt.Println(total)
}
