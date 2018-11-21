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
	answer := 0
	g1, g2, g3, a := make([]int, N+1), make([]int, N+1), make([]int, N+1), make([]int, N+1)
	for i := 1; i <= N; i++ {
		g1[i] = NextInt()
	}
	for i := 1; i <= N; i++ {
		g2[i] = NextInt()
	}
	for i := 1; i <= N; i++ {
		g3[i] = NextInt()
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			a[j] = 0
		}
		for j := 1; j <= N; j++ {
			if g1[j] == i {
				break
			}
			a[g1[j]]++
		}
		for j := 1; j <= N; j++ {
			if g2[j] == i {
				break
			}
			a[g2[j]]++
		}
		for j := 1; j <= N; j++ {
			if g3[j] == i {
				break
			}
			a[g3[j]]++
		}
		great := true
		for j := 1; j <= N; j++ {
			if a[j] == 3 {
				great = false
				break
			}
		}
		if great {
			answer++
		}
	}
	fmt.Println(answer)
}
