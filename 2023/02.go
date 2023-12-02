package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var input = flag.String("input", "02.txt", "")
var part = flag.Int("part", 0, "")

var re = regexp.MustCompile("(?P<count>[0-9]+) (?P<color>[a-x]+)")

var MAX_CUBES = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	game := 1
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		possible := true
		maxes := make(map[string]int)
		for _, turn := range strings.Split(line[1], ";") {
			matches := re.FindAllStringSubmatch(turn, -1)
			turn := make(map[string]int)
			for _, m := range matches {
				turn[m[2]], _ = strconv.Atoi(m[1])
			}
			for c := range MAX_CUBES {
				possible = possible && turn[c] <= MAX_CUBES[c]
				maxes[c] = max(maxes[c], turn[c])
			}

		}
		if *part == 0 {
			if possible {
				total += game
			}
		} else {
			subtotal := 1
			for c := range MAX_CUBES {
				subtotal *= maxes[c]
			}
			total += subtotal
		}
		game++
	}

	fmt.Println(total)
}
