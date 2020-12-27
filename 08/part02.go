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
	complete := false
	for i := 0; i < len(program); i++ {

		var modified_program = make([]instruction, len(program))
		copy(modified_program, program)
		switch program[i].operation {
		case ("jmp"):
			modified_program[i].operation = "nop"
		case ("nop"):
			modified_program[i].operation = "jmp"
		case ("acc"):
			continue
		}
		acc = 0
		pc = 0
		for !modified_program[pc].run_before {
			modified_program[pc].run_before = true
			switch modified_program[pc].operation {
			case ("acc"):
				acc += modified_program[pc].operand
				pc++
			case ("jmp"):
				pc += modified_program[pc].operand
			case ("nop"):
				pc++
			}
			if pc >= len(program) {
				complete = true
				break
			}
		}
		if complete {
			break
		}
	}

	fmt.Println(acc)
}
