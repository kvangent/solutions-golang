package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// [Metadata]
// Title: Red Black Tree
// URL: https://open.kattis.com/problems/redblacktree
// Categories: dp, recursive
// Difficulty: 3.0

const MOD = 1000000007

type Node struct {
	red bool
	children []int
}

var nodes []Node

// solve returns array where array[i] is the number of subsets with i red nodes
func solve(par Node) []int {
	// Empty leafs
	if len(par.children) == 0 {
		if par.red {
			return []int{0, 1}
		} else {
			return []int{1}
		}
	}
	childCts := make([][]int, 0, len(par.children))
	maxRed := 0
	// Calculate counts for subtrees
	for _, child := range par.children {
		ct := solve(nodes[child])
		maxRed += len(ct)-1
		childCts = append(childCts, ct)
	}
	// Combine subtree counts
	curCt := make([]int, maxRed + 1)
	redCt := 0
	for _, ct := range childCts {
		newCts := make([]int, maxRed + 1)
		copy(newCts, curCt)
		// Use i nodes from the current tree
		for i := 0; i < len(ct); i++ {
			// Using just this subtree
			newCts[i] += ct[i]
			newCts[i] %= MOD
			// Use j red nodes from trees so far
			for j:= 0; j <= redCt; j++ {
				newCts[i+j] += ct[i] * curCt[j]
				newCts[i+j] %= MOD
			}
		}
		redCt += len(ct)-1
		curCt = newCts

	}
	// Count the node itself
	if par.red {
		if len(curCt) < 2 {
			curCt = []int{curCt[0], 0}
		}
		curCt[1] = (curCt[1] + 1) % MOD
	} else {
		curCt[0] = (curCt[0] + 1) % MOD
	}
	return curCt
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	n := nextInt(scanner) // number of nodes
	m := nextInt(scanner) // number of red nodes

	// Parse input
	nodes = make([]Node, n)
	for i := 1; i < n; i++ {
		p := nextInt(scanner) - 1
		nodes[p].children = append(nodes[p].children, i)
	}
	for i := 0; i < m; i++ {
		id := nextInt(scanner) - 1
		nodes[id].red = true
	}
	counts := solve(nodes[0])
	// Empty set
	counts[0]++
	counts[0] %= MOD
	for _, ct := range counts {
		fmt.Println(ct)
	}
	for i := len(counts); i < m+1; i++ {
		fmt.Println(0)
	}
}

func nextInt(scanner *bufio.Scanner) int{
	scanner.Scan()
	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return i
}
