package main

import (
	"flag"
	"fmt"
)

var input = flag.String("input", "test16.txt", "")
var part = flag.Int("part", 1, "")

type Coords struct {
	x int
	y int
}

func RunTrial(d Coords, t1 Coords, t2 Coords) (bool, int) {
	max_y := -1000
	x := 0
	y := 0
	for i := 0; i < 10000; i++ {
		x += d.x
		y += d.y
		if y > max_y {
			max_y = y
		}
		if d.x > 0 {
			d.x -= 1
		} else if d.x < 0 {
			d.x += 1
		}
		d.y -= 1

		if x >= t1.x && x <= t2.x && y >= t1.y && y <= t2.y {
			return true, max_y
		}
	}
	return false, -1
}

func main() {
	// This is bad and I feel bad.
	t1 := Coords{x: 144, y: -100}
	t2 := Coords{x: 178, y: -76}

	result := 0
	count := 0
	for x := -1000; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			hit, max_y := RunTrial(Coords{x, y}, t1, t2)
			if hit && max_y > result {
				result = max_y
			}
			if hit {
				count++
			}
		}
	}
	fmt.Println(result)
	fmt.Println(count)
}
