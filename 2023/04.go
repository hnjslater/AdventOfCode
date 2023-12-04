package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	//	"unicode"
)

var input = flag.String("input", "04.txt", "")
var part = flag.Int("part", 0, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	copies := make(map[int]int)
	card := 0
	for scanner.Scan() {
		line := scanner.Text()
		winners := make(map[int]bool)
		card_total := 0

		winner_string := strings.Split(strings.Split(line, ":")[1], "|")[0]
		for _, num_str := range strings.Fields(winner_string) {
			num, _ := strconv.Atoi(num_str)
			winners[num] = true
		}

		entries_string := strings.Split(strings.Split(line, ":")[1], "|")[1]
		for _, num_str := range strings.Fields(entries_string) {
			num, _ := strconv.Atoi(num_str)
			if winners[num] {
				if *part == 0 {
					if card_total == 0 {
						card_total = 1
					} else {
						card_total = card_total * 2
					}
				} else {
					card_total++
				}
			}
		}
		if *part == 0 {
			total += card_total
		}
		copies[card]++
		for i := 0; i < card_total; i++ {
			copies[card+i+1] += copies[card]
		}
		card++
	}
	if *part != 0 {
		for _, v := range copies {
			total += v
		}
	}
	fmt.Println(total)
}
