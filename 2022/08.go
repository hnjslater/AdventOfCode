package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type location struct {
	r int
	c int
}

func (l location) Add(d location) location {
	return location{l.r + d.r, l.c + d.c}
}

type Forest map[location]int

func (f Forest) Contains(l location) bool {
	_, ok := f[l]
	return ok
}

var input = flag.String("input", "input08.txt", "")
var part = flag.Int("part", 0, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	count := 0
	scanner := bufio.NewScanner(file)
	r := 0
	forest := make(Forest)
	for scanner.Scan() {
		line := scanner.Text()
		for c, ru := range line {
			h, _ := strconv.Atoi(string(ru))
			forest[location{r, c}] = h
		}
		r++
	}
	if *part == 0 {
		for start, starth := range forest {
			visible := 4
			for _, d := range []location{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
				l := start
				l = start.Add(d)
				for forest.Contains(l) {
					if l != start && forest[l] >= starth {
						visible--
						break
					}

					l = l.Add(d)
				}
			}
			if visible > 0 {
				count++
			}
		}
	} else {
		for start, starth := range forest {
			curcount := 1
			for _, d := range []location{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
				l := start
				innercount := 0
				l = l.Add(d)
				for forest.Contains(l) {
					innercount++
					if l != start && forest[l] >= starth {
						break
					}
					l = l.Add(d)
				}
				curcount *= innercount
				innercount = 0
			}
			if curcount > count {
				count = curcount
			}
		}
	}
	fmt.Println(count)
}
