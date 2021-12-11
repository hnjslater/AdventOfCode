package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var input = flag.String("input", "input11.txt", "")
var part = flag.Int("part", 1, "")

type Coords struct {
	x int
	y int
}

type Octopus struct {
	energy  int
	flashed bool
}

type Floor map[Coords]Octopus

func (f Floor) Neighbors(c Coords) []Coords {
	var n []Coords
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			candidate := Coords{c.x + dx, c.y + dy}
			if _, present := f[candidate]; present && candidate != c {
				n = append(n, candidate)
			}
		}
	}
	return n
}

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	floor := make(Floor)
	x := 0
	for scanner.Scan() {
		line := scanner.Text()
		for y, r := range line {
			e, _ := strconv.Atoi(string(r))
			floor[Coords{x, y}] = Octopus{e, false}
		}
		x++
	}
	max_steps := 100
	if *part == 2 {
		max_steps = 10000
	}
stepper:
	for s := 1; s <= max_steps; s++ {
		for c, _ := range floor {
			floor[c] = Octopus{floor[c].energy + 1, false}
		}
		flashes := 0
		changed := true
		for changed {
			changed = false
			for c, _ := range floor {
				if floor[c].energy > 9 && !floor[c].flashed {
					flashes++
					changed = true
					for _, cn := range floor.Neighbors(c) {
						if !floor[cn].flashed {
							floor[cn] = Octopus{floor[cn].energy + 1, floor[cn].flashed}
						}
					}
					floor[c] = Octopus{0, true}
				}
			}
		}
		if *part == 1 {
			result += flashes
		} else {
			if flashes == len(floor) {
				result = s
				break stepper
			}
		}
	}
	fmt.Println(result)
}
