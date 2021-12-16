package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var input = flag.String("input", "input15.txt", "")
var part = flag.Int("part", 1, "")

type Coords struct {
	x int
	y int
}

type Candidate struct {
	v int
	c Coords
}

func get(world map[Coords]int, cs Coords, max int) int {
	val := world[cs]
	if cs.x >= max {
		val = get(world, Coords{cs.x - max, cs.y}, max) + 1
	} else if cs.y >= max {
		val = get(world, Coords{cs.x, cs.y - max}, max) + 1
	}
	if val == 10 {
		val = 1
	}
	return val
}

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	world := make(map[Coords]int)
	y := 0
	ws := 0
	for scanner.Scan() {
		line := scanner.Text()
		if ws == 0 {
			ws = len(line)
		}
		for x, r := range line {
			value, _ := strconv.Atoi(string(r))
			world[Coords{x, y}] = value
		}
		y++
	}

	done := make(map[Coords]int)
	candidates := []Candidate{Candidate{0, Coords{0, 0}}}

	max := ws
	if *part == 2 {
		max = ws * 5
	}

	for len(candidates) > 0 {
		//FIXME this whole loop is horrible, check every place 4 times is not optimal!
		c := candidates[len(candidates)-1]
		candidates = candidates[:len(candidates)-1]
		if _, ok := done[c.c]; ok {
			continue
		}
		done[c.c] = c.v

		if c.c.x+1 < max {
			cs := Coords{c.c.x + 1, c.c.y}
			candidates = append(candidates, Candidate{get(world, cs, ws) + c.v, cs})
		}
		if c.c.y+1 < max {
			cs := Coords{c.c.x, c.c.y + 1}
			candidates = append(candidates, Candidate{get(world, cs, ws) + c.v, cs})
		}
		if c.c.x > 0 {
			cs := Coords{c.c.x - 1, c.c.y}
			candidates = append(candidates, Candidate{get(world, cs, ws) + c.v, cs})
		}
		if c.c.y > 0 {
			cs := Coords{c.c.x, c.c.y - 1}
			candidates = append(candidates, Candidate{get(world, cs, ws) + c.v, cs})
		}

		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].v > candidates[j].v
		})
	}

	fmt.Println(done[Coords{max - 1, max - 1}])
}
