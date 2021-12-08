package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var input = flag.String("input", "test08.txt", "")
var part = flag.Int("part", 1, "")

func sorted(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}

func contains(needle byte, haystack string) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}

func stringAnd(values ...string) string {
	result := ""

	for b := byte('a'); b <= byte('g'); b++ {
		count := 0
		for _, v := range values {
			if contains(b, v) {
				count++
			}
		}
		if count == len(values) {
			result += string(rune(b))
		}
	}
	return result
}

func determineDigit(unknown *map[int][]string, length int, pattern string, common int) string {
	for i, poss := range (*unknown)[length] {
		r := stringAnd(pattern, poss)
		if len(r) == common {
			(*unknown)[length] = append((*unknown)[length][:i], (*unknown)[length][i+1:]...)
			return poss
		}
	}
	panic("No Match")
}
func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input := strings.Split(line, " | ")

		digits := make(map[int]string)
		unknown := make(map[int][]string)

		for _, d := range strings.Split(input[0], " ") {
			switch len(d) {
			case 2:
				digits[1] = d
			case 3:
				digits[7] = d
			case 4:
				digits[4] = d
			case 7:
				digits[8] = d
			default:
				unknown[len(d)] = append(unknown[len(d)], d)
			}
		}

		digits[6] = determineDigit(&unknown, 6, digits[1], 1)
		digits[0] = determineDigit(&unknown, 6, stringAnd(unknown[5]...), 2)
		digits[9] = unknown[6][0]
		digits[3] = determineDigit(&unknown, 5, digits[1], 2)
		digits[5] = determineDigit(&unknown, 5, digits[6], 5)
		digits[2] = unknown[5][0]

		rdigits := make(map[string]int)
		for k, v := range digits {
			rdigits[sorted(v)] = k
		}
		cresult := 0
		for _, d := range strings.Split(input[1], " ") {
			val := rdigits[sorted(d)]
			if *part == 1 {
				if val == 1 || val == 4 || val == 7 || val == 8 {
					result++
				}
			} else {
				cresult = cresult * 10
				cresult += val
			}
		}
		result += cresult
	}
	fmt.Println(result)
}
