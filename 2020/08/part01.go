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

type instruction struct {
	operation  string
	operand    int
	run_before bool
}

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var program []instruction

	for scanner.Scan() {
		line := scanner.Text()
		split_line := strings.Split(line, " ")
		opr := split_line[0]
		opa, err := strconv.Atoi(split_line[1])
		if err != nil {
			panic("Error parsing input")
		}

		program = append(program, instruction{opr, opa, false})
	}

	pc := 0
	acc := 0

	for !program[pc].run_before {
		program[pc].run_before = true
		switch program[pc].operation {
		case ("acc"):
			acc += program[pc].operand
			pc++
		case ("jmp"):
			pc += program[pc].operand
		case ("nop"):
			pc++
		}
	}

	fmt.Print(acc)
}
