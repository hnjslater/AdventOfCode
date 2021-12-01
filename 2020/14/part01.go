package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var input_file = flag.String("input", "input.txt", "Input file")
var write_re = regexp.MustCompile("mem\\[(?P<address>[0-9]+)\\] = (?P<value>[0-9]+)")

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	memory := make(map[int]uint64)
	var mask uint64
	var value uint64
	for scanner.Scan() {
		line := scanner.Text()
		if line[1] == 'a' {
			mask_str := line[7:]
			mask_str = strings.ReplaceAll(mask_str, "1", "0")
			mask_str = strings.ReplaceAll(mask_str, "X", "1")
			mask, err = strconv.ParseUint(mask_str, 2, 64)
			if err != nil {
				log.Fatal(err)
			}

			value_str := line[7:]
			value_str = strings.ReplaceAll(value_str, "X", "0")
			value, err = strconv.ParseUint(value_str, 2, 64)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			matches := write_re.FindStringSubmatch(line)
			address, err := strconv.Atoi(matches[write_re.SubexpIndex("address")])
			if err != nil {
				log.Fatal(err)
			}
			wval, err := strconv.ParseUint(matches[write_re.SubexpIndex("value")], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			memory[address] = (wval & mask) | value
		}
	}

	var total uint64 = 0
	for _, v := range memory {
		total += v
	}
	fmt.Println(total)
}
