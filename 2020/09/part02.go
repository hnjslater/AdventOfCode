package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var input_file = flag.String("input", "input.txt", "Input file")

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int
	value := 0
	needle := 57195069
	for scanner.Scan() {
		line := scanner.Text()
		value, err = strconv.Atoi(line)
		if err != nil {
			log.Fatal("Error parsing input")
		}
		numbers = append(numbers, value)
	}
	i := 0
	j := 0
	for ; i < len(numbers) && value != needle; i++ {
		j = i
		value = 0
		for value < needle {
			value += numbers[j]
			j++
		}
	}
	smallest := numbers[i]
	largest := numbers[i]

	for k := i; k < j; k++ {
		if numbers[k] > smallest {
			smallest = numbers[k]
		}
		if numbers[k] < largest {
			largest = numbers[k]
		}
	}
	fmt.Println(largest + smallest)
}
