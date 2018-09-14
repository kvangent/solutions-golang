package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// [Metadata]
// Title: Wireless is the New Fiber
// URL: https://open.kattis.com/problems/newfiber
// Categories: graph, sorting
// Difficulty: 2.7

type Node struct {
	id    int
	cap   int
	edges []int
}

func parseLine(scanner *bufio.Scanner) (int, int) {
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	i, _ := strconv.Atoi(input[0])
	j, _ := strconv.Atoi(input[1])
	return i, j
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Scan input
	n, m := parseLine(scanner)
	nodes := make([]Node, n, n)
	for i := range nodes {
		nodes[i].id = i
	}
	// Count edges for each node
	for i := 0; i < m; i++ {
		n1, n2 := parseLine(scanner)
		nodes[n1].cap++
		nodes[n2].cap++
	}
	// Remove capacity from nodes until at required number of edges
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].cap > nodes[j].cap // descending
	})
	extraDegrees, changed := 2*(m-n+1), 0
	for extraDegrees > 0 {
		changed++
		rmv := nodes[0].cap - 1
		if rmv > extraDegrees {
			nodes[0].cap -= extraDegrees
			break // partial shrink
		}
		nodes[0].cap -= rmv
		extraDegrees -= rmv
		// Move the shrunk node to the end
		nodes = append(nodes[1:], nodes[0])
	}
	// Assemble Tree and print
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fprintln(out, changed)
	fmt.Fprintf(out, "%d %d\n", len(nodes), len(nodes)-1)
	nodeQ, nbrQ := nodes[:], nodes[1:]
	nodes[0].cap++ // add the initial connection
	for len(nodeQ) != 0 {
		// pop next node to be assigned edges
		var node *Node
		node, nodeQ = &nodeQ[0], nodeQ[1:]
		// Add Edges (minus ed
		for i := 1; i < node.cap; i++ {
			// pop next edge to attach
			var nbr *Node
			nbr, nbrQ = &nbrQ[0], nbrQ[1:]
			fmt.Fprintf(out, "%d %d\n", node.id, nbr.id)
		}
	}
}
