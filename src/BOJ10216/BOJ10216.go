package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)
var x, y, r, parent [3001]int

func main() {
	in.Split(bufio.ScanWords)
	T := nextInt()
	for t := 0; t < T; t++ {
		N := nextInt()
		for n := 1; n <= N; n++ {
			x[n], y[n], r[n], parent[n] = nextInt(), nextInt(), nextInt(), n
		}
		G := N
		for a := 1; a <= N; a++ {
			for b := a + 1; b <= N; b++ {
				dx := x[b] - x[a]
				dy := y[b] - y[a]
				if dx*dx+dy*dy <= (r[b]+r[a])*(r[b]+r[a]) {
					if find(a) != find(b) {
						union(a, b)
						G--
					}
				}
			}
		}
		fmt.Println(G)
	}
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x == y {
		return
	}
	parent[y] = x
}

func find(x int) int {
	for x != parent[x] {
		x = parent[x]
	}
	return x
}

func nextInt() int {
	in.Scan()
	r := 0
	for _, c := range in.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}
