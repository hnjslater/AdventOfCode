package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var input = flag.String("input", "input12.txt", "")
var part = flag.Int("part", 1, "")

func contains(needle string, haystack []string) bool {
	for _, s := range haystack {
		if needle == s {
			return true
		}
	}
	return false
}

func walk(caves map[string][]string, path []string, cave string, small_cave_allowed bool) int {
	if cave == "end" {
		return 1
	}
	if !unicode.IsUpper(rune(cave[0])) {
		if contains(cave, path) {
			if !small_cave_allowed {
				return 0
			} else {
				small_cave_allowed = false
			}
		}
	}
	result := 0
	next_path := append(path, cave)

	for _, next_cave := range caves[cave] {
		result += walk(caves, next_path, next_cave, small_cave_allowed)
	}
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
	caves := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		vertexes := strings.Split(line, "-")

		if vertexes[0] != "end" && vertexes[1] != "start" {
			caves[vertexes[0]] = append(caves[vertexes[0]], vertexes[1])
		}
		if vertexes[1] != "end" && vertexes[0] != "start" {
			caves[vertexes[1]] = append(caves[vertexes[1]], vertexes[0])
		}
	}
	fmt.Println(walk(caves, []string{}, "start", *part == 2))
}
