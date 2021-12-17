package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var input = flag.String("input", "test16.txt", "")
var part = flag.Int("part", 1, "")
var literal = flag.String("literal", "", "")

type PacketStream struct {
	s string
	c int
}

func NewPacketStream(s string) PacketStream {
	// a string of zeros and ones is the obvious way to encode this right?
	r := ""

	for i := range s {
		x, _ := strconv.ParseInt(string(s[i]), 16, 64)
		b := strconv.FormatInt(x, 2)
		for len(b) < 4 {
			b = "0" + b
		}
		r += b
	}
	return PacketStream{r, 0}
}

func (ps *PacketStream) getInt(length int) int {
	r, _ := strconv.ParseInt(string(ps.s[ps.c:ps.c+length]), 2, 64)

	ps.c += length
	return int(r)
}

func (ps *PacketStream) getvInt() int {
	val := 0
	for {
		conbit := ps.getInt(1)
		num := ps.getInt(4)
		val = (val * 16) + int(num)
		if conbit == 0 {
			break
		}
	}

	return val
}

func (ps *PacketStream) Packet() int {
	version := ps.getInt(3)
	typeid := ps.getInt(3)
	if typeid == 4 {
		value := ps.getvInt()
		if *part == 1 {
			return version
		} else {
			return value
		}
	} else {

		lentype := ps.getInt(1)
		var subpackets []int
		if lentype == 0 {
			plen := ps.getInt(15)
			end := ps.c + plen
			for ps.c < end {
				subpackets = append(subpackets, ps.Packet())
			}
		} else {
			pcount := ps.getInt(11)
			for i := 0; i < pcount; i++ {
				subpackets = append(subpackets, ps.Packet())
			}
		}
		result := 0
		if *part == 1 {
			typeid = 0
			result = version
		}

		switch typeid {
		case 0:
			for _, x := range subpackets {
				result += x
			}
		case 1:
			result = 1
			for _, x := range subpackets {
				result *= x
			}
		case 2:
			result = math.MaxInt
			for _, x := range subpackets {
				if x < result {
					result = x
				}
			}
		case 3:
			result = 0
			for _, x := range subpackets {
				if x > result {
					result = x
				}
			}
		case 5:
			if subpackets[0] > subpackets[1] {
				result = 1
			} else {
				result = 0
			}
		case 6:
			if subpackets[0] < subpackets[1] {
				result = 1
			} else {
				result = 0
			}
		case 7:
			if subpackets[0] == subpackets[1] {
				result = 1
			} else {
				result = 0
			}
		}
		return result
	}
}

func main() {
	flag.Parse()
	text := *literal
	if text == "" {
		file, err := os.Open(*input)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Scan()
		text = scanner.Text()
	}
	ps := NewPacketStream(text)
	result := ps.Packet()

	fmt.Println(result)
}
