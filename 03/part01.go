package main

import (
    "bufio"
    "log"
    "os"
    "fmt"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    x := 0
    y := 0
    width := 31
    trees := 0
    for scanner.Scan() {
        line := scanner.Text()
        if line[x] == '#' {
            trees++
        }
        x=(x+3)%width
        y++
    }
    fmt.Println(trees)
}
