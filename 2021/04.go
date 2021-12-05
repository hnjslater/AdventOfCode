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

var input = flag.String("input", "input04.txt", "")
var part = flag.Int("part", 1, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	var draws []int
	for _, draw := range strings.Split(line, ",") {
		draw_parsed, _ := strconv.Atoi(draw)
		draws = append(draws, draw_parsed)
	}

	cards := make(map[int][10]map[int]bool)
	card_id := -1
	row_id := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			card_id++
			row_id = 0
			var new_card [10]map[int]bool
			for i := 0; i < 10; i++ {
				new_card[i] = make(map[int]bool)
			}
			cards[card_id] = new_card
		} else {
			for col_id, numstr := range strings.Fields(line) {
				num, _ := strconv.Atoi(numstr)
				cards[card_id][row_id][num] = true
				cards[card_id][col_id+5][num] = true
			}
			row_id++
		}
	}
bingo:
	for _, draw := range draws {
		for c := range cards {
			winner := false
			for _, g := range cards[c] {
				delete(g, draw)
				if len(g) == 0 {
					winner = true
				}
			}
			if winner {
				if *part == 1 || (*part == 2 && len(cards) == 1) {
					result := 0
					for _, g := range cards[c] {
						for n := range g {
							result += n
						}
					}
					fmt.Println((result * draw) / 2)
					break bingo
				} else {
					delete(cards, c)
				}
			}
		}
	}
}
