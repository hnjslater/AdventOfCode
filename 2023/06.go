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

var input = flag.String("input", "06.txt", "")
var part = flag.Int("part", 0, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 1
	scanner.Scan()
	times := strings.Fields(scanner.Text())[1:]
	scanner.Scan()
	dists := strings.Fields(scanner.Text())[1:]

	if *part > 0 {
		times = []string{strings.Join(times, "")}
		dists = []string{strings.Join(dists, "")}
	}

	for i := 0; i < len(times); i++ {
		t, _ := strconv.ParseFloat(times[i], 4)
		d, _ := strconv.ParseFloat(dists[i], 4)

		root1 := (t - math.Sqrt(t*t-4*d)) / (2.0)
		root2 := (t + math.Sqrt(t*t-4*d)) / (2.0)

		fixed1 := int(root1) + 1
		fixed2 := int(root2)

		if float64(fixed2) == root2 {
			fixed2 -= 1
		}

		total = total * (fixed2 - fixed1 + 1)
	}
	fmt.Println(total)
}
