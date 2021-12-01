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
	max_seat_id := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := BinarySearchForSeat(line[0:7], 'F', 'B', 127, 0)
		col := BinarySearchForSeat(line[7:10], 'L', 'R', 7, 0)
		seat_id := (row * 8) + col
		if seat_id > max_seat_id {
			max_seat_id = seat_id
		}
	}
	fmt.Println(max_seat_id)
}
