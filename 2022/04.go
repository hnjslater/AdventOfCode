package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var input = flag.String("input", "input04.txt", "")
var part = flag.Int("part", 0, "")

var lineExp = regexp.MustCompile(`([[:digit:]]+)-([[:digit:]]+),([[:digit:]]+)-([[:digit:]]+)`)

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match := lineExp.FindStringSubmatch(line)

		r00, _ := strconv.Atoi(match[1])
		r01, _ := strconv.Atoi(match[2])
		r10, _ := strconv.Atoi(match[3])
		r11, _ := strconv.Atoi(match[4])

		if r00 <= r10 && r01 >= r11 {
			total++
		} else if r10 <= r00 && r11 >= r01 {
			total++
		} else if *part == 1 {
			if r00 <= r10 && r10 <= r01 {
				total++
			} else if r00 <= r11 && r11 <= r01 {
				total++
			} else if r10 <= r00 && r00 <= r11 {
				total++
			} else if r10 <= r01 && r01 <= r11 {
				total++
			}
		}
	}
	fmt.Println(total)
}
