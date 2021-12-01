package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var input = flag.String("input", "", "")
var part = flag.Int("part", 0, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	previous := math.MaxInt64
	window_size := 1
	if *part == 2 {
		window_size = 3
	}
	d := make([]int, 0)

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		d = append(d, value)
		if len(d) > window_size {
			d = d[1:]
		}
		if len(d) == window_size {
			current := 0
			for _, i := range d {
				current += i
			}
			if current > previous {
				count++
			}
			previous = current
		}
	}
	fmt.Println(count)
}
