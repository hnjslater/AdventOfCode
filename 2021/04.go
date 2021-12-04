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

	var draws []int
	scanner.Scan()
	line := scanner.Text()

	for _, draw := range strings.Split(line, ",") {
		draw_parsed, _ := strconv.Atoi(draw)
		draws = append(draws, draw_parsed)
	}
	var cards [][][]int
	card_id := 0
	cards = append(cards, [][]int{})
	removed := make(map[int]bool)

	scanner.Scan() // skip blank line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			card_id++
			cards = append(cards, [][]int{})
			continue
		}
		numstrs := strings.Split(line, " ")
		var row []int
		for _, numstr := range numstrs {
			if numstr == "" {
				continue
			}
			num_parsed, _ := strconv.Atoi(numstr)
			row = append(row, num_parsed)
		}
		cards[card_id] = append(cards[card_id], row)
	}

bingo:
	for _, draw := range draws {
		for c := range cards {
			for r := range cards[c] {
				for e := range cards[c][r] {
					if cards[c][r][e] == draw {
						cards[c][r][e] = -1
					}
				}
			}
		}

		for c := range cards {
			if removed[c] {
				continue
			}
			row_match := true
			for r := range cards[c] {
				row_match = true
				for e := range cards[c][r] {
					if cards[c][r][e] != -1 {
						row_match = false
					}
				}
				if row_match {
					break
				}
			}

			col_match := true
			for e := 0; e < 5; e++ {
				col_match = true
				for r := 0; r < 5; r++ {
					if cards[c][r][e] != -1 {
						col_match = false
					}
				}
				if col_match {
					break
				}
			}

			if col_match || row_match {

				result := 0
				if row_match || col_match {
					for _, r := range cards[c] {
						for _, e := range r {
							if e > -1 {
								result += e
							}
						}
					}
				}

				result *= draw
				if *part == 1 {
					fmt.Println(result)
					break bingo
				} else if len(removed)+1 == len(cards) {
					fmt.Println(result)
					break bingo
				} else {
					removed[c] = true
				}
			}
		}
	}
}
