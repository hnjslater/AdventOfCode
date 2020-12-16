package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
        "strings"
)

var input_file = flag.String("input", "input.txt", "Input file")
var rule_re = regexp.MustCompile("(?P<field_name>[a-z ]+): (?P<min0>[0-9]+)-(?P<max0>[0-9]+) or (?P<min1>[0-9]+)-(?P<max1>[0-9]+)")

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
        var rules [][]int
	for scanner.Scan() {
            line := scanner.Text()
            if line == "" {
                break
            }
            matches := rule_re.FindStringSubmatch(line)
            min0,_ := strconv.Atoi(matches[rule_re.SubexpIndex("min0")])
            max0,_ := strconv.Atoi(matches[rule_re.SubexpIndex("max0")])
            min1,_ := strconv.Atoi(matches[rule_re.SubexpIndex("min1")])
            max1,_ := strconv.Atoi(matches[rule_re.SubexpIndex("max1")])
            
            rules = append(rules, []int{min0,max0})
            rules = append(rules, []int{min1,max1})
	}

        scanner.Scan()
        scanner.Scan()
        //your_ticket_str := scanner.Text()
        scanner.Scan()
        total := 0
        for scanner.Scan() {
            line := scanner.Text()
            for _,vs := range strings.Split(line,",") {
                v,_ := strconv.Atoi(vs)
                matched := false
                for _,r := range rules {
                    if v >= r[0] && v <= r[1] {
                        matched = true
                    }
                }
                if !matched {
                    total += v
                }
            }
        }

	fmt.Println(total)
}
