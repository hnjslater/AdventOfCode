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

var input = flag.String("input", "input09.txt", "")
var part = flag.Int("part", 1, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	var seafloor [][]int
	seafloor = append(seafloor, []int{})
	for scanner.Scan() {
		line := scanner.Text()
		sline := make([]int, len(line)+2)
		sline[0] = 100
		sline[len(line)+1] = 100
		for i, r := range line {
			sline[i+1], _ = strconv.Atoi(string(r))
		}
		seafloor = append(seafloor, sline)
	}
	padding := make([]int, len(seafloor[1]))
	for i := range seafloor[1] {
		padding[i] = 100
	}
	seafloor = append(seafloor, padding)
	seafloor[0] = padding

	var basins [][2]int
	for x := 1; x < len(seafloor)-1; x++ {
		for y := 1; y < len(seafloor[0])-1; y++ {
			pnt := seafloor[x][y]
			if pnt < seafloor[x-1][y] && pnt < seafloor[x+1][y] && pnt < seafloor[x][y-1] && pnt < seafloor[x][y+1] {
				basins = append(basins, [...]int{x, y})
			}
		}
	}
	if *part == 1 {
		for _, b := range basins {
			result += seafloor[b[0]][b[1]] + 1
		}
	} else {
		for i, b := range basins {
			seafloor[b[0]][b[1]] = -i - 1
		}

		changed := 1
		for changed != 0 {
			changed = 0
			for x := 1; x < len(seafloor)-1; x++ {
				for y := 1; y < len(seafloor[0])-1; y++ {
					pnt := seafloor[x][y]
					if pnt < 9 && pnt >= 0 {
						if seafloor[x-1][y] < 0 {
							pnt = seafloor[x-1][y]
						} else if seafloor[x+1][y] < 0 {
							pnt = seafloor[x+1][y]
						} else if seafloor[x][y+1] < 0 {
							pnt = seafloor[x][y+1]
						} else if seafloor[x][y-1] < 0 {
							pnt = seafloor[x][y-1]
						}
						if seafloor[x][y] != pnt {
							changed++
							seafloor[x][y] = pnt
						}
					}
				}
			}
		}

		basin_sizes := make(map[int]int)

		for x := 1; x < len(seafloor)-1; x++ {
			for y := 1; y < len(seafloor[0])-1; y++ {
				if seafloor[x][y] < 0 {
					basin_sizes[seafloor[x][y]]++
				}
			}
		}

		var sizes []int
		for _, v := range basin_sizes {
			sizes = append(sizes, v)
		}
		sort.IntSlice.Sort(sizes)
		result = 1
		for _, x := range sizes[len(sizes)-3:] {
			result *= x
		}
	}
	fmt.Println(result)
}
