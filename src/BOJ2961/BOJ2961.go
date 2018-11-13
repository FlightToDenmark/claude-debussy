package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var in = bufio.NewScanner(os.Stdin)
var answer, S, B int
var A [][]int

func main() {
	in.Split(bufio.ScanWords)
	N := nextInt()
	A = make([][]int, N)
	for n := 0; n < N; n++ {
		A[n] = make([]int, 2)
		A[n][0], A[n][1] = nextInt(), nextInt()
	}
	answer = math.MaxInt64
	for s := 1; s < 1<<uint(N); s++ {
		S, B = 1, 0
		for n := 0; n < N; n++ {
			if s&(1<<uint(n)) > 0 {
				S *= (A[n][0])
				B += (A[n][1])
			}
		}
		answer = min(answer, abs(S-B))
	}
	fmt.Println(answer)
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

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
