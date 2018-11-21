package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	for n := N; n >= 1; n-- {
		L := strings.Repeat("*", n)
		fmt.Println(L)
	}
}
