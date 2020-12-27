package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	valid_count := 0
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("(?P<min>[0-9]+)-(?P<max>[0-9]+) (?P<char>[a-z]): (?P<pass>[a-z]+)+")
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		min, _ := strconv.Atoi(matches[re.SubexpIndex("min")])
		max, _ := strconv.Atoi(matches[re.SubexpIndex("max")])
		char := matches[re.SubexpIndex("char")]
		pass := matches[re.SubexpIndex("pass")]

		count := strings.Count(pass, char)
		if count >= min && count <= max {
			valid_count++
		}

	}
	fmt.Println(valid_count)
}
