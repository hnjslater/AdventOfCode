package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	l := list.New()
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		l.PushBack(value)
	}

	for it := l.Front(); it != nil; it = it.Next() {
		for jt := it.Next(); jt != nil; jt = jt.Next() {
			if it.Value.(int)+jt.Value.(int) == 2020 {
				fmt.Println(it.Value.(int) * jt.Value.(int))
			}
		}
	}
}
