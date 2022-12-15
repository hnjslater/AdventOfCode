package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

var input = flag.String("input", "input12.txt", "")
var part = flag.Int("part", 0, "")

type loc struct {
	r int
	c int
}

type World map[loc]int

func (w World) Contains(k loc) bool {
	_, ok := w[k]
	return ok
}

func (w World) GetSmallest() loc {
	minLoc := loc{-1, -1}

	for k, v := range w {
		if minLoc.r == -1 {
			minLoc = k
		} else if w[minLoc] > v {
			minLoc = k
		}
	}

	return minLoc
}

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	world := make(World)
	var start loc
	var end loc

	r := 0
	for scanner.Scan() {
		line := scanner.Text()
		for c := 0; c < len(line); c++ {
			b := line[c]
			if b == 'S' {
				start = loc{r, c}

			} else if b == 'E' {
				end = loc{r, c}
			} else {
				world[loc{r, c}] = int(b - 'a')
			}
		}
		r++
	}
	world[start] = 0
	world[end] = 'z' - 'a'

	candidates := make(World)
	candidates[end] = 0
	done := make(World)

	for len(candidates) > 0 {
		l := candidates.GetSmallest()
		done[l] = candidates[l]
		for _, l2 := range []loc{loc{l.r - 1, l.c}, loc{l.r + 1, l.c}, loc{l.r, l.c - 1}, loc{l.r, l.c + 1}} {
			if done.Contains(l2) || !world.Contains(l2) {
				continue
			} else if world[l] > world[l2]+1 {
				continue
			} else if !candidates.Contains(l2) || (candidates.Contains(l2) && candidates[l2] > candidates[l]+1) {
				candidates[l2] = candidates[l] + 1
			}

		}
		delete(candidates, l)

	}
	if *part == 0 {
		fmt.Println(done[start])
	} else {
		min := math.MaxInt64
		for k, v := range world {
			if v == 0 && done[k] < min && done[k] > 0 {
				min = done[k]
			}
		}
		fmt.Println(min)
	}

}
