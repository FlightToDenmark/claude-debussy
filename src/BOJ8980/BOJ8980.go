package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in = bufio.NewScanner(os.Stdin)

type info struct{ from, to, box int }

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
	N, C, M := NextInt(), NextInt(), NextInt()
	var infos []info
	truck := make([]int, N)
	for m := 0; m < M; m++ {
		infos = append(infos, info{NextInt(), NextInt(), NextInt()})
	}
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].to < infos[j].to
	})
	answer := 0
	for _, i := range infos {
		take := 0
		for j := i.from; j < i.to; j++ {
			if truck[j] > take {
				take = truck[j]
			}
		}
		if C-take > i.box {
			take = i.box
		} else {
			take = C - take
		}
		answer += take
		for j := i.from; j < i.to; j++ {
			truck[j] += take
		}
	}
	fmt.Println(answer)
}
