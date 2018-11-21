package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)
var visited [25]bool
var connected [5][5]bool
var input []byte
var answer int
var before int
var connect int

const S, Y = 83, 89

func main() {
	in.Split(bufio.ScanWords)
	for i := 0; i < 5; i++ {
		in.Scan()
		for _, c := range in.Bytes() {
			input = append(input, c)
		}
	}
	answer = 0
	before = -1
	find(0)
	fmt.Println(answer)
}

func memset() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			connected[i][j] = false
		}
	}
}

func rangeCheck(i, j int) bool {
	if i < 0 || j < 0 || i >= 5 || j >= 5 {
		return false
	}
	return true

}

func find(total int) {

	if total == 7 {
		connect = 1
		memset()
		for i := 0; i < 25; i++ {
			if visited[i] {
				connected[i%5][i/5] = true
				isConnected(i)
				break
			}
		}
	}

	for i := before + 1; i < 25; i++ {
		visited[i] = true
		before = i
		find(total + 1)
		visited[i] = false
	}

}

func isConnected(index int) {

	if connect == 7 {
		ss := 0
		for i := 0; i < 25; i++ {
			if visited[i] && input[i] == S {
				ss++
			}
		}
		if ss >= 4 {
			answer++
		}
	} else {
		i := index % 5
		j := index / 5
		if rangeCheck(i+1, j) && !connected[i+1][j] && visited[index+1] {
			connected[i+1][j] = true
			connect++
			isConnected(index + 1)
		}
		if rangeCheck(i-1, j) && !connected[i-1][j] && visited[index-1] {
			connected[i-1][j] = true
			connect++
			isConnected(index - 1)
		}
		if rangeCheck(i, j+1) && !connected[i][j+1] && visited[index+5] {
			connected[i][j+1] = true
			connect++
			isConnected(index + 5)
		}
		if rangeCheck(i, j-1) && !connected[i][j-1] && visited[index-5] {
			connected[i][j-1] = true
			connect++
			isConnected(index - 5)
		}
	}
}
