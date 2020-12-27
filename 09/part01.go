package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
	buffer := make([]int, 25)
	i := 0
	value := 0
outer:
	for scanner.Scan() {
		line := scanner.Text()
		value, err = strconv.Atoi(line)
		if err != nil {
			log.Fatal("Error parsing input")
		}
		if i < 25 {
			buffer[i] = value
		} else {
			found := false
		middle:
			for j := 0; j < len(buffer); j++ {
				for k := j + 1; k < len(buffer); k++ {
					if buffer[j]+buffer[k] == value {
						found = true
						break middle
					}
				}

			}
			if !found {
				break outer
			}
			buffer = append(buffer[1:], value)
		}
		i++
	}
	fmt.Println(value)
}
