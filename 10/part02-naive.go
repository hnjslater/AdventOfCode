package main

// After the fact, I came back and fixed my naive solution.

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var input_file = flag.String("input", "input.txt", "Input file")

func CountSubtrees2(start int, buffer []int, cache map[int]int) int {
	value, found := cache[start]
	if found {
		return value
	}
	if len(buffer)-start == 1 {
		return 1
	}

	total := 0
	if buffer[start+1]-buffer[start] <= 3 {
		total += CountSubtrees2(start+1, buffer, cache)
	}
	if len(buffer)-start > 2 && buffer[start+2]-buffer[start] <= 3 {
		total += CountSubtrees2(start+2, buffer, cache)
	}
	if len(buffer)-start > 3 && buffer[start+3]-buffer[start] <= 3 {
		total += CountSubtrees2(start+3, buffer, cache)
	}
	cache[start] = total
	return total
}

func CountSubtrees(start int, buffer []int) int {
	cache := make(map[int]int)
	return CountSubtrees2(start, buffer, cache)
}

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var buffer []int
	buffer = append(buffer, 0)
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Error parsing input")
		}
		buffer = append(buffer, value)
	}
	sort.Ints(buffer)
	buffer = append(buffer, buffer[len(buffer)-1]+3)

	fmt.Println(CountSubtrees(0, buffer))
}
