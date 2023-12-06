package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var input = flag.String("input", "05.txt", "")
var part = flag.Int("part", 0, "")

type map_entry struct {
	source int
	dest   int
	length int
}

func (me map_entry) end() int {
	return me.source + me.length - 1
}

func (me map_entry) Map(s seed) seed {
	diff := me.dest - me.source
	return seed{s.start + diff, s.end + diff}
}

type seed struct {
	start int
	end   int
}

func process(maps [][]map_entry, s seed, stage int) int {
	if stage == len(maps) {
		return s.start
	}

	min_final := math.MaxInt32
	todo := []seed{s}
	undoable := []seed{}
	for len(todo) > 0 {
		s := todo[0]
		todo = todo[1:]
		matched := false
		for _, entry := range maps[stage] {
			r := math.MaxInt32
			if s.start >= entry.source && s.start <= entry.end() && s.end >= entry.source && s.end <= entry.end() {
				r = process(maps, entry.Map(s), stage+1)
				matched = true
			} else if s.start >= entry.source && s.start <= entry.end() {
				r = process(maps, entry.Map(seed{s.start, entry.end()}), stage+1)
				todo = append(todo, seed{entry.end() + 1, s.end})
				matched = true
			} else if s.end >= entry.source && s.end <= entry.end() {
				r = process(maps, entry.Map(seed{entry.source, s.end}), stage+1)
				todo = append(todo, seed{s.start, entry.source - 1})
				matched = true
			} else if s.start < entry.source && s.end > entry.end() {
				todo = append(todo, seed{s.start, entry.source - 1})
				todo = append(todo, seed{entry.end() + 1, s.end})
				r = process(maps, seed{entry.source, entry.end()}, stage+1)
				matched = true
			}
			if r < min_final {
				min_final = r
			}
		}
		if !matched {
			undoable = append(undoable, s)
		}
	}

	for _, s := range undoable {
		r := process(maps, s, stage+1)
		if r < min_final {
			min_final = r
		}
	}

	return min_final
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
	seeds_raw := strings.Split(scanner.Text(), " ")[1:]

	var seeds []seed
	if *part == 0 {
		seeds = make([]seed, len(seeds_raw))
		for i := 0; i < len(seeds_raw); i++ {
			seed_number, _ := strconv.Atoi(seeds_raw[i])
			seeds[i] = seed{seed_number, seed_number}
		}
	} else {
		seeds = make([]seed, len(seeds_raw)/2)
		for i := 0; i < len(seeds_raw); i += 2 {
			seed_number, _ := strconv.Atoi(seeds_raw[i])
			length, _ := strconv.Atoi(seeds_raw[i+1])
			seeds[i/2] = seed{seed_number, seed_number + length - 1}
		}
	}

	var maps [][]map_entry
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if line[len(line)-1] == ':' {
			maps = append(maps, make([]map_entry, 0))
			continue
		}
		entries := strings.Split(line, " ")
		prev, _ := strconv.Atoi(entries[0])
		next, _ := strconv.Atoi(entries[1])
		size, _ := strconv.Atoi(entries[2])

		maps[len(maps)-1] = append(maps[len(maps)-1], map_entry{next, prev, size})

	}

	min_final := math.MaxInt32

	for _, seed := range seeds {
		val := process(maps, seed, 0)
		if val < min_final {
			min_final = val
		}
	}

	fmt.Println(min_final)
}
