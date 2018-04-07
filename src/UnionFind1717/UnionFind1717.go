package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var N int
var M int
var parent []int
var rank []int

func main() {

	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &M)

	parent = make([]int, N+1)
	for i, _ := range parent {
		parent[i] = i
	}
	rank = make([]int, N+1)
	for i, _ := range rank {
		rank[i] = 1
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for m := 0; m < M; m++ {
		var n []int
		scanner.Scan()
		for _, f := range strings.Fields(scanner.Text()) {
			if i, err := strconv.Atoi(f); err == nil {
				n = append(n, i)
			}
		}

		if n[0] == 0 {
			union(n[1], n[2])
		} else {
			p1, p2 := find(n[1]), find(n[2])
			if p1 == p2 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}

}

func union(u int, v int) {

	u, v = find(u), find(v)

	if u == v {
		return
	}

	if rank[u] > rank[v] {
		u, v = v, u
	}

	parent[u] = v
	if rank[u] == rank[v] {
		rank[v]++
	}

}

func find(u int) int {
	
	for u != parent[u] {
		u = parent[u]	
	}
	
	return u

}
