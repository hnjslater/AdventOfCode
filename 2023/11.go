package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var input = flag.String("input", "11.txt", "")
var part = flag.Int("part", 0, "")

type coord struct {
	row int
	col int
}

func Sort(x int, y int) (int, int) {
	if x < y {
		return x, y
	} else {
		return y, x
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

	total := 0

	universe := make(map[coord]byte)
	rows_seen := make(map[int]bool)
	cols_seen := make(map[int]bool)

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for col := range line {
			if line[col] == '#' {
				universe[coord{row + 1, col + 1}] = line[col]
				rows_seen[row+1] = true
				cols_seen[col+1] = true
			}
		}
		row++
	}

	for first := range universe {
		for second := range universe {
			start_row, end_row := Sort(first.row, second.row)
			start_col, end_col := Sort(first.col, second.col)

			dist := 0
			for row := start_row; row < end_row; row++ {
				dist++
				if !rows_seen[row] {
					if *part == 0 {
						dist += 1
					} else {
						dist += 999999
					}
				}
			}
			for col := start_col; col < end_col; col++ {
				dist++
				if !cols_seen[col] {
					if *part == 0 {
						dist += 1
					} else {
						dist += 999999
					}
				}
			}

			total += dist

		}
	}
	fmt.Println(universe)
	fmt.Println(total / 2)
}
