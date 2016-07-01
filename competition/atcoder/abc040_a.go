package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var s = bufio.NewScanner(os.Stdin)

func main() {
	a := strings.Split(nextLine(), " ")

	n, _ := strconv.Atoi(a[0])
	x, _ := strconv.Atoi(a[1])

	a1 := x - 1
	a2 := n - x

	if a1 >= a2 {
		fmt.Println(a2)
	} else {
		fmt.Println(a1)
	}
}

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
