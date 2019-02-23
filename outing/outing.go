package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// [Metadata]
// Title: Outing
// URL: https://open.kattis.com/problems/outing
// Categories: graph, DAG, dfs
// Difficulty: 5.9

type Node struct {
	reqs  *Node
	reqBy []*Node
	visit int
}

var nodes []Node

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// participants, bus size
	n, k := nextInt(scanner), nextInt(scanner)
	nodes = make([]Node, n)
	for i := range nodes {
		req := nextInt(scanner) - 1
		nodes[i].reqs = &nodes[req]
		nodes[req].reqBy = append(nodes[req].reqBy, &nodes[i])
	}
	// Group participants with requirements with a min and max size
	var grps [][2]int
	for _, node := range nodes {
		if node.visit == 0 {
			// If the node hasn't been visited, it in a new group
			min, max := groupSizes(&node)
			grps = append(grps, [2]int{min, max})
		}
	}

	// Every possible passenger ct to be reached for each potential group
	possible := make([][]bool, len(grps)+1) // group, passengers
	for i := range possible {
		possible[i] = make([]bool, k+1)
	}
	possible[0][0] = true // No passengers is always possible
	for g := 0; g < len(grps); g++ { // for each group to potentially add
		for p := 0; p <= k; p++ { // num of passengers possible on the bus
			if !possible[g][p] {
				continue
			}
			possible[g+1][p] = true // take no one
			// Try taking every potential group size
			for i := grps[g][0]; i <= grps[g][1] && p+i < len(possible[g]); i++ {
				possible[g+1][p+i] = true
			}
		}
	}
	// Print the highest number of passengers possible
	for p := k; p >= 0; p-- {
		if possible[len(grps)][p] {
			fmt.Println(p)
			break
		}
	}
}

// groupSizes returns the min and max a graph component can be
func groupSizes(cur *Node) (int, int) {
	depth := 1
	for ; cur.visit == 0; depth++ {
		cur.visit = depth
		cur = cur.reqs
	}
	min := depth - cur.visit
	return min, groupMax(cur)
}

// Returns the max size of a graph component
func groupMax(cur *Node) int {
	if cur.visit == -1 {
		return 0
	}
	cur.visit = -1
	max := 1
	for _, nbr := range cur.reqBy {
		max += groupMax(nbr)
	}
	return max
}

// Returns the next token and converts it to an int
func nextInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return i
}
