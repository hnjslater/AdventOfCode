package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var input = flag.String("input", "18.txt", "")
var part = flag.Int("part", 0, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	height := 0
	fudge := 0

	for scanner.Scan() {
		line := scanner.Text()
		var direction byte
		var count int
		if *part == 0 {
			direction = line[0]
			count, _ = strconv.Atoi(string(line[2]))
			if line[3] != ' ' {
				count2, _ := strconv.Atoi(string(line[3]))
				count = (count * 10) + count2
			}
		} else if *part == 1 {
			direction = line[len(line)-2]
			blah := line[len(line)-7 : len(line)-2]
			count2, _ := strconv.ParseInt(blah, 16, 32)
			count = int(count2)
		}
		diff := 0
		switch direction {
		case 'U', '3':
			height = height - count
		case 'D', '1':
			height = height + count
		case 'L', '2':
			diff = (height * count)
		case 'R', '0':
			diff = -(height * count)
		}
		total += diff
		fudge += count
	}
	fmt.Println(total + (fudge / 2) + 1)
}
