package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func run(dx int, dy int) int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	x := 0
	width := 31
	trees := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line[x] == '#' {
			trees++
		}
		x = (x + dx) % width
		for i := 1; i < dy; i++ {
			scanner.Scan()
			fmt.Println("Skipped")
		}
	}
	return trees
}

func main() {
	fmt.Println(run(1, 1) * run(3, 1) * run(5, 1) * run(7, 1) * run(1, 2))
}
