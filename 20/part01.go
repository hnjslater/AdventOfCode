package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
        "strings"
        "strconv"
        "time"
        "runtime/pprof"
)
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
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
    return (x*3)+y
}

func rAddr(a int) (int, int) {
    return a / 3, a % 3
}

func Rotate(tile []string) []string {
    tile1 := make([]string, len(tile))

    for x := 0; x < len(tile); x++ {
        rs := make([]byte, 10)
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
var cache map[Entry]string

func GetEdge(tiles map[int][]string, num int, rotate int, flip bool, mirror bool, edge int) string {
    e,found := cache[Entry{num,rotate,flip,mirror,edge}]
    if found {
        return e
    }
    t := make([]string, len(tiles[num]))
    copy(t, tiles[num])

    if flip {
        t = Flip(t)
    }
    if mirror {
        t = Mirror(t)
    }
    for i:=0; i < rotate; i++ {
        t = Rotate(t)
    }
    result := ""
    if edge == 0 {
        result = North(t)
    } else if edge == 1 {
        result = East(t)
    } else if edge == 2 {
        result = South(t)
    } else if edge == 3 {
        result = West(t)
    }
    cache[Entry{num,rotate,flip,mirror,edge}] = result
    return result
}

func TryTile(tiles map[int][]string, edges map[string][]int, used []Tile) ([]Tile, bool) {
    x, y := rAddr(len(used))
    if len(used) == 1 {
        fmt.Print(time.Now().Format("2006-01-02 15:04:05"), " ")
        fmt.Println(used[0].num)
    }
    if len(tiles) == 0 {
        return used, true
    }

    var matches []Tile
    var e string
    var s string
    c := ""
    if y > 0 {
        e = East(used[len(used)-1].value)
    }
    if x > 0 {
        s = South(used[Addr(x-1,y)].value)
    }
    loop: for _,k := range edges[c] {
        _,found := tiles[k]
        if ! found {
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
                if len(e) > 0 {
                    ok = ok && (e == GetEdge(tiles, k, j, flip, mirror, 3))
                }
                if len(s) > 0 {
                    ok = ok && (s == GetEdge(tiles, k, j, flip, mirror, 0))
                }
                if ok {
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
                    matches = append(matches, Tile{k,other_tile})
                }
            }
        }
    }
    for _,m := range matches {
        old_tile := tiles[m.num]
        delete(tiles, m.num)
        r, ok := TryTile(tiles, edges, append(used, m))
        tiles[m.num] = old_tile
        if ok {
            return r, ok
        }
    }
    return used, false
}

func main() {
    cache = make(map[Entry]string)
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }

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
        edges[""] = append(edges[""], k)
    }

    var layout []Tile;
    layout2, ok := TryTile(tiles, edges, layout)
    if ! ok {
        log.Fatal("No result found")
    }
    for _,v := range layout2 {
        fmt.Print(v.num, " ")
    }
    fmt.Println()
    if *debug {
        fmt.Print(layout2[0].num*layout2[2].num*layout2[6].num*layout2[8].num)
    }
    fmt.Println()
}
