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
	var N, S int
	fmt.Scan(&N)
	fmt.Scan(&S)
	var num []int
	for n := 0; n < N; n++ {
		var x int
		fmt.Scan(&x)
		num = append(num, x)
	}
	answer := 0
	for subset := 1; subset < (1 << uint32(N)); subset++ {
		sum := 0
		for n := 0; n < N; n++ {
			if subset&(1<<uint32(n)) > 0 {
				sum += num[n]
			}
		}
		if sum == S {
			answer++
		}
	}
	fmt.Println(answer)
}
