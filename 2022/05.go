package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var input = flag.String("input", "input05.txt", "")
var part = flag.Int("part", 0, "")

var lineExp = regexp.MustCompile(`move ([[:digit:]]+) from ([[:digit:]]+) to ([[:digit:]]+)`)

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var stacks [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		if line[1] == '1' {
			break
		}

		if len(stacks) == 0 {
			num_stacks := (len(line) + 1) / 4
			stacks = make([][]rune, num_stacks)
		}
		for i := 0; i < len(stacks); i++ {
			col := (i * 4) + 1
			if line[col] != ' ' {
				stacks[i] = append(stacks[i], rune(line[col]))
			}
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		match := lineExp.FindStringSubmatch(line)
		if len(match) == 0 {
			continue
		}
		count, _ := strconv.Atoi(match[1])
		from, _ := strconv.Atoi(match[2])
		to, _ := strconv.Atoi(match[3])
		to = to - 1
		from = from - 1

		if *part == 0 {
			for i := 0; i < count; i++ {
				stacks[to] = append([]rune{stacks[from][0]}, stacks[to]...)
				stacks[from] = stacks[from][1:]
			}
		} else {
			new_to := make([]rune, count)
			copy(new_to, stacks[from][0:count])
			new_to = append(new_to, stacks[to]...)
			stacks[to] = new_to
			stacks[from] = stacks[from][count:]
		}
	}

	result := ""
	for i := 0; i < len(stacks); i++ {
		result = result + string(stacks[i][0])
	}
	fmt.Println(result)
}
