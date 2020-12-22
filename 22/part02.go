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

func hash(a [][]int) string {
	return fmt.Sprint(a)
}

func my_copy(hands [][]int, p0c int, p1c int) [][]int {
	hands2 := make([][]int, 2)
	hands2[0] = make([]int, p0c)
	hands2[1] = make([]int, p1c)
	copy(hands2[0], hands[0])
	copy(hands2[1], hands[1])

	return hands2
}

func play(game int, hands [][]int) (int, int) {
	history := make(map[string]bool)

	i := 1
	subgame := game
	for len(hands[0]) > 0 && len(hands[1]) > 0 {
		if history[hash(hands)] {
			return 0, subgame
		}
		history[hash(hands)] = true

		if *debug {
			fmt.Printf("-- Round %v (Game %v) --\n", i, game)
			fmt.Printf("Player 1's deck: %v\n", hands[0])
			fmt.Printf("Player 2's deck: %v\n", hands[1])
			fmt.Printf("Player 1 plays: %v\n", hands[0][0])
			fmt.Printf("Player 2 plays: %v\n", hands[1][0])
		}

		played := []int{hands[0][0], hands[1][0]}
		hands[0] = hands[0][1:len(hands[0])]
		hands[1] = hands[1][1:len(hands[1])]

		winner := 0
		loser := 0
		if len(hands[0]) >= played[0] && len(hands[1]) >= played[1] {
			if *debug {
				fmt.Println("NEW GAME")
			}
			subgame++
			winner, subgame = play(game, my_copy(hands, played[0], played[1]))
		} else if played[1] > played[0] {
			winner = 1
		}
		if winner == 0 {
			loser = 1
		}

		if *debug {
			fmt.Printf("Player %v wins ther round!\n\n", winner+1)
		}

		hands[winner] = append(hands[winner], played[winner])
		hands[winner] = append(hands[winner], played[loser])

		i++

	}

	if len(hands[0]) == 0 {
		if *debug {
			fmt.Printf("The Winner of Game %v is Player 2!\n", game)
		}
		return 1, subgame
	} else {
		if *debug {
			fmt.Printf("The Winner of Game %v is Player 1!\n", game)
		}
		return 0, subgame
	}
}

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
	play(1, hands)

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
