package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var input_file = flag.String("input", "input.txt", "Input file")
var debug = flag.Bool("debug", false, "Debug")

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	a_to_i := make(map[string]map[string]bool)
	all_ingts := make(map[string]bool)
	all_algns := make(map[string]bool)

	var uses []string

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line[0:len(line)-1], "(contains ")

		ingts := strings.Split(parts[0], " ")
		algns := strings.Split(parts[1], ", ")
		ingts = ingts[0 : len(ingts)-1]

		for _, algn := range algns {
			all_algns[algn] = true
		}
		for _, ingt := range ingts {
			all_ingts[ingt] = true
			uses = append(uses, ingt)
		}

		for _, algn := range algns {
			imap := make(map[string]bool)
			for _, ingt := range ingts {
				imap[ingt] = true
			}
			options, found := a_to_i[algn]
			if !found {
				a_to_i[algn] = imap
			} else {
				for k := range options {
					if !imap[k] {
						delete(options, k)
					}
				}
			}
		}
	}

	clean_ingts := make(map[string]bool)
	for k := range all_ingts {
		clean_ingts[k] = true
	}

	for _, v := range a_to_i {
		for k2 := range v {
			delete(clean_ingts, k2)
		}
	}

	count := 0

	for _, u := range uses {
		if clean_ingts[u] {
			count++
		}
	}

	fmt.Println(count)
}
