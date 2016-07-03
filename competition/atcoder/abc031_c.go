package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var memo map[int]int

func main() {
	memo = make(map[int]int)
	n, _ := strconv.Atoi(nextLine())
	a := strings.Split(nextLine2(), " ")

	aa := make([]int, 0, len(a))
	for _, bb := range a {
		bbb, _ := strconv.Atoi(bb)
		aa = append(aa, bbb)
	}

	score := 0
	for i := 0; i <= n; i++ {
		s := search(i, aa)
		if s > score {
			score = s
		}
	}

	fmt.Println(score)
}

func search(p int, a []int) int {
	s1 := -1
	s2 := -1
	for i := 0; i <= len(a); i++ {
		if i == p {
			continue
		}

		var p1, p2 int
		if i > p {
			p1, p2 = calc(a[p:i])
		} else {
			p1, p2 = calc(a[i:p])
		}

		if s2 == -1 || p2 > s2 {
			s1 = p1
			s2 = p2
		}
	}

	return s1
}

func calc(a []int) (int, int) {
	p1, p2 := 0, 0
	for i := range a {
		if i%2 == 0 { // even
			p1 += a[i]
		} else { // odd
			p2 += a[i]
		}
	}

	return p1, p2
}

var s = bufio.NewScanner(os.Stdin)

func next() string {
	s.Split(bufio.ScanWords)
	s.Scan()
	return s.Text()
}

func nextLine2() string {
	s.Scan()
	return s.Text()
}

func nextLine() string {
	s.Split(bufio.ScanLines)
	s.Scan()
	return s.Text()
}

func nextInt() int {
	i, e := strconv.Atoi(next())
	if e != nil {
		panic(e)
	}
	return i
}
