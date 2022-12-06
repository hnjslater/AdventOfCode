package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var input = flag.String("input", "input06.txt", "")
var part = flag.Int("part", 0, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := 0
	length := 4
	if *part != 0 {
		length = 14
	}

	scanner.Scan()
	line := scanner.Text()
	for i := length; i < len(line); i++ {
		slot := line[i-length : i]
		m := make(map[rune]bool)
		for _, r := range slot {
			m[r] = true
		}
		if len(m) == length {
			result = i
			break
		}
	}
	fmt.Println(result)
}
