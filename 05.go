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

var input = flag.String("input", "test05.txt", "")
var part = flag.Int("part", 0, "")

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

	var fishies [9]int
	for _, entry := range strings.Split(line, ",") {
		x, _ := strconv.Atoi(entry)
		fishies[x]++
	}
	generations := 80
	if *part == 2 {
		generations = 256
	}

	for g := 0; g < generations; g++ {
		breeders := fishies[0]
		for i := 1; i <= 8; i++ {
			fishies[i-1] = fishies[i]
		}
		fishies[8] = breeders
		fishies[6] += breeders
	}
	pop := 0
	for _, e := range fishies {
		pop += e
	}
	fmt.Println(pop)
}
