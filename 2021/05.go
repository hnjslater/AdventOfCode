package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var input = flag.String("input", "", "")
var part = flag.Int("part", 0, "")
var lineExp = regexp.MustCompile(`(?P<x1>\d+),(?P<y1>\d+) -> (?P<x2>\d+),(?P<y2>\d+)`)

type place struct {
	x int
	y int
}

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	world := make(map[place]int)
	for scanner.Scan() {
		line := scanner.Text()
		match := lineExp.FindStringSubmatch(line)
		x1, _ := strconv.Atoi(match[1])
		y1, _ := strconv.Atoi(match[2])
		x2, _ := strconv.Atoi(match[3])
		y2, _ := strconv.Atoi(match[4])

		dx := 0
		dy := 0
		if x1 < x2 {
			dx = 1
		} else if x1 > x2 {
			dx = -1
		}

		if y1 < y2 {
			dy = 1
		} else if y1 > y2 {
			dy = -1
		}

		x := x1
		y := y1
		if *part == 1 && dx != 0 && dy != 0 {
			continue
		}
		for {
			world[place{x, y}] += 1
			if x == x2 && y == y2 {
				break
			}
			x += dx
			y += dy

		}
	}
	for _, p := range world {
		if p >= 2 {
			result++
		}
	}
	fmt.Println(result)
}
