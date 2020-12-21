package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
        "strings"
        "strconv"
)
var input_file = flag.String("input", "input.txt", "Input file")

// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func Reverse( s string ) string {
    b := make([]byte, len(s));
    var j int = len(s) - 1;
    for i := 0; i <= j; i++ {
        b[j-i] = s[i]
    }

    return string ( b );
}

func East(tile []string) string {
    str := make([]byte, len(tile))
    for i,s := range tile {
        str[i] = s[len(s)-1]
    }
    return string(str)
}
func West(tile []string) string {
    str := make([]byte, len(tile))
    for i,s := range tile {
        str[i] = s[0]
    }
    return string(str)
}

func North(tile []string) string {
    return tile[0];
}

func South(tile []string) string {
    return tile[len(tile)-1]
}

func main() {
    flag.Parse()
    file, err := os.Open(*input_file)
    if err != nil {
            log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)

    tiles := make(map[int][]string)
    edges := make(map[string][]int)
    tile := 0

    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "Tile") {
            tile, err = strconv.Atoi(line[5:9])
            if err != nil {
                log.Fatal(err)
            }
        } else if line != "" {
            tiles[tile] = append(tiles[tile], line)
        }
    }

    for k,v := range tiles {
        edges[North(v)] = append(edges[North(v)], k)
        edges[East(v)] = append(edges[East(v)], k)
        edges[South(v)] = append(edges[South(v)], k)
        edges[West(v)] = append(edges[West(v)], k)
        edges[Reverse(North(v))] = append(edges[Reverse(North(v))], k)
        edges[Reverse(East(v))] = append(edges[Reverse(East(v))], k)
        edges[Reverse(South(v))] = append(edges[Reverse(South(v))], k)
        edges[Reverse(West(v))] = append(edges[Reverse(West(v))], k)
        edges[""] = append(edges[""], k)
    }

    counts := make(map[int]int)
    for _,v := range edges {
        if len(v) == 1 {
            counts[v[0]]++
        }
    }
    var corners []int
    for k,v := range counts {
        if v == 4 {
            corners = append(corners, k)
        }
    }
    total := 1
    for _,c := range corners {
        total = total * c
    }
    fmt.Println(total)
}
