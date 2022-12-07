package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type inode struct {
	size  int
	files map[string]*inode
}

func (n inode) Du() int {
	result := 0
	result += n.size
	for k, v := range n.files {
		if k != ".." {
			result += v.Du()
		}
	}
	return result
}

func newInode() *inode {
	var i inode
	i.files = make(map[string]*inode)
	return &i
}

func part1(n *inode) int {
	result := 0
	if n.size == 0 {
		if n.Du() < 100000 {
			result += n.Du()
		}
		for k, f := range n.files {
			if k != ".." {
				result += part1(f)
			}
		}
	}

	return result
}

func part2Imp(n *inode) []int {
	var sizes []int
	if n.size == 0 {
		sizes = append(sizes, n.Du())
		for k, v := range n.files {
			if k != ".." {
				sizes = append(sizes, part2Imp(v)...)
			}
		}
	}

	return sizes
}

func part2(root *inode) int {
	r := part2Imp(root)
	sort.Ints(r)
	unused := 70000000 - root.Du()

	for _, v := range r {
		if unused+v > 30000000 {
			return v
		}
	}

	return 0
}

var input = flag.String("input", "input07.txt", "")
var part = flag.Int("part", 0, "")

var cd = regexp.MustCompile(`\$ cd ([a-z.]+)`)
var fileListing = regexp.MustCompile(`([[:digit:]]+) (.+)`)

func main() {
	flag.Parse()
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	root := newInode()
	var current *inode
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "$ cd /" {
			current = root
		} else if line == "$ cd .." {
			current = current.files[".."]
		} else if match := cd.FindStringSubmatch(line); len(match) > 0 {
			inode := newInode()
			inode.files[".."] = current
			current.files[match[1]] = inode
			current = inode
		} else if match := fileListing.FindStringSubmatch(line); len(match) > 0 {
			size, _ := strconv.Atoi(match[1])
			inode := newInode()
			inode.size = size
			current.files[match[2]] = inode
		}
	}
	if *part == 0 {
		fmt.Println(part1(root))
	} else {
		fmt.Println(part2(root))
	}
}
