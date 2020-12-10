package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
        "sort"
)

var input_file = flag.String("input", "input.txt", "Input file")

func CountSubtrees(start int, buffer []int) int {
    if len(buffer)-start == 1 {
        return 1
    }

    total := 0
    if buffer[start+1] - buffer[start] <= 3 {
        total += CountSubtrees(start+1, buffer)
    }
    if len(buffer)-start > 2 && buffer[start+2] - buffer[start] <= 3 {
        total += CountSubtrees(start+2, buffer)
    }
    if len(buffer)-start > 3 && buffer[start+3] - buffer[start] <= 3 {
        total += CountSubtrees(start+3, buffer)
    }
    return total
}

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var buffer []int
        buffer = append(buffer, 0)
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Error parsing input")
		}
                buffer = append(buffer, value)
	}
        sort.Ints(buffer)
        buffer = append(buffer, buffer[len(buffer)-1]+3)

        fmt.Println(CountSubtrees(0,buffer))


}
