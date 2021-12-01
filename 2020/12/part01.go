package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var input_file = flag.String("input", "input.txt", "Input file")

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	x := 0
	y := 0
	var d byte = 'E'
	compass := "NESW"
	for scanner.Scan() {
		line := scanner.Text()
		operation := line[0]
		operand, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("Error parsing input.")
		}
		if operation == 'F' {
			operation = d
		}
		switch operation {
		case 'N':
			y += operand
		case 'S':
			y -= operand
		case 'E':
			x += operand
		case 'W':
			x -= operand
		case 'L', 'R':
			operand = operand / 90
			if operation == 'L' {
				operand = -operand
			}
			index := (strings.IndexByte(compass, d) + operand) % 4
			if index < 0 {
				index += len(compass)
			}
			d = compass[index]
		}
	}
	fmt.Println(Abs(x) + Abs(y))
}
