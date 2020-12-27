package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var input_file = flag.String("input", "input.txt", "Input file")

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	first_line := scanner.Text()
	arrival, err := strconv.Atoi(first_line)
	if err != nil {
		log.Fatal(err)
	}
	min_wait := math.MaxInt64
	best_bus := 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, busstr := range strings.Split(line, ",") {
			bus, err := strconv.Atoi(busstr)
			if err != nil {
				continue
			}

			wait := ((arrival / bus) * bus) - arrival
			if wait < 0 {
				wait += bus
			}
			if wait < min_wait {
				min_wait = wait
				best_bus = bus
			}
		}
	}

	fmt.Println(min_wait * best_bus)
}
