package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	_, _ = strconv.Atoi(nextLine())
	ss := strings.Split(nextLine2(), " ")
	ii := make([]int, 0, len(ss))
	for _, iii := range ss {
		i, _ := strconv.Atoi(iii)
		ii = append(ii, i)
	}

	orig := make([]int, len(ii))
	copy(orig, ii)
	bubbleSort(ii)

	for _, i := range ii {
		for j, v := range orig {
			if v == i {
				fmt.Println(j + 1)
				break
			}
		}
	}
}

func bubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if s[j] > s[j-1] {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}
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
