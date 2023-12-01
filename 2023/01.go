package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var input = flag.String("input", "01.txt", "")
var part = flag.Int("part", 0, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	numbers := make(map[string]int)
	if *part > 0 {
		numbers["zero"] = 0
		numbers["one"] = 1
		numbers["two"] = 2
		numbers["three"] = 3
		numbers["four"] = 4
		numbers["five"] = 5
		numbers["six"] = 6
		numbers["seven"] = 7
		numbers["eight"] = 8
		numbers["nine"] = 9
		numbers["ten"] = 10
	}

	for scanner.Scan() {
		line := scanner.Text()
		first := -1
		last := 0

		for i := range line {
			n, e := strconv.Atoi(string(line[i]))
			if e == nil {
				if first == -1 {
					first = n
				}
				last = n
			} else {
				n = -1
				for k, v := range numbers {
					if len(k) <= len(line)-i {
						candidate := string(line[i : i+len(k)])
						if candidate == k {
							n = v
							break
						}
					}
				}
				if n > -1 {
					if first == -1 {
						first = n
					}
					last = n
				}
			}
		}
		total += (first * 10) + last
	}
	fmt.Println(total)
}
