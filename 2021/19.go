package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var input = flag.String("input", "example19b.txt", "")
var part = flag.Int("part", 1, "")

func Abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

type Point struct {
	x, y, z int
}

type Scan map[Point]bool

func (s Scan) First() Point {
	for p := range s {
		return p
	}
	panic("Empty map")
}

func (p Point) Roll() Point {
	return Point{p.x, p.z, -p.y}
}

func (p Point) Turn() Point {
	return Point{-p.y, p.x, p.z}
}

func (p Point) Rotate(orientation int) Point {
	for i := 0; i < orientation; i++ {
		if i%4 == 0 {
			p = p.Roll()
		} else {
			p = p.Turn()
		}
		if i == 11 {
			p = p.Roll().Turn().Roll()
		}
	}
	return p
}

func (p Point) Translate(t Point) Point {
	return Point{p.x + t.x, p.y + t.y, p.z + t.z}
}

func Check(s1 Scan, s2 Scan, ori int, t Point) int {
	overlap := 0
	for p := range s2 {
		if s1[p.Rotate(ori).Translate(t)] {
			overlap++
		}
	}
	return overlap
}

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var scanners []Scan
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 2 && text[0:2] == "--" {
			scanners = append(scanners, Scan{})
		} else if len(text) > 0 {
			var point Point
			nums := strings.Split(text, ",")
			point.x, _ = strconv.Atoi(nums[0])
			point.y, _ = strconv.Atoi(nums[1])
			point.z, _ = strconv.Atoi(nums[2])
			scanners[len(scanners)-1][point] = true
		}
	}

	found := Scan{}
	for k := range scanners[0] {
		found[k] = true
	}
	done := make(map[int]bool)
	done[0] = true
	true_scanners := make(map[Point]bool)
outer:
	for true {
		for i := 0; i < len(scanners); i++ {
		sloop:
			for p1 := range scanners[i] {
				for p0 := range found {
					if done[i] {
						continue
					}
					for r := 0; r < 24; r++ {
						p1r := p1.Rotate(r)
						t := Point{p0.x - p1r.x, p0.y - p1r.y, p0.z - p1r.z}

						if Check(found, scanners[i], r, t) >= 6 {
							for k := range scanners[i] {
								found[k.Rotate(r).Translate(t)] = true
							}
							done[i] = true
							true_scanners[t] = true
							if len(done) == len(scanners) {
								break outer
							} else {
								continue sloop
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(len(found))

	distance := 0
	for p0 := range true_scanners {
		for p1 := range true_scanners {
			cd := Abs(p0.x - p1.x)
			cd += Abs(p0.y - p1.y)
			cd += Abs(p0.z - p1.z)
			if cd > distance {
				distance = cd
			}
		}
	}

	fmt.Println(distance)
}
