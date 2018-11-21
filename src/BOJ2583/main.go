package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in = bufio.NewScanner(os.Stdin)

type point struct{ x, y int }

func NextInt() int {
	in.Scan()
	r := 0
	for _, c := range in.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}

func main() {
	in.Split(bufio.ScanWords)
	m, n, k := NextInt(), NextInt(), NextInt()
	r := make(map[point]bool, m*n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			r[point{i, j}] = true
		}
	}
	for ; k > 0; k-- {
		x1, y1, x2, y2 := NextInt(), NextInt(), NextInt(), NextInt()
		for i := x1; i < x2; i++ {
			for j := y1; j < y2; j++ {
				delete(r, point{i, j})
			}
		}
	}
	var f func(p point) int
	f = func(p point) int {
		if !r[p] {
			return 0
		}
		delete(r, p)
		return 1 + f(point{p.x - 1, p.y}) + f(point{p.x + 1, p.y}) + f(point{p.x, p.y - 1}) + f(point{p.x, p.y + 1})
	}
	var c []int
	for len(r) > 0 {
		for p := range r {
			c = append(c, f(p))
			break
		}
	}
	sort.Ints(c)
	fmt.Println(len(c))
	for _, v := range c {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
