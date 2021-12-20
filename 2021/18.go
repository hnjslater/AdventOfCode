// THIS CODE IS HORRIBLE AND I FEEL BAD ABOUT IT

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var input = flag.String("input", "test18.txt", "")
var part = flag.Int("part", 1, "")
var test = flag.Bool("test", false, "")

type Pair struct {
	leftp  *Pair
	rightp *Pair
	leftv  int
	rightv int
}

type Stack []*Pair

func (s Stack) Top() *Pair {
	return s[len(s)-1]
}

func (p *Pair) ProcessExplode(s Stack) bool {
	if len(s) == 4 {
		nums := append(s, p)[0].Collect()
		for i := 1; i < len(nums); i++ {
			if nums[i] == &p.leftv {
				*nums[i-1] += p.leftv
				break
			}
		}
		nums = append(s, p)[0].Collect()
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] == &p.rightv {
				*nums[i+1] += p.rightv
				break
			}
		}
		if s.Top().leftp == p {
			s.Top().leftp = nil
			s.Top().leftv = 0
		} else {
			s.Top().rightp = nil
			s.Top().rightv = 0
		}
		return true
	}

	if p.leftp != nil {
		changed := p.leftp.ProcessExplode(append(s, p))
		if changed {
			return true
		}
	}
	if p.rightp != nil {
		changed := p.rightp.ProcessExplode(append(s, p))
		if changed {
			return true
		}
	}
	return false
}

func (p *Pair) ProcessSplit(s Stack) bool {
	if p.leftp != nil {
		changed := p.leftp.ProcessSplit(append(s, p))
		if changed {
			return true
		}
	} else if p.leftv > 9 {
		left := p.leftv / 2
		right := p.leftv/2 + p.leftv%2
		p.leftv = -1
		p.leftp = &Pair{leftv: left, rightv: right}
		return true
	}

	if p.rightp != nil {
		changed := p.rightp.ProcessSplit(append(s, p))
		if changed {
			return true
		}
	} else if p.rightv > 9 {
		left := p.rightv / 2
		right := p.rightv/2 + p.rightv%2
		p.rightv = -1
		p.rightp = &Pair{leftv: left, rightv: right}
		return true
	}

	return false
}

func (p *Pair) Collect() []*int {
	var ret []*int

	if p.leftp == nil {
		ret = append(ret, &p.leftv)
	} else {
		ret = append(ret, p.leftp.Collect()...)
	}

	if p.rightp == nil {
		ret = append(ret, &p.rightv)
	} else {
		ret = append(ret, p.rightp.Collect()...)
	}

	return ret
}

func (p Pair) String(sin ...string) string {
	sout := ""
	if len(sin) > 0 {
		sout += sin[0]
	}
	sout += "["
	if p.leftp != nil {
		sout += p.leftp.String()
	} else {
		sout += strconv.Itoa(p.leftv)
	}
	sout += ","
	if p.rightp != nil {
		sout += p.rightp.String()
	} else {
		sout += strconv.Itoa(p.rightv)
	}
	sout += "]"

	return sout
}

func (p1 Pair) Add(p2 Pair) *Pair {
	p := Pair{leftv: -1, rightv: -1, leftp: &p1, rightp: &p2}
	changed := true
	for changed {
		changed = p.ProcessExplode(Stack{})
		if !changed {
			changed = p.ProcessSplit(Stack{})
		}
	}
	return &p
}

func (p Pair) Mag() int {
	left := 0
	right := 0

	if p.leftp != nil {
		left = p.leftp.Mag()
	} else {
		left = p.leftv
	}
	if p.rightp != nil {
		right = p.rightp.Mag()
	} else {
		right = p.rightv
	}

	return 3*left + 2*right
}

func MakePair(s string) *Pair {
	var outer *Pair
	var stack Stack
	for _, r := range s {
		switch r {
		case ('['):
			p := Pair{leftv: -1, rightv: -1}
			if outer == nil {
				outer = &p
			}
			if len(stack) > 0 {
				if stack.Top().leftp == nil && stack.Top().leftv == -1 {
					stack.Top().leftp = &p
				} else {
					stack.Top().rightp = &p
				}
			}
			stack = append(stack, &p)
		case (']'):
			if len(stack) > 1 {
				stack = stack[0 : len(stack)-1]
			}
		case (','):
		default:
			v, _ := strconv.Atoi(string(r))
			if stack.Top().leftv == -1 && stack.Top().leftp == nil {
				stack.Top().leftv = v
			} else {
				stack.Top().rightv = v
			}
		}
	}
	return outer
}

func Test(in string, expected string) {
	p := MakePair(in)
	p.ProcessExplode(Stack{})
	result := p.String()
	if result != expected {
		fmt.Println("      In: " + in)
		fmt.Println("     Out: " + result)
		fmt.Println("Expected: " + expected)
		fmt.Println()
	} else {
		fmt.Println("Success")
	}
}

func TestAdd(op1 string, op2 string, expected string) {
	p1 := MakePair(op1)
	p2 := MakePair(op2)
	result := p1.Add(*p2).String()

	if result != expected {
		fmt.Println("      In: "+op1, "+", op2)
		fmt.Println("     Out: " + result)
		fmt.Println("Expected: " + expected)
		fmt.Println()
	} else {
		fmt.Println("success")
	}
}

func main() {
	flag.Parse()

	if *test {
		Test("[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]")
		Test("[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]")
		Test("[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]")
		Test("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
		Test("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]")

		TestAdd("[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
		TestAdd("[[[[1,1],[2,2]],[3,3]],[4,4]]", "[5,5]", "[[[[3,0],[5,3]],[4,4]],[5,5]]")
		TestAdd("[[[[3,0],[5,3]],[4,4]],[5,5]]", "[6,6]", "[[[[5,0],[7,4]],[5,5]],[6,6]]")
		TestAdd("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]", "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]", "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]")
		TestAdd("[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]", "[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]", "[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]")
	} else {
		file, err := os.Open(*input)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		if *part == 1 {
			scanner := bufio.NewScanner(file)
			var tot *Pair
			for scanner.Scan() {
				text := scanner.Text()
				if tot == nil {
					tot = MakePair(text)
				} else {
					tot = tot.Add(*MakePair(text))
				}
			}

			fmt.Println(tot.Mag())
		} else {
			scanner := bufio.NewScanner(file)
			var pairs []string
			for scanner.Scan() {
				pairs = append(pairs, scanner.Text())
			}
			result := 0
			for i, p1s := range pairs {
				for j, p2s := range pairs {
					if i != j {
						p1 := MakePair(p1s)
						p2 := MakePair(p2s)
						sum := p1.Add(*p2)
						if sum.Mag() > result {
							result = sum.Mag()
						}
					}
				}
			}
			fmt.Println(result)
		}
	}
}
