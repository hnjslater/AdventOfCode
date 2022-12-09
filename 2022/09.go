package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type location struct {
	r int
	c int
}

func (l location) Add(d location) location {
	return location{l.r + d.r, l.c + d.c}
}

var input = flag.String("input", "input09.txt", "")
var part = flag.Int("part", 0, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	visited := make(map[location]bool)
	var snake []location
	if *part == 0 {
		snake = make([]location, 2)
	} else {
		snake = make([]location, 10)
	}
	for scanner.Scan() {
		line := scanner.Text()

		delta := location{}
		direction := line[0]
		magnitude, _ := strconv.Atoi(line[2:])

		switch direction {
		case 'U':
			delta = location{-1, 0}
		case 'D':
			delta = location{1, 0}
		case 'L':
			delta = location{0, -1}
		case 'R':
			delta = location{0, 1}
		}

		for i := 0; i < magnitude; i++ {
			snake[0] = snake[0].Add(delta)
			for i := 1; i < len(snake); i++ {
				h := &snake[i-1]
				t := &snake[i]

				extraMove := Abs(t.r-h.r)+Abs(t.c-h.c) > 2

				if Abs(t.r-h.r) > 1 || extraMove {
					if h.r > t.r {
						t.r += 1
					} else if h.r < t.r {
						t.r -= 1
					}
				}

				if Abs(t.c-h.c) > 1 || extraMove {
					if h.c > t.c {
						t.c += 1
					} else if h.c < t.c {
						t.c -= 1
					}
				}

				visited[snake[len(snake)-1]] = true
			}
		}
	}
	fmt.Println(len(visited))
}
