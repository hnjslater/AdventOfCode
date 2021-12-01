package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
	wx := 10
	wy := 1
	for scanner.Scan() {
		line := scanner.Text()
		operation := line[0]
		operand, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("Error parsing input.")
		}
		switch operation {
		case 'N':
			wy += operand
		case 'S':
			wy -= operand
		case 'E':
			wx += operand
		case 'W':
			wx -= operand
		case 'F':
			x += wx * operand
			y += wy * operand
		case 'L', 'R':
			if operation == 'L' {
				operand = 360 - operand
			}
			for i := 0; i < operand/90; i++ {
				wx, wy = wy, -wx
			}
		}
	}
	fmt.Println(Abs(x) + Abs(y))
}
