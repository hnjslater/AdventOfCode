package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

var input = flag.String("input", "input12.txt", "")
var part = flag.Int("part", 1, "")

type Pair struct {
	b1 byte
	b2 byte
}

type CacheKey struct {
	p    Pair
	step int
}

type Submarine struct {
	Rules map[Pair]byte
	Cache map[CacheKey]map[byte]int
}

func NewSubmarine() Submarine {
	s := Submarine{}
	s.Rules = make(map[Pair]byte)
	s.Cache = make(map[CacheKey]map[byte]int)
	return s
}

func (s Submarine) Eval(step int, b1 byte, b2 byte) map[byte]int {
	if step == 0 {
		return map[byte]int{}
	}
	r, ok := s.Cache[CacheKey{Pair{b1, b2}, step}]
	if ok {
		return r
	}


	result := make(map[byte]int)
	mid := s.Rules[Pair{b1, b2}]
	result[mid]++
	for k, v := range s.Eval(step-1, b1, mid) {
		result[k] += v
	}
	for k, v := range s.Eval(step-1, mid, b2) {
		result[k] += v
	}
	s.Cache[CacheKey{Pair{b1, b2}, step}] = result
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

	scanner.Scan()
	molecule := scanner.Text()
	scanner.Scan()

	s := NewSubmarine()

	for scanner.Scan() {
		line := scanner.Text()
		s.Rules[Pair{line[0], line[1]}] = line[6]
	}

	steps := 10
	if *part == 2 {
		steps = 40
	}

	counts := make(map[byte]int)
	for i := range molecule {
		counts[molecule[i]]++
	}

	for i := 0; i < len(molecule)-1; i++ {
		for k, v := range s.Eval(steps, molecule[i], molecule[i+1]) {
			counts[k] += v
		}
	}

	max := 0
	min := math.MaxInt

	for _, v := range counts {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(max - min)
}
