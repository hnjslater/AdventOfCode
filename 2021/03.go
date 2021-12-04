package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var input = flag.String("input", "", "")
var part = flag.Int("part", 0, "")

func Part1() int64 {
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	width := int64(-1)
	count := 0
	var counts []int
	for scanner.Scan() {
		line := scanner.Text()
		if width == -1 {
			width = int64(len(line))
			counts = make([]int, width)
		}
		for i, c := range line {
			if c == '1' {
				counts[i]++
			}
		}
		count++
	}
	var gamma int64 = 0
	for i := int64(0); i < width; i++ {
		if counts[i] > (count / 2) {
			gamma = (gamma << 1) | 1

		} else {
			gamma = (gamma << 1)
		}
	}

	epsilon := gamma ^ ((1 << width) - 1)
	return gamma * epsilon
}
func Part2(popular bool) int64 {
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	width := int64(-1)
	scanner := bufio.NewScanner(file)
	values := make(map[int64]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if width == -1 {
			width = int64(len(line))
		}
		value, _ := strconv.ParseInt(line, 2, 64)
		values[value] = true
	}

	for i := int64(0); i < width; i++ {
		count := 0
		test := int64(1 << (width - i - 1))
		for v := range values {
			if v&test > 0 {
				count++
			}
		}
		crit := (count * 2) >= len(values)
		if !popular {
			crit = !crit
		}
		for v := range values {
			if crit {
				if v&test == 0 {
					delete(values, v)
				}
			} else if v&test > 0 {
				delete(values, v)
			}
		}

		if len(values) == 1 {
			break
		}
	}
	for v := range values {
		return v
	}
	return -1
}

func main() {
	flag.Parse()
	if *part == 1 {
		fmt.Println(Part1())
	} else {
		fmt.Println(Part2(true) * Part2(false))
	}
}
