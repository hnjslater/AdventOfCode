package main

import (
    "bufio"
    "log"
    "os"
    "regexp"
    "fmt"
    "strconv"
    "strings"
)

func check_range(value string, min int, max int) bool {
    parsed_value, err := strconv.Atoi(value)
    if err != nil {
        return false
    }
    return parsed_value >= min && parsed_value <= max
}

func check_regex(value string, regex string) bool {
    re := regexp.MustCompile(regex)
    return re.FindString(value) != ""
}

func check_height(name string, value string) bool {
    if len(value) < 3 {
        return false
    }
    measure := value[:len(value)-2]
    if strings.HasSuffix(value, "cm") {
        return check_range(measure, 150, 193)
    } else if strings.HasSuffix(value, "in") {
        return check_range(measure, 59, 76)
    }
    return false
}

func is_valid(name string, value string) bool {
    switch name {
    case "byr":
        return check_range(value, 1920, 2002)
    case "iyr":
        return check_range(value, 2010, 2020)
    case "eyr":
        return check_range(value, 2020, 2030)
    case "hgt":
        return check_height(name, value)
    case "hcl":
        return check_regex(value, "^#[0-9a-f]{6}$")
    case "ecl":
        return check_regex(value, "^(amb|blu|brn|gry|grn|hzl|oth)$")
    case "pid":
        return check_regex(value, "^[0-9]{9}$")
    case "cid":
        return false
    }
    return false
}

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
                field_value := matches[2]
                if is_valid(field_name, field_value) {
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
