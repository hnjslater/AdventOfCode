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

var input_file = flag.String("input", "input.txt", "Input file")

// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func reverse(s string) string {
	b := make([]byte, len(s))
	var j int = len(s) - 1
	for i := 0; i <= j; i++ {
		b[j-i] = s[i]
	}

	return string(b)
}

func east(tile []string) string {
	str := make([]byte, len(tile))
	for i, s := range tile {
		str[i] = s[len(s)-1]
	}
	return string(str)
}
func west(tile []string) string {
	str := make([]byte, len(tile))
	for i, s := range tile {
		str[i] = s[0]
	}
	return string(str)
}

func north(tile []string) string {
	return tile[0]
}

func south(tile []string) string {
	return tile[len(tile)-1]
}

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	tiles := make(map[int][]string)
	edges := make(map[string][]int)
	tile := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Tile") {
			tile, err = strconv.Atoi(line[5:9])
			if err != nil {
				log.Fatal(err)
			}
		} else if line != "" {
			tiles[tile] = append(tiles[tile], line)
		}
	}

	for k, v := range tiles {
		edges[north(v)] = append(edges[north(v)], k)
		edges[east(v)] = append(edges[east(v)], k)
		edges[south(v)] = append(edges[south(v)], k)
		edges[west(v)] = append(edges[west(v)], k)

		edges[reverse(north(v))] = append(edges[reverse(north(v))], k)
		edges[reverse(east(v))] = append(edges[reverse(east(v))], k)
		edges[reverse(south(v))] = append(edges[reverse(south(v))], k)
		edges[reverse(west(v))] = append(edges[reverse(west(v))], k)

		edges[""] = append(edges[""], k)
	}

	counts := make(map[int]int)
	for _, v := range edges {
		if len(v) == 1 {
			counts[v[0]]++
		}
	}
	var corners []int
	for k, v := range counts {
		if v == 4 {
			corners = append(corners, k)
		}
	}
	total := 1
	for _, c := range corners {
		total = total * c
	}
	fmt.Println(total)
}
