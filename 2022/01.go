package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var input = flag.String("input", "01.txt", "")
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

	var elves []int
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			value, _ := strconv.Atoi(line)
			total += value
		} else {
			elves = append(elves, total)
			total = 0
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	if *part == 1 {
		fmt.Println(elves[0])
	} else {
		fmt.Println(elves[0] + elves[1] + elves[2])
	}
}
