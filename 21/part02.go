package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var input_file = flag.String("input", "input.txt", "Input file")
var debug = flag.Bool("debug", false, "Debug")

func AKey(m map[string]bool) string {
	for k, _ := range m {
		return k
	}
	return ""
}

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
				for k, _ := range options {
					if !imap[k] {
						delete(options, k)
					}
				}
			}
		}
	}
	a_to_i2 := make(map[string]string)

	changed := true
	for changed {
		changed = false
		a := ""
		i := ""
	inner:
		for k, v := range a_to_i {
			if len(v) == 1 {
				a = k
				i = AKey(v)
				break inner
			}
		}
		if a != "" {
			a_to_i2[a] = i
			for _, v := range a_to_i {
				delete(v, i)
			}
			changed = true
		}
	}

	var allegens []string
	for k, _ := range a_to_i2 {
		allegens = append(allegens, k)
	}

	sort.Strings(allegens)

	fmt.Printf("%v", a_to_i2[allegens[0]])
	for i := 1; i < len(allegens); i++ {
		fmt.Printf(",%v", a_to_i2[allegens[i]])
	}
	fmt.Println()
}
