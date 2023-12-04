package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var input = flag.String("input", "03.txt", "")
var part = flag.Int("part", 0, "")

type coords struct {
	r int
	c int
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

	parts := make(map[coords]byte)
	gears := make(map[coords]int)
	r := 0
	for scanner.Scan() {
		line := scanner.Text()

		for c, _ := range line {
			parts[coords{r, c}] = line[c]
		}
		r++
	}

	for coord, v := range parts {

		c := coord.c
		r := coord.r

		if !unicode.IsDigit(rune(v)) {
			continue
		}
		if unicode.IsDigit(rune(parts[coords{r, c - 1}])) {
			continue
		}

		var pals []coords

		pals = append(pals, coords{r, c - 1})
		pals = append(pals, coords{r + 1, c - 1})
		pals = append(pals, coords{r - 1, c - 1})

		val := ""
		for w := 0; unicode.IsDigit(rune(parts[coords{r, c + w}])); w++ {
			val += string(parts[coords{r, c + w}])
			pals = append(pals, coords{r - 1, c + w})
			pals = append(pals, coords{r + 1, c + w})
		}
		value, _ := strconv.Atoi(val)

		pals = append(pals, coords{r, c + len(val)})
		pals = append(pals, coords{r - 1, c + len(val)})
		pals = append(pals, coords{r + 1, c + len(val)})

		if *part == 0 {
			for _, coord := range pals {
				if parts[coord] != 0 && rune(parts[coord]) != '.' {
					total += value
					break
				}
			}
		} else {
			for _, coord := range pals {
				if rune(parts[coord]) == '*' {
					gear := gears[coord]
					if gear == 0 {
						gears[coord] = value
					} else {
						total += gears[coord] * value
					}
				}
			}

		}
	}
	fmt.Println(total)
}
