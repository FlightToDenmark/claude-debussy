package main

import (
	"bufio"
	"fmt"
	"os"
)

var N, size, ret int
var tree, A []int
var in = bufio.NewScanner(os.Stdin)

func main() {
	in.Split(bufio.ScanWords)
	N = nextInt()
	size = 1
	for size < N {
		size *= 2
	}
	tree = make([]int, size*2)
	for i := 0; i < size; i++ {
		update(i, 1)
	}
	A = make([]int, N)
	for n := 1; n <= N; n++ {
		ret = query(1, nextInt())
		update(ret, 0)
		A[ret] = n
	}
	for n := 0; n < N; n++ {
		fmt.Println(A[n])
	}
}

func update(i, value int) {
	i += size
	tree[i] = value
	for i > 1 {
		i /= 2
		tree[i] = tree[2*i] + tree[2*i+1]
	}
}

func query(node, k int) int {
	if node >= size {
		return node - size
	}
	pivot := tree[2*node]
	if k < pivot {
		return query(2*node, k)
	} else {
		return query(2*node+1, k-pivot)
	}
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
