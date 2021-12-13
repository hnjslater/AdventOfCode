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

var input = flag.String("input", "input12.txt", "")
var part = flag.Int("part", 1, "")

type Coords struct {
	x int
	y int
}

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sheet := make(map[Coords]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		spline := strings.Split(line, ",")
		x, _ := strconv.Atoi(spline[1])
		y, _ := strconv.Atoi(spline[0])

		sheet[Coords{x, y}] = true
	}

	for scanner.Scan() {
		fold_line := scanner.Text()
		fold, _ := strconv.Atoi(strings.Split(fold_line, "=")[1])
		fold_dir := fold_line[11]

		switch fold_dir {
		case 'x':
			for c := range sheet {
				if c.y > fold {
					delete(sheet, c)
					sheet[Coords{c.x, fold - (c.y - fold)}] = true
				}
			}
		case 'y':
			for c := range sheet {
				if c.x > fold {
					delete(sheet, c)
					sheet[Coords{fold - (c.x - fold), c.y}] = true
				}
			}
		}
		if *part == 1 {
			fmt.Println(len(sheet))
			break
		}
	}

	if *part == 2 {
		max_x := 0
		max_y := 0
		for k := range sheet {
			if k.x > max_x {
				max_x = k.x
			}
			if k.y > max_y {
				max_y = k.y
			}
		}

		for x := 0; x <= max_x; x++ {
			for y := 0; y <= max_y; y++ {
				if !sheet[Coords{x, y}] {
					fmt.Print(".")
				} else {
					fmt.Print("#")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
