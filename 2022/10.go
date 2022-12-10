package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var input = flag.String("input", "input10.txt", "")
var part = flag.Int("part", 0, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	result := 0
	hangover := false
	reg := 1
	operand := 0

	importantCycles := map[int]bool{20: true, 60: true, 100: true, 140: true, 180: true, 220: true}
	for cycle := 1; cycle <= 240; cycle++ {
		if *part != 0 {
			if Abs(((cycle-1)%40)-reg) < 2 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			if cycle%40 == 0 {
				fmt.Println()
			}
		}
		if importantCycles[cycle] {
			result += (cycle * reg)
		}
		if hangover {
			reg += operand
			hangover = false
		} else {
			if !scanner.Scan() {
				continue
			}
			line := scanner.Text()
			op := line[0:4]
			if op == "addx" {
				operand, _ = strconv.Atoi(line[5:])
				hangover = true
			}
		}

	}
	if *part == 0 {
		fmt.Println(result)
	}
}
