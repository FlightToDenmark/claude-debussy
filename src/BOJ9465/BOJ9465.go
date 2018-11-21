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
	T := NextInt()
	var s [200000]int
	var cache [200000][2]int
	for t := 0; t < T; t++ {
		N := NextInt()
		for n := 0; n < 2*N; n += 2 {
			s[n] = NextInt()
			cache[n][0], cache[n][1] = -1, -1
		}
		for n := 1; n < 2*N; n += 2 {
			s[n] = NextInt()
			cache[n][0], cache[n][1] = -1, -1
		}
		var solve func(i, nextClose int) int
		solve = func(i, nextClose int) int {
			ret := 0
			if i >= 2*N {
				return 0
			}
			if cache[i][nextClose] != -1 {
				return cache[i][nextClose]
			}
			if i%2 == 0 {
				ret = s[i] + solve(i+3, 0)
				if nextClose == 0 {
					ret = max(ret, solve(i+1, 0))
				} else {
					ret = max(ret, solve(i+2, 0))
				}
			} else {
				ret = max(s[i]+solve(i+1, 1), solve(i+1, 0))
			}
			cache[i][nextClose] = ret
			return ret
		}
		fmt.Println(solve(0, 0))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
