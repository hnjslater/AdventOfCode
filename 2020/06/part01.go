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
	answers := make(map[rune]bool)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			for _, r := range line {
				answers[r] = true
			}
		} else {
			total += len(answers)
			answers = make(map[rune]bool)
		}
	}
	total += len(answers)
	fmt.Println(total)

}
