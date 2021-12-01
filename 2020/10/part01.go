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
	one_jolt := 0
	three_jolt := 0
	for i := 1; i < len(buffer); i++ {
		diff := (buffer[i] - buffer[i-1])
		if diff == 1 {
			one_jolt++
		} else if diff == 3 {
			three_jolt++
		}
	}
	fmt.Println(one_jolt * three_jolt)

}
