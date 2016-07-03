package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	l := nextLine()
	s.Scan()
	c, _ := strconv.Atoi(s.Text())
	p := [...]int{0, 0}

	hatena := 0
	for _, s := range l {
		switch s {
		case 'L':
			p[0] -= 1
		case 'R':
			p[0] += 1
		case 'U':
			p[1] += 1
		case 'D':
			p[1] -= 1
		case '?':
			hatena += 1
		}
	}

	if c == 1 {
		// max
		if p[0] >= 0 {
			p[0] += hatena
		} else {
			p[0] -= hatena
		}
	} else {
		// min
		for i := 0; i < hatena; i++ {
			if p[0] > 0 {
				p[0] -= 1
			} else if p[0] < 0 {
				p[0] += 1
			} else if p[1] > 0 {
				p[1] -= 1
			} else {
				p[1] += 1
			}
		}
	}

	fmt.Println(math.Abs(float64(p[0])) + math.Abs(float64(p[1])))
}

var s = bufio.NewScanner(os.Stdin)

func next() string {
	s.Split(bufio.ScanWords)
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
