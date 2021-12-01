package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var input_file = flag.String("input", "input.txt", "Input file")
var debug = flag.Bool("debug", false, "Debug")

type Coords struct {
	x, y int
}

func neighbour(coords Coords, direction string) Coords {
	switch direction {
	case "e":
		return Coords{coords.x + 1, coords.y}
	case "w":
		return Coords{coords.x - 1, coords.y}
	case "nw":
		return Coords{coords.x, coords.y - 1}
	case "sw":
		return Coords{coords.x - 1, coords.y + 1}
	case "ne":
		return Coords{coords.x + 1, coords.y - 1}
	case "se":
		return Coords{coords.x, coords.y + 1}
	default:
		log.Fatal("Unknown Direction.")
		return coords
	}
}

var compass = []string{"e", "w", "nw", "sw", "ne", "se"}

func countNeighbours(floor map[Coords]bool, coords Coords) int {
	count := 0
	for _, c := range compass {
		if floor[neighbour(coords, c)] {
			count++
		}
	}
	return count
}

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	floor := make(map[Coords]bool)

	for scanner.Scan() {
		line := scanner.Text()
		var coords Coords
		for i := 0; i < len(line); i++ {
			direction := line[i : i+1]
			if direction == "n" || direction == "s" {
				i++
				direction += line[i : i+1]
			}
			coords = neighbour(coords, direction)

		}
		if floor[coords] {
			delete(floor, coords)
		} else {
			floor[coords] = true
		}
	}
        for g := 0; g < 100; g++ {
                todo := make(map[Coords]bool)
                next := make(map[Coords]bool)

                for k := range floor {
                        count := countNeighbours(floor, k)
                        if count == 1 || count == 2 {
                                next[k] = true
                        }
                        for _, d := range compass {
                                todo[neighbour(k, d)] = true
                        }
                }

                for k := range todo {
                        if countNeighbours(floor, k) == 2 {
                                next[k] = true
                        }
                }

                floor = next
        }
	fmt.Println(len(floor))
}
