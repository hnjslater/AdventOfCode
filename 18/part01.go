package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var input_file = flag.String("input", "input.txt", "Input file")
var debug = flag.Bool("debug", false, "Debug")

func Eval(s string) (int, int) {
    acc := 0
    var op byte = '?'
    var i int
    for i=0; i < len(s); i++ {
        r := s[i]
        switch (r) {
            case ' ':
                // don't care
            case '+', '-', '*':
                op = r
            case ')':
                return acc, i
            default:
                val := 0
                if r == '(' {
                    var j int
                    val, j = Eval(s[i+1:])
                    i += j + 1
                } else if r >= '0' && r <= '9' {
                    val = int(r) - '0'
                } else {
                    log.Fatal("don't recognise ", rune(r))
                }
                if op == '?' {
                    acc = val
                } else if op == '+' {
                    acc += val
                } else if op == '-' {
                    acc -= val
                } else if op == '*' {
                    acc *= val
                } else {
                    log.Fatal("don't recognise ", op)
                }
        }
    }
    return acc, i
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
                t,_ := Eval(line)
                total += t
	}
        fmt.Println(total)
}
