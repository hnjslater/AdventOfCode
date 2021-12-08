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

var input = flag.String("input", "test06.txt", "")
var part = flag.Int("part", 1, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	var crabs []int
	for _, entry := range strings.Split(line, ",") {
		x, _ := strconv.Atoi(entry)
		crabs = append(crabs, x)
	}

	mincrab := crabs[0]
	maxcrab := crabs[1]

	for _, c := range crabs {
		if c > maxcrab {
			maxcrab = c
		}
		if c < mincrab {
			mincrab = c
		}
	}
	fuel := math.MaxInt64

	for i := mincrab; i <= maxcrab; i++ {
		cfuel := 0
		for _, c := range crabs {
			ccfuel := int(math.Abs(float64(i - c)))
			if *part == 1 {
				cfuel += ccfuel
			} else {
				cfuel += (ccfuel * (ccfuel + 1)) / 2
			}
		}
		if cfuel < fuel {
			fuel = cfuel
		}
	}
	fmt.Println(fuel)
}
