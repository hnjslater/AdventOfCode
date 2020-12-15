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

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	input_str := strings.Split(line, ",")
	history := make(map[int][]int)
	prev_value := 0
	for i, v := range input_str {
		prev_value, _ = strconv.Atoi(v)
		history[prev_value] = []int{i}
	}
	for i := len(history); i < 2020; i++ {
		last_seen := history[prev_value]
		value := 0
		if len(last_seen) > 1 {
			value = last_seen[len(last_seen)-1] - last_seen[len(last_seen)-2]
		}
		history[value] = append(history[value], i)
		prev_value = value

	}
	fmt.Println(prev_value)
}
