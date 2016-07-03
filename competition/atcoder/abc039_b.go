package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	a := nextInt()
	math.Sqrt

	a = divide(a)
	a = divide(a)
	fmt.Println(a)
}

func divide(a int) int {
	if a == 1 {
		return 1
	}
	if a == -1 {
		return -1
	}

	b := a / 2
	for {
		t := b * b
		if a == t {
			return b
		}

		if t > a {
			b -= 1
		} else {
			b += 1
		}
	}
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
