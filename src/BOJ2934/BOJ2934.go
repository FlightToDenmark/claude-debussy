package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)
var N int
var size = 100000
var tree []int

func main() {
	in.Split(bufio.ScanWords)
	N = nextInt()
	tree = make([]int, size+1)
	a, b, A, B := 0, 0, 0, 0
	for n := 0; n < N; n++ {
		a = nextInt()
		b = nextInt()
		update(a+1, 1)
		update(b, -1)
		A = sum(a)
		B = sum(b)
		fmt.Println(A + B)
		update(a, -A)
		update(a+1, A)
		update(b, -B)
		update(b+1, B)
	}
}

func update(i, val int) {

	for i <= size {
		tree[i] += val
		i += (i & -i)
	}

}

func sum(i int) int {

	answer := 0
	for i > 0 {
		answer += tree[i]
		i -= (i & -i)
	}
	return answer

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
