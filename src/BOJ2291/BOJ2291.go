package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in.Split(bufio.ScanWords)
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
