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
	x int
	y int
	z int
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func CountNeighbours(world map[Coords]bool, c Coords) int {
	n := 0
	for x := c.x - 1; x <= c.x+1; x++ {
		for y := c.y - 1; y <= c.y+1; y++ {
			for z := c.z - 1; z <= c.z+1; z++ {
				if world[Coords{x, y, z}] {
					n++
				}
			}
		}
	}
	if world[c] {
		n--
	}
	return n
}

func PrintWorld(world map[Coords]bool) {
	var minx, miny, minz, maxx, maxy, maxz int
	for k, _ := range world {
		minx = Min(k.x, minx)
		miny = Min(k.y, miny)
		minz = Min(k.z, minz)
		maxx = Max(k.x, maxx)
		maxy = Max(k.y, maxy)
		maxz = Max(k.z, maxz)
	}

	for z := minz; z <= maxz; z++ {
		fmt.Println("z=", z)
		for x := minx; x <= maxx; x++ {
			for y := miny; y <= maxy; y++ {
				if world[Coords{x, y, z}] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	world := make(map[Coords]bool)

	scanner := bufio.NewScanner(file)
	x := 0
	for scanner.Scan() {
		line := scanner.Text()
		for y, r := range line {
			if r == '#' {
				world[Coords{x, y, 0}] = true
			}
		}
		x++
	}

	for i := 0; i < 6; i++ {
		if *debug {
			fmt.Println("=============", i, "=============")
			PrintWorld(world)
		}
		var minx, miny, minz, maxx, maxy, maxz int
		for k, _ := range world {
			minx = Min(k.x, minx)
			miny = Min(k.y, miny)
			minz = Min(k.z, minz)
			maxx = Max(k.x, maxx)
			maxy = Max(k.y, maxy)
			maxz = Max(k.z, maxz)
		}
		world2 := make(map[Coords]bool)
		for x := minx - 1; x < maxx+2; x++ {
			for y := miny - 1; y < maxy+2; y++ {
				for z := minz - 1; z < maxz+2; z++ {
					c := Coords{x, y, z}
					n := CountNeighbours(world, c)
					if world[c] && (n == 2 || n == 3) {
						world2[c] = true
					} else if (!world[c]) && n == 3 {
						world2[c] = true
					}
				}
			}
		}
		world = world2
	}

	fmt.Println(len(world))
}
