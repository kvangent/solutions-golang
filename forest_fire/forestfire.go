package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// [Metadata]
// Title: Forest Fires
// URL: https://open.kattis.com/problems/forestfires
// Categories: data structures
// Difficulty: 6.2

var r int
var dirs = [4][2]int{{0,1},{1,0},{0,-1},{-1,0}}

func main(){
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for in.Scan() {
		input := strings.Split(in.Text(), " ")
		r, _ = strconv.Atoi(input[0])
		n, _ := strconv.Atoi(input[1])
		grid := make([][]bool, 100)
		for c := range grid {
			grid[c] = make([]bool, 100)
		}
		uf := NewUnionFind(10000)

		trees := make([]int, n)
		var fireCt int
		for i := 0; i < n; i++ {
			// Find an empty space for a tree
			var m, x, y int
			for ok := true; ok; ok = grid[x][y] {
				m = nextR() % 10000
				x, y = m/100, m%100
			}
			// Place tree
			grid[x][y] = true
			trees[i] = m

			// If any neighbors are trees, join the sets
			for _, d := range dirs {
				nX, nY := x+d[0], y+d[1]
				if nX >= 0 && nX < 100 && nY >= 0 && nY < 100 && grid[nX][nY] {
					uf.join((100*nX)+nY, m)
				}
			}

			// Check 'random' A and B
			a := trees[nextR() % (i+1)]
			b := trees[nextR() % (i+1)]
			if uf.find(a) == uf.find(b) {
				fireCt++
			}

			// Print and reset fire queries every 100
			if i % 100 == 99 {
				fmt.Fprintf(out, "%d ", fireCt)
				fireCt = 0
			}
		}
		fmt.Fprintln(out)
	}
}

func nextR() int{
	r = (r * 5171 + 13297) % 50021
	return r
}

type UnionFind struct{
	parent []int
}

func NewUnionFind(size int) *UnionFind{
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent}
}

func (uf *UnionFind) find(x int) int{
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) join(x, y int) {
	uf.parent[uf.find(x)] = uf.find(y)
}
