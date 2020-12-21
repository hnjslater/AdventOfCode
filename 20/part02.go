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
var debug = flag.Bool("debug", false, "Debug")

// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func Reverse( s string ) string {
    b := make([]byte, len(s));
    var j int = len(s) - 1;
    for i := 0; i <= j; i++ {
        b[j-i] = s[i]
    }

    return string ( b );
}

func Addr(x int, y int) int {
    return (x*12)+y
}

func rAddr(a int) (int, int) {
    return a / 12, a % 12
}

func Rotate(tile []string) []string {
    tile1 := make([]string, len(tile))

    for x := 0; x < len(tile); x++ {
        rs := make([]byte, len(tile))
        for y:=0; y < len(tile); y++ {
            rs[y] = tile[y][len(tile) - x -1]
        }
        tile1[x] = string(rs)
    }

    return tile1
}

func Flip(tile []string) []string {
    tile1 := make([]string, len(tile))
    for i := 0; i < len(tile); i++ {
        tile1[len(tile)-i-1] = tile[i]
    }
    return tile1
}

func Mirror(tile []string) []string {
    var tile1 []string
    for _,s := range(tile) {
        tile1 = append(tile1, Reverse(s))
    }
    return tile1
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

func PrintTile(s []string) {
    for _,s := range s {
        fmt.Println(s)
    }
    fmt.Println()
}

func PrintTiles(s [][]string) {
    for row:=0; row < 10; row++ {
        for tnum :=0; tnum < len(s); tnum++ {
            fmt.Print(s[tnum][row], " ")
        }
    fmt.Println()
    }
}
type Tile struct {
    num int
    value []string
}

type Entry struct {
    num int
    rotate int
    flip bool
    mirror bool
    edge int
}

func TryTile(tiles map[int][]string, edges map[string][]int, used []Tile, used2 map[int]bool) ([]Tile, bool) {
    x, y := rAddr(len(used))
    if len(tiles) == len(used2) {
        return used, true
    }

    var matches []Tile
    var e string
    var s string
    c := ""

    if x > 0 {
        s = South(used[Addr(x-1,y)].value)
        c = s
    }
    if y > 0 {
        e = East(used[Addr(x,y-1)].value)
        c = e
    }
    loop: for _,k := range edges[c] {
        if used2[k] {
            continue loop
        }
        for i := 0; i < 4; i++ {
            flip := false
            mirror := false
            if i == 0 {
            } else if i == 1 {
                flip = true
            } else if i == 2 {
                mirror = true
            } else if i == 3 {
                flip = true
                mirror = true
            }
            for j := 0; j < 4; j++ {
                ok := true
                other_tile := make([]string, len(tiles[k]))
                copy(other_tile, tiles[k])

                if flip {
                    other_tile = Flip(other_tile)
                }
                if mirror {
                    other_tile = Mirror(other_tile)
                }
                for r:=0; r < j; r++ {
                    other_tile = Rotate(other_tile)
                }
                if len(e) > 0 {
                    ok = ok && (e == West(other_tile))
                }
                if len(s) > 0 {
                    ok = ok && (s == North(other_tile))
                }
                if ok {
                    matches = append(matches, Tile{k,other_tile})
                }
            }
        }
    }
    for _,m := range matches {
        used2[m.num] = true
        r, ok := TryTile(tiles, edges, append(used, m), used2)
        delete(used2, m.num)
        if ok {
            return r, ok
        }
    }
    return used, false
}

func Search(needle []string, haystack []string) (int) {
    found := 0
    for r:=0; r< len(haystack)-len(needle); r++ {
        placeloop: for c:=0; c < len(haystack[0])-len(needle[0]); c++ {
            for i,nr := range needle {
                for j,n := range nr {
                    if n == '#' && haystack[r+i][c+j] == '.' {
                        continue placeloop
                    }
                }
            }
            for i,nr := range needle {
                row := []byte(haystack[r+i])
                for j,n := range nr {
                    if n == '#' {
                        row[c+j] = '.'
                    }
                }
                haystack[r+i] = string(row)
            }
            found++
        }
    }
    return found
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
        edges[South(v)] = append(edges[South(v)], k)
        edges[East(v)] = append(edges[East(v)], k)
        edges[West(v)] = append(edges[West(v)], k)
        edges[Reverse(North(v))] = append(edges[Reverse(North(v))], k)
        edges[Reverse(South(v))] = append(edges[Reverse(South(v))], k)
        edges[Reverse(East(v))] = append(edges[Reverse(East(v))], k)
        edges[Reverse(West(v))] = append(edges[Reverse(West(v))], k)
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

    edges[""] = append(edges[""], corners[0])

    layout, ok := TryTile(tiles, edges, make([]Tile,0), make(map[int]bool))
    if ! ok {
        log.Fatal("No result found")
    }

    var topt [][]string
    for i,t := range layout {
        if i % 12 == 0 && i != 0 {
            topt = make([][]string,0)
        }
        topt = append(topt, t.value)
    }

    var smap []string

    for row := 0; row < 120; row++ {
        smap_row := ""
        if row % 10 == 0 || row % 10 == 9 {
            continue
        }
        x := row / 10
        for col := 0; col < 12; col++ {
            trow := layout[Addr(x, col)].value[row%10]
            smap_row += trow[1:len(trow)-1]
        }
        smap = append(smap, smap_row)
    }

    needle := []string{"                  # ",
                       "#    ##    ##    ###",
                       " #  #  #  #  #  #   "}
    outer: for k := 0; k < 4; k++ {
        smap2 := make([]string, len(smap))
        copy(smap2, smap)
        if k == 1 || k == 3 {
            smap2 = Flip(smap2)
        }
        if k == 2 || k == 3 {
            smap2 = Mirror(smap2)
        }

        for i := 0; i < 4; i++ {
            found := Search(needle, smap2)
            if found > 0 {
                count := 0
                for _,v := range smap2 {
                    for _,r := range v {
                        if r == '#' {
                            count++
                        }
                    }
                }
                fmt.Println(count)
                break outer
            }
            smap2 = Rotate(smap2)
        }
    }
}
