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
var rule_re = regexp.MustCompile("(?P<name>[a-z ]+): (?P<min0>[0-9]+)-(?P<max0>[0-9]+) or (?P<min1>[0-9]+)-(?P<max1>[0-9]+)")

type Rule struct {
    name string
    ranges []int
    possibles map[int]bool
}

func main() {
	flag.Parse()
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
        var rules []Rule
	for scanner.Scan() {
            line := scanner.Text()
            if line == "" {
                break
            }
            matches := rule_re.FindStringSubmatch(line)
            name := matches[rule_re.SubexpIndex("name")]
            min0,_ := strconv.Atoi(matches[rule_re.SubexpIndex("min0")])
            max0,_ := strconv.Atoi(matches[rule_re.SubexpIndex("max0")])
            min1,_ := strconv.Atoi(matches[rule_re.SubexpIndex("min1")])
            max1,_ := strconv.Atoi(matches[rule_re.SubexpIndex("max1")])
            possibles := make(map[int]bool)
            rules = append(rules, Rule{name, []int{min0,max0,min1,max1}, possibles})
	}

        for _,r := range rules {
            for i:=0; i < len(rules); i++ {
                r.possibles[i] = true
            }
        }

        scanner.Scan()
        scanner.Scan()
        your_ticket := strings.Split(scanner.Text(), ",")
        scanner.Scan()
        scanner.Scan()
        for scanner.Scan() {
            line := scanner.Text()
            matched := true
            for _,vs := range strings.Split(line,",") {
                v,_ := strconv.Atoi(vs)
                field_matched := false
                for _,r := range rules {
                    if (v >= r.ranges[0] && v <= r.ranges[1]) || ( v >= r.ranges[2] && v <= r.ranges[3]) {
                        field_matched = true
                    }
                }
                if !field_matched {
                    matched = false
                }
            }
            if matched {
                for i,vs := range strings.Split(line,",") {
                    v,_ := strconv.Atoi(vs)
                    for _,r := range rules {
                        if v < r.ranges[0] || (v > r.ranges[1] && v < r.ranges[2]) || v > r.ranges[3] {
                            delete(r.possibles, i)
                        }
                    }
                }
            }
        }
        rmap := make(map[int]int)
        for len(rmap) < len(rules) {
            for i,r := range rules {
                if len(r.possibles) == 1 {
                    var k int
                    for k,_ = range r.possibles {
                    }
                    for _,r := range rules {
                        delete (r.possibles, k)
                    }
                    rmap[i] = k
                }
            }
        }
        total := 1
        for i,r := range rules {
            if strings.HasPrefix(r.name, "departure") {
                val,_ := strconv.Atoi(your_ticket[rmap[i]])
                total *= val
            }
        }
        fmt.Println(total)
}
