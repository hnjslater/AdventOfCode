package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var input = flag.String("input", "input03.txt", "")
var part = flag.Int("part", 0, "")

func priority(b byte) int {
	if b > byte('`') {
		return int(b) - 96
	} else {
		return int(b) - 38
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
	var rucksacks []string
	for scanner.Scan() {
		line := scanner.Text()
		rucksacks = append(rucksacks, line)
	}
	if *part == 0 {
		for _, line := range rucksacks {
			var comps [2]string
			comps[0] = line[0 : len(line)/2]
			comps[1] = line[len(line)/2 : len(line)]

			var sets [2]map[byte]bool
			sets[0] = make(map[byte]bool)
			sets[1] = make(map[byte]bool)
			for i := 0; i < 2; i++ {
				for j := 0; j < len(comps[i]); j++ {
					sets[i][comps[i][j]] = true
				}
			}
			for k, _ := range sets[0] {
				if sets[1][k] {
					total += priority(k)
				}

			}
		}
	} else {
		for i := 0; i < len(rucksacks)/3; i++ {
			candidates := make(map[byte]int)

			for j := 0; j < 3; j++ {
				r := rucksacks[(i*3)+j]

				for _, b := range r {
					if candidates[byte(b)] == j {
						candidates[byte(b)] += 1
					}
				}
			}

			for k, v := range candidates {
				if v == 3 {
					total += priority(k)
				}
			}
		}
	}
	fmt.Println(total)
}
