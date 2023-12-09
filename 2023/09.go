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

var input = flag.String("input", "09.txt", "")
var part = flag.Int("part", 0, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		var ns []int

		for _, s := range line {
			v, _ := strconv.Atoi(s)
			ns = append(ns, v)
		}
		var ns2 []int
		all_equal := false
		level := 0
		var fronts []int
		for {
			fronts = append(fronts, ns[0])
			if *part == 0 {
				total += ns[len(ns)-1]
			}
			all_equal = true
			for _, n := range ns {
				if n != ns[0] {
					all_equal = false
				}
			}
			if all_equal == true {
				break
			}
			for i := 0; i < len(ns)-1; i++ {
				ns2 = append(ns2, (ns[i+1] - ns[i]))
			}
			ns = ns2
			ns2 = make([]int, 0)
			level++
		}
		if *part > 0 {
			x := 0
			for i := len(fronts) - 1; i >= 0; i-- {
				x = fronts[i] - x
			}
			total += x
		}
	}
	fmt.Println(total)
}
