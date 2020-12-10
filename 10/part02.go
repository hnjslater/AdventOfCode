package main

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
	counts := make([]int, len(buffer))
	counts[0] = 1
	for i := 0; i < len(buffer); i++ {
		if i+1 < len(buffer) && buffer[i+1]-buffer[i] <= 3 {
			counts[i+1] += counts[i]
		}
		if i+2 < len(buffer) && buffer[i+2]-buffer[i] <= 3 {
			counts[i+2] += counts[i]
		}
		if i+3 < len(buffer) && buffer[i+3]-buffer[i] <= 3 {
			counts[i+3] += counts[i]
		}
	}
	fmt.Println(counts[len(counts)-1])
}
