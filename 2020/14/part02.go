package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var input_file = flag.String("input", "input.txt", "Input file")
var write_re = regexp.MustCompile("mem\\[(?P<address>[0-9]+)\\] = (?P<value>[0-9]+)")

// Sigh
func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	memory := make(map[uint64]uint64)
	var mask uint64
	var floating_bits []int
	for scanner.Scan() {
		line := scanner.Text()
		if line[1] == 'a' {
			mask_str := line[7:]
			mask_str = strings.ReplaceAll(mask_str, "X", "0")
			mask, err = strconv.ParseUint(mask_str, 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			floating_bits = make([]int, 0)
			for i, r := range line[7:] {
				if r == 'X' {
					floating_bits = append(floating_bits, 35-i)
				}
			}
		} else {
			matches := write_re.FindStringSubmatch(line)
			address, err := strconv.ParseUint(matches[write_re.SubexpIndex("address")], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			wval, err := strconv.ParseUint(matches[write_re.SubexpIndex("value")], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			address = address | mask
			for i := 0; i < powInt(2, len(floating_bits)); i++ {
				address2 := address
				for j, v := range floating_bits {
					if i&(1<<j) > 0 {
						address2 = address2 | (1 << v)
					} else {
						address2 = address2 &^ (1 << v)
					}
				}
				memory[address2] = wval
			}
		}
	}

	var total uint64 = 0
	for _, v := range memory {
		total += v
	}
	fmt.Println(total)
}
