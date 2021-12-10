package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

var input = flag.String("input", "input10.txt", "")
var part = flag.Int("part", 1, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	var scores = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	var scores2 = map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
	var closer = map[rune]rune{'(': ')', '{': '}', '<': '>', '[': ']'}
	var lineScores []int

	for scanner.Scan() {
		line := scanner.Text()
		var stack []rune
	loop:
		for _, r := range line {
			switch r {
			case '(', '[', '{', '<':
				stack = append(stack, r)
			case ')', ']', '}', '>':
				if r == closer[stack[len(stack)-1]] {
					stack = stack[0 : len(stack)-1]
				} else if *part == 1 {
					result += scores[r]
					break loop
				} else {
					stack = []rune{}
					break loop
				}
			}
		}
		if *part == 2 && len(stack) > 0 {
			score := 0
			for i := len(stack) - 1; i >= 0; i-- {
				score = (score * 5) + scores2[closer[stack[i]]]
			}

			lineScores = append(lineScores, score)
		}
	}
	if *part == 2 {
		sort.IntSlice.Sort(lineScores)
		result = lineScores[len(lineScores)/2]
	}

	fmt.Println(result)
}
