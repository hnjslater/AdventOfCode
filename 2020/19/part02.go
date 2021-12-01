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

var input_file = flag.String("input", "input.txt", "Input file")
var debug = flag.Bool("debug", false, "Debug")

var line_re = regexp.MustCompile("(?P<num>[0-9]+): (?P<pattern>.*)")

func ExpandRule(rules map[int]string, n int) string {
	if n == 8 {
		r42 := ExpandRule(rules, 42)
		return "(" + r42 + ")+"
	} else if n == 11 {
		output := "("
		for i := 1; i < 10; i++ {
			if i != 1 {
				output += "|"
			}
			output += "("
			for j := 0; j < i; j++ {
				output += ExpandRule(rules, 42)
			}
			for j := 0; j < i; j++ {
				output += ExpandRule(rules, 31)
			}
			output += ")"
		}
		output += ")"
		return output
	}

	output := ""
	acc := 0
	brackets := false
	for _, r := range []byte(rules[n]) {
		if r >= '0' && r <= '9' {
			acc = (acc * 10) + int(r-'0')
		} else {
			if acc > 0 {
				output += ExpandRule(rules, acc)
				acc = 0
			}
			if r == '|' {
				output = "(" + output + ")|("
				brackets = true
			} else if r == 'a' || r == 'b' {
				output += fmt.Sprintf("%c", r)
			}
		}
	}
	if acc > 0 {
		output += ExpandRule(rules, acc)
	}
	if brackets {
		output += ")"
	}
	return "(" + output + ")"
}
func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rules := make(map[int]string)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		matches := line_re.FindStringSubmatch(line)
		num, err := strconv.Atoi(matches[line_re.SubexpIndex("num")])
		if err != nil {
			log.Fatal(err)
		}
		rules[num] = matches[line_re.SubexpIndex("pattern")]

	}

	expanded := "^" + ExpandRule(rules, 0) + "$"
	rule_re := regexp.MustCompile(expanded)
	rule_re.Longest()
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		match := rule_re.FindStringIndex(line)
		if len(match) > 0 {
			if line[match[0]:match[1]] == line {
				total++
			}
		}
	}
	fmt.Println(total)
}
