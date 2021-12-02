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

var input = flag.String("input", "", "")
var part = flag.Int("part", 1, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	depth := 0
	forward := 0
	aim := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		direction := line[0]
		value, _ := strconv.Atoi(line[1])

		if *part == 1 {
			if direction == "up" {
				depth -= value
			} else if direction == "down" {
				depth += value
			} else if direction == "forward" {
				forward += value
			}
		} else {
			if direction == "up" {
				aim -= value
			} else if direction == "down" {
				aim += value
			} else if direction == "forward" {
				forward += value
				depth += (aim * value)
			}

		}
	}
	fmt.Println(depth * forward)
}
