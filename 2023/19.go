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

var input = flag.String("input", "19.txt", "")
var part = flag.Int("part", 0, "")

type rule struct {
	category byte
	comp     byte
	value    int
	workflow string
}

type workflow struct {
	rules    []rule
	fallback string
}

func CopyParts(part map[byte][]int) map[byte][]int {
	part2 := make(map[byte][]int)
	for k, v := range part {
		part2[k] = make([]int, 2)
		copy(part2[k], v)
	}
	return part2
}

func Max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func Min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func process(rules map[string]workflow, part map[byte][]int, curr string, path []string) int {
	result := 0
	if curr == "A" {
		result = 1
		for _, v := range part {
			result = result * (v[1] - v[0])
		}
		return result
	}
	if curr == "R" {
		return 0
	}
	for _, r := range rules[curr].rules {
		part2 := CopyParts(part)
		if r.comp == '>' {
			part2[r.category][0] = Max(part2[r.category][0], r.value)
			part[r.category][1] = Min(part[r.category][1], r.value)
			result += process(rules, part2, r.workflow, append(path, curr))
		} else {
			part2[r.category][1] = Min(part2[r.category][1], r.value-1)
			part[r.category][0] = Max(part[r.category][0], r.value-1)
			result += process(rules, part2, r.workflow, append(path, curr))
		}
	}
	part2 := CopyParts(part)
	result += process(rules, part2, rules[curr].fallback, append(path, curr))

	return result
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
	rules := make(map[string]workflow)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var w workflow
		nameEnd := strings.Index(line, "{")
		name := line[0:nameEnd]

		for _, rulestr := range strings.Split(line[strings.Index(line, "{")+1:strings.Index(line, "}")], ",") {
			if strings.Index(rulestr, ":") == -1 {
				w.fallback = rulestr
			} else {
				var r rule
				r.category = rulestr[0]
				r.comp = rulestr[1]
				r.value, err = strconv.Atoi(rulestr[2:strings.Index(rulestr, ":")])
				if err != nil {
					panic(err)
				}
				r.workflow = rulestr[strings.Index(rulestr, ":")+1:]
				w.rules = append(w.rules, r)
			}
		}
		rules[name] = w
	}

	if *part == 0 {
		for scanner.Scan() {
			line := scanner.Text()
			item := make(map[byte]int)

			for _, itemstr := range strings.Split(line[1:len(line)-1], ",") {
				item[itemstr[0]], _ = strconv.Atoi(itemstr[2:])
			}
			w := "in"

			for w != "R" && w != "A" {
				matched := false
				for _, r := range rules[w].rules {
					if r.comp == '>' {
						if item[r.category] > r.value {
							w = r.workflow
							matched = true
							break
						}
					} else {
						if item[r.category] < r.value {
							w = r.workflow
							matched = true
							break
						}
					}
				}
				if !matched {
					w = rules[w].fallback
				}

			}

			if w == "A" {
				for _, c := range item {
					total += c
				}
			}
		}
	} else {
		part := make(map[byte][]int)
		for _, c := range []byte{'x', 'm', 'a', 's'} {
			part[c] = []int{0, 4000}
		}

		total = process(rules, part, "in", []string{})
	}
	fmt.Println(total)
}
