package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)
var rank [][]int
var tree []int
var size = 1

const INF = 99999999

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
	N, index := NextInt(), 0
	rank = make([][]int, 3)
	var firstOrder []int
	for i := 0; i < 3; i++ {
		rank[i] = make([]int, N+1)
		for r := 1; r <= N; r++ {
			index = NextInt()
			rank[i][index] = r
			if i == 0 {
				firstOrder = append(firstOrder, index)
			}
		}
	}
	treeInit(N)
	answer := 0
	for _, v := range firstOrder {
		if query(0, size, 1, 0, rank[1][v]) > rank[2][v] {
			answer++
		}
		update(rank[1][v], rank[2][v])
	}
	fmt.Println(answer)

}

func treeInit(n int) {

	for size < n {
		size *= 2
	}
	tree = make([]int, size*2)
	for i := range tree {
		tree[i] = INF
	}

}

func min(x, y int) int {

	if x > y {
		return y
	}
	return x

}

func update(i, value int) {

	i += size
	tree[i] = value
	for i > 1 {
		i /= 2
		tree[i] = min(tree[2*i], tree[2*i+1])
	}

}

func query(nodeL, nodeR, nodeIndex, queryL, queryR int) int {

	if queryR < nodeL || queryL > nodeR {
		return INF
	}
	if queryL <= nodeL && nodeR <= queryR {
		return tree[nodeIndex]
	}
	mid := (nodeL + nodeR) / 2
	return min(query(nodeL, mid, nodeIndex*2, queryL, queryR), query(mid+1, nodeR, nodeIndex*2+1, queryL, queryR))

}
