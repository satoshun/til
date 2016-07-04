package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := nextLine()
	ss := strings.Split(s, " ")
	a, _ := strconv.Atoi(ss[0])
	b, _ := strconv.Atoi(ss[1])
	c, _ := strconv.Atoi(ss[2])
	d, _ := strconv.Atoi(ss[3])
	e, _ := strconv.Atoi(ss[4])

	if e+d+a >= e+c+b {
		fmt.Println(e + d + a)
	} else {
		fmt.Println(e + c + b)
	}
}

var s = bufio.NewScanner(os.Stdin)

func next() string {
	s.Split(bufio.ScanWords)
	s.Scan()
	return s.Text()
}

func nextValue() string {
	s.Scan()
	return s.Text()
}

func nextLine() string {
	s.Split(bufio.ScanLines)
	s.Scan()
	return s.Text()
}
