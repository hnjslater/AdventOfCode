package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var input = flag.String("input", "07.txt", "")
var part = flag.Int("part", 0, "")

type hand struct {
	cards string
	bid   int
}

var rank = map[byte]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

const FIVE_KIND = 7
const FOUR_KIND = 6
const FULL_HOUS = 5
const TRIPS = 4
const TWO_PAIR = 3
const PAIR = 2
const HIGH = 1

func (h hand) strength() int {
	counts := make(map[byte]int)

	jokers := 0
	for i := range h.cards {
		if *part > 0 && h.cards[i] == 'J' {
			jokers++
			counts['J'] = 0
		} else {
			counts[h.cards[i]]++
		}
	}
	num_ranks := len(counts)
	max_same := 0

	for _, v := range counts {
		if v > max_same {
			max_same = v
		}
	}

	if (max_same + jokers) == 5 {
		return FIVE_KIND
	}
	if (max_same + jokers) == 4 {
		return FOUR_KIND
	}
	if max_same == 3 {
		if num_ranks == 2 {
			return FULL_HOUS
		}
		return TRIPS
	}
	if jokers == 2 {
		return TRIPS
	}
	if max_same == 2 {
		if num_ranks == 3 && jokers == 0 {
			return TWO_PAIR
		}
		if num_ranks == 3 && jokers == 1 {
			return FULL_HOUS
		}
		if jokers == 1 {
			return TRIPS
		}
		return PAIR
	}
	if jokers == 1 {
		return PAIR
	}

	return HIGH

}

func (h hand) value() int {
	value := h.strength()
	for i := range h.cards {
		value = (value * (len(rank) + 1)) + rank[h.cards[i]]
	}

	return value
}

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if *part > 0 {
		for k := range rank {
			rank[k]++
		}
		rank['J'] = 0
	}

	scanner := bufio.NewScanner(file)

	total := 0

	var hands []hand

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		h := line[0]
		b, _ := strconv.Atoi(line[1])

		hands = append(hands, hand{h, b})
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].value() < hands[j].value()
	})

	for i, v := range hands {
		total += (v.bid * (i + 1))
	}
	fmt.Println(total)
}
