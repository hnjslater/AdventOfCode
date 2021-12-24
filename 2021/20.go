package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

var input = flag.String("input", "input20.txt", "")
var part = flag.Int("part", 1, "")

type Point struct {
	x, y int
}

func Min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func Max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	lookup := scanner.Text()
	image := make(map[Point]bool)
	y := 0
	scanner.Scan()
	for scanner.Scan() {
		text := scanner.Text()
		for x, r := range text {
			if r == '#' {
				image[Point{x, y}] = true
			} else {
				image[Point{x, y}] = false
			}
		}
		y++
	}

	iterations := 2
	if *part == 2 {
		iterations = 50
	}
	for i := 0; i < iterations; i++ {
		image2 := make(map[Point]bool)
		var minx, miny, maxx, maxy int = 0, 0, 0, 0
		for p := range image {
			minx = Min(minx, p.x)
			miny = Min(miny, p.y)
			maxx = Max(maxx, p.x)
			maxy = Max(maxy, p.y)
		}

		for x := minx - 1; x <= maxx+1; x++ {
			for y := miny - 1; y <= maxy+1; y++ {
				num := 0
				var exp float64 = 8
				for yy := y - 1; yy <= y+1; yy++ {
					for xx := x - 1; xx <= x+1; xx++ {
						v := (i % 2) == 1
						if v2, ok := image[Point{xx, yy}]; ok {
							v = v2
						}
						if v {
							num = num + int(math.Pow(2, exp))
						}
						exp--
					}
				}
				image2[Point{x, y}] = (lookup[num] == '#')
			}
		}
		image = image2
	}
	count := 0
	for k := range image {
		if image[k] {
			count++
		}
	}
	fmt.Println(count)
}
