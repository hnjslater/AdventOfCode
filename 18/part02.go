package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
        "unicode"
)

var input_file = flag.String("input", "input.txt", "Input file")
var debug = flag.Bool("debug", false, "Debug")

func Push(s []int, r int) []int {
    return append(s, r)
}

func Pop (s []int) []int {
    return s[:len(s)-1]
}

func Peep(s []int) int {
    return s[len(s)-1]
}

func Prec(b rune) int {
    if b == '*' {
        return 0
    } else {
        return 1
    }
}

func Eval(line string) int {
    var exp []int
    var stack []int

    for _,r := range line {
        if unicode.IsDigit(r) {
            exp = append(exp, int(r))
        } else if unicode.IsSpace(r) {
        } else if r == ')' {
            inner: for len(stack) > 0 {
                v := Peep(stack)
                stack = Pop(stack)
                if v == '(' {
                    break inner
                }
                exp = append(exp, v)
            }
        } else if r == '(' {
            stack = Push(stack, int(r))
        } else {
            for len(stack) > 0 && Peep(stack) != '(' && Prec(r) <= Prec(rune(Peep(stack))) {
                exp = append(exp, Peep(stack))
                stack = Pop(stack)
            }
            stack = Push(stack, int(r))
        }
    }

    for len(stack) > 0 {
        exp = append(exp, Peep(stack))
        stack = Pop(stack)
    }

    for _,r := range exp {
        if unicode.IsDigit(rune(r)) {
            stack = Push(stack, r-'0')
        } else if r == '+' || r == '*' {
            operand1 := Peep(stack)
            stack = Pop(stack)
            operand2 := Peep(stack)
            stack = Pop(stack)
            switch (r) {
                case '+':
                    stack = Push(stack, operand1 + operand2)
                case '*':
                    stack = Push(stack, operand1 * operand2)
            }
        }
    }
    return Peep(stack)
}

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
        total := 0
	for scanner.Scan() {
		line := scanner.Text()
                t := Eval(line)
                total += t
	}
        fmt.Println(total)
}
