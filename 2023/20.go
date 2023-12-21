package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var input = flag.String("input", "20.txt", "")
var part = flag.Int("part", 0, "")

type module struct {
	mtype           byte
	outputs         []string
	flipflop_memory bool
	inputs          map[string]bool
}

type pulse struct {
	source      string
	destination string
	high        bool
}

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	network := make(map[string]*module)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		name := line[0]
		mtype := byte('b')
		if name != "broadcaster" {
			name = line[0][1:]
			mtype = line[0][0]
		}
		outputs := strings.Split(line[1], ", ")

		network[name] = &module{mtype, outputs, false, make(map[string]bool)}
	}

	for name, source_module := range network {
		for _, output := range source_module.outputs {
			_, found := network[output]
			if !found {
				network[output] = &module{'b', []string{}, false, make(map[string]bool)}
			}
			network[output].inputs[name] = false
		}
	}

	low_pulses := 0
	high_pulses := 0
	var first_high = map[string]int{"ft": 1, "ng": 1, "sv": 1, "jz": 1}
	count_found := 0
	stop := 1000
	if *part > 0 {
		stop = 10000
	}
	for i := 0; i < stop; i++ {
		pulses := []pulse{{"button", "broadcaster", false}}
		for len(pulses) > 0 {
			curr_pulse := pulses[0]
			pulses = pulses[1:]

			if *part > 0 {
				if val, found := first_high[curr_pulse.destination]; found && !curr_pulse.high {
					if val == 1 {
						fmt.Println(i)
						first_high[curr_pulse.destination] = 2
						count_found++
					}
				}
				if count_found == 4 {
					break
				}
				// I should do an LCM calculation here but I didn't because I hated this one
			}

			if curr_pulse.high {
				high_pulses++
			} else {
				low_pulses++
			}
			var next_high bool
			curr_module := network[curr_pulse.destination]
			switch curr_module.mtype {
			case ('b'):
				next_high = curr_pulse.high
			case ('%'):
				if !curr_pulse.high {
					if curr_module.flipflop_memory == false {
						next_high = true
						curr_module.flipflop_memory = true
					} else {
						next_high = false
						curr_module.flipflop_memory = false
					}
				} else {
					continue
				}
			case ('&'):
				curr_module.inputs[curr_pulse.source] = curr_pulse.high
				all_high := true
				for _, v := range curr_module.inputs {
					all_high = all_high && v
				}
				if all_high {
					next_high = false
				} else {
					next_high = true
				}
			}

			for _, output := range curr_module.outputs {
				pulses = append(pulses, pulse{curr_pulse.destination, output, next_high})
			}
		}
	}
	if *part == 0 {
		fmt.Println(high_pulses * low_pulses)
	}
}
