package main

import (
    "bufio"
    "log"
    "os"
    "regexp"
    "fmt"
    "flag"
    "strings"
    "strconv"
)

var input_file = flag.String("input", "input.txt", "Input file")

type BagContents struct {
    desc string
    count int
}

var line_re = regexp.MustCompile("(?P<bag>[a-z]+ [a-z]+) bags contain (?P<contents>[\\w\\s,]*)\\.")
var desc_re = regexp.MustCompile("(?P<count>[0-9]+) (?P<bag>[a-z]+ [a-z]+) (?:bags|bag)")

func ResolveBags(bags map[string][]BagContents, desc string) int {
    var total = 0
    for _,content := range bags[desc] { 
        if content.count > 0 {
            total += content.count + content.count * ResolveBags(bags, content.desc)
        }
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

    bags := make(map[string][]BagContents)

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        matches := line_re.FindStringSubmatch(line)
        if len(matches) == 0 {
            panic("Regex didn't match line")
        } else {
            bag := matches[line_re.SubexpIndex("bag")]
            contents := matches[line_re.SubexpIndex("contents")]
            if contents == "no other bags" {
                bags[bag] = nil
            } else {
                for _,content_desc := range strings.Split(contents, ",") {
                    content_desc = strings.TrimSpace(content_desc)
                    dmatches := desc_re.FindStringSubmatch(content_desc)
                    count,err := strconv.Atoi(dmatches[desc_re.SubexpIndex("count")])
                    if err != nil {
                        panic("error parsing file.")
                    }
                    desc := dmatches[desc_re.SubexpIndex("bag")]
                    bags[bag] = append(bags[bag], BagContents{desc: desc, count: count})
                }
            }
        }
    }

    fmt.Println(ResolveBags(bags, "shiny gold"))
}
