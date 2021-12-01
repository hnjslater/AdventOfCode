package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var input_file = flag.String("input", "input.txt", "Input file")

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var buffer [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		line = "L" + line + "L"
		if len(buffer) == 0 {
			buffer = append(buffer, []byte(strings.Repeat("L", len(line))))
		}
		buffer = append(buffer, []byte(line))
	}
	buffer = append(buffer, []byte(strings.Repeat("L", len(buffer[0]))))
	changed := true
	for changed {
		changed = false
		next_buffer := make([][]byte, len(buffer))

		next_buffer[0] = make([]byte, len(buffer[0]))
		copy(next_buffer[0], buffer[0])

		next_buffer[len(buffer)-1] = make([]byte, len(buffer[0]))
		copy(next_buffer[len(buffer)-1], buffer[len(buffer)-1])

		for r := 1; r < len(buffer)-1; r++ {
			next_buffer[r] = make([]byte, len(buffer[r]))
			copy(next_buffer[r], buffer[r])
			for c := 1; c < len(buffer[r])-1; c++ {
				if buffer[r][c] == '.' {
					continue
				}
				count := 0
				for _, ri := range []int{-1, 0, +1} {
					for _, ci := range []int{-1, 0, +1} {
						if ri != 0 || ci != 0 {
						loop:
							for i := 1; ; i++ {
								switch buffer[r+(ri*i)][c+(ci*i)] {
								case 'L':
									break loop
								case '#':
									count++
									break loop
								}
							}
						}
					}
				}
				if buffer[r][c] == 'L' && count == 0 {
					next_buffer[r][c] = '#'
					changed = true
				} else if buffer[r][c] == '#' && count >= 5 {
					next_buffer[r][c] = 'L'
					changed = true
				}
			}
		}

		buffer = next_buffer
	}
	count := 0
	for r := 0; r < len(buffer); r++ {
		for c := 0; c < len(buffer[r]); c++ {
			if buffer[r][c] == '#' {
				count++
			}
		}
	}
	fmt.Println(count)
}
