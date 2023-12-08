package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var input = flag.String("input", "08.txt", "")
var part = flag.Int("part", 0, "")

type node struct {
	left  string
	right string
}

// borrowed from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	path := scanner.Text()
	scanner.Scan()

	network := make(map[string]node)

	for scanner.Scan() {
		line := scanner.Text()

		start := line[0:3]
		left := line[7:10]
		right := line[12:15]

		fmt.Println(start, left, right)
		network[start] = node{left, right}

	}

	var starts []string
	lengths := make(map[string]int)

	for k := range network {
		if k[2] == 'A' {
			starts = append(starts, k)
		}
	}

	for i, curr := range starts {
		for curr[2] != 'Z' {
			if path[lengths[starts[i]]%len(path)] == 'L' {
				curr = network[curr].left
			} else {
				curr = network[curr].right
			}
			lengths[starts[i]]++
		}
	}
	if *part == 0 {
		fmt.Println(lengths["AAA"])
	} else {
		var llengths []int
		for _, v := range lengths {
			llengths = append(llengths, v)
		}
		fmt.Println(LCM(llengths[0], llengths[1], llengths[2:]...))
	}
}
