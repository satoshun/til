package main

import "fmt"

const (
	N = 4
	U = 5
)

type item struct {
	w int
	v int
}

var items map[int]item

func init() {
	items = make(map[int]item)
	items[0] = item{w: 2, v: 3}
	items[1] = item{w: 2, v: 2}
	items[2] = item{w: 3, v: 4}
	items[3] = item{w: 2, v: 2}
}

func search(i, u int) int {
	if i == N-1 {
		return 0
	} else if u < items[i].w {
		return search(i+1, u)
	} else {
		r1 := search(i+1, u)
		r2 := search(i+1, u-items[i].w) + items[i].v
		if r1 >= r2 {
			return r1
		}
		return r2
	}
}

func main() {
	fmt.Println(search(0, U))
}
