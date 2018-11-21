package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)

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
	N := NextInt()
	var r []int
	for n := 0; n < N; n++ {
		r = append(r, NextInt())
	}
	var stack []int
	answer := 0
	for n := 0; n < N; n++ {
		for len(stack) > 0 && r[stack[len(stack)-1]] > r[n] {
			height := r[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := n
			if len(stack) > 0 {
				width -= stack[len(stack)-1] + 1
			}
			if answer < width*height {
				answer = width * height
			}
		}
		stack = append(stack, n)
	}
	for len(stack) > 0 {
		height := r[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		width := N
		if len(stack) > 0 {
			width -= stack[len(stack)-1] + 1
		}
		if answer < width*height {
			answer = width * height
		}
	}
	fmt.Println(answer)
}
