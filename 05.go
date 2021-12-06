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

var input = flag.String("input", "test05.txt", "")
var part = flag.Int("part", 0, "")

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	var fish []int
	for _,entry := range strings.Split(line, ",") {
		x,_ := strconv.Atoi(entry)
		fish = append(fish, x)
	}
	fmt.Println(fish)
	if *part == 1 {
		for g := 0; g < 256; g++ {
			for i := 0; i < len(fish); i++ {
				fish[i]-= 1
				if fish[i] == -1 {
				   fish[i] = 6
				   fish = append(fish, 9)
			   }
			}
		}
	} else {
		var fishies [10]int
		for _,ent := range fish {
			fishies[ent]++
		}

		for g := 0; g < 256; g++ {
			breeders := fishies[0]
			for i := 1; i <= 9; i++ {
				fishies[i-1] = fishies[i]
			}
			fishies[8] = breeders
			fishies[6] += breeders
		}
		pop := 0
		for _,e := range fishies {
			pop += e
		}
		fmt.Println(pop)
	}

}
