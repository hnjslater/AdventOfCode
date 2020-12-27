package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func BinarySearchForSeat(input string, take_bottom rune, take_top rune, max int, min int) int {
	mid := 0
	for _, c := range input {
		mid = min + ((max - min) / 2)
		if c == take_top {
			min = mid
		} else if c == take_bottom {
			max = mid
		} else {
			panic("Unexpected character in seat string")
		}
	}
	return max
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var seats [128 * 8]bool
	for scanner.Scan() {
		line := scanner.Text()
		row := BinarySearchForSeat(line[0:7], 'F', 'B', 127, 0)
		col := BinarySearchForSeat(line[7:10], 'L', 'R', 7, 0)
		seat_id := (row * 8) + col
		seats[seat_id] = true
	}
	for i := 1; i < len(seats)-1; i++ {
		if seats[i-1] && !seats[i] && seats[i+1] {
			fmt.Println(i)
		}
	}
}
