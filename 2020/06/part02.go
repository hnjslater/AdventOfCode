package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	group_answers := make(map[rune]bool)
	answers := make(map[rune]bool)
	for r := 'a'; r <= 'z'; r++ {
		group_answers[r] = true
	}
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		answers = make(map[rune]bool)
		if line != "" {
			for _, r := range line {
				answers[r] = true
			}
			for r := 'a'; r <= 'z'; r++ {
				if !answers[r] {
					delete(group_answers, r)
				}
			}
		} else {
			total += len(group_answers)
			for r := 'a'; r <= 'z'; r++ {
				group_answers[r] = true
			}
		}
	}
	total += len(group_answers)
	fmt.Println(total)

}
