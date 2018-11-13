package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in = bufio.NewScanner(os.Stdin)
var budget int
var A []int

func main() {
	in.Split(bufio.ScanWords)
	N := nextInt()
	A = make([]int, N)
	sum := 0
	for i := range A {
		A[i] = nextInt()
		sum += A[i]
	}
	budget = nextInt()
	sort.Ints(A)
	if budget >= sum {
		fmt.Println(A[N-1])
		os.Exit(0)
	}
	left, right := 1, A[N-1]
	mid := 0
	for left < right {
		mid = (left + right) / 2
		if possible(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	fmt.Println(left - 1)

}

func possible(x int) bool {
	sum := 0
	for _, a := range A {
		sum += min(x, a)
	}
	if sum <= budget {
		return true
	}
	return false
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

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
