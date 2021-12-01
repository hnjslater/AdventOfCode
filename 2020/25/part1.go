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
var debug = flag.Bool("debug", false, "Debug")
var part = flag.Int("part", 1, "Part2?")

func transform(value int, subject int) int {
	value = value * subject
	value = value % 20201227
	return value
}

func getloop(key int, subject int) int {
	value := 1
	loop := 0
	for value != key {
		value = transform(value, subject)
		loop++
	}
	return loop
}

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	card_key, _ := strconv.Atoi(line)
	card_loop := getloop(card_key, 7)

	scanner.Scan()
	line = scanner.Text()
	door_key, _ := strconv.Atoi(line)
	door_loop := getloop(door_key, 7)

	x := 1
	for i := 0; i < card_loop; i++ {
		x = transform(x, door_key)
	}
	fmt.Println(x)

	y := 1
	for i := 0; i < door_loop; i++ {
		y = transform(y, card_key)
	}
	fmt.Println(y)
}
