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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Bus struct {
	arrival int
	number  int
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
	scanner.Text()
	if err != nil {
		log.Fatal(err)
	}
	max_bus := Bus{0, 0}
	var buses []Bus
	for scanner.Scan() {
		line := scanner.Text()
		for arrival, busstr := range strings.Split(line, ",") {
			bus, err := strconv.Atoi(busstr)
			if err != nil {
				continue
			}
			buses = append(buses, Bus{arrival, bus})
			if bus > max_bus.number {
				max_bus = Bus{arrival, bus}
			}
		}

	}
outer:
	for i := 0; ; i++ {
		time := (max_bus.number * i) - max_bus.arrival
		for _, bus := range buses {
			if (time+bus.arrival)%bus.number != 0 {
				continue outer
			}
		}
		fmt.Print(time)
		break outer
	}
}
