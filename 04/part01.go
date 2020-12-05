package main

import (
    "bufio"
    "log"
    "os"
    "regexp"
    "fmt"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    re := regexp.MustCompile("(?P<fieldname>[a-z]+):(?P<value>[a-z0-9#]+)")
    found_fields := make(map[string]bool)
    valid_passports := 0
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            if len(found_fields) == 7 {
                valid_passports++
            }
            found_fields = make(map[string]bool)
        } else {
            matches := re.FindAllStringSubmatch(line, -1)
            for _, matches := range matches {
                field_name := matches[1]
                if field_name != "cid" {
                    found_fields[field_name] = true
                }
            }
        }
    }
    if len(found_fields) == 7 {
        valid_passports++
    }
    fmt.Println(valid_passports)
}
