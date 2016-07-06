package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nextLine()
	ns := strings.Split(nextLine2(), " ")
	total := 0
	i := 1
	for j, n := range ns {
		total += i
		if len(ns)-1 == j {
			break
		}
		if n < ns[j+1] {
			i++
		} else {
			i = 1
		}
	}
	fmt.Println(total)
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

func nextInt2() int {
	s.Scan()
	i, e := strconv.Atoi(s.Text())
	if e != nil {
		panic(e)
	}
	return i
}
