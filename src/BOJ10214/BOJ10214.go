package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)

func main() {
	in.Split(bufio.ScanWords)
	T := nextInt()
	for t := 0; t < T; t++ {
		Y, K := 0, 0
		for i := 0; i < 9; i++ {
			Y += nextInt()
			K += nextInt()
		}
		if Y > K {
			fmt.Println("Yonsei")
		} else if K > Y {
			fmt.Println("Korea")
		} else {
			fmt.Println("Draw")
		}
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
