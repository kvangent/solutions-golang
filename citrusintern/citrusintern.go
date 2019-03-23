package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// [Metadata]
// Title: The Citrus Intern
// URL: https://open.kattis.com/problems/citrusintern
// Categories: dp
// Difficulty: 5.8

const INF = int(1e18)

var scanner *bufio.Scanner
var members []member
var dp [2][]int

type member struct {
	isChild bool
	cost    int
	subrds  []int
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	n := nextInt()
	members = make([]member, n)
	dp[0], dp[1] = make([]int, n), make([]int, n)
	for i := range dp[0] {
		dp[0][i], dp[1][i] = -1, -1
	}

	for i := range members {
		members[i].cost = nextInt()                // cost to bribe
		members[i].subrds = make([]int, nextInt()) // subordinates
		for j := range members[i].subrds {
			id := nextInt()
			members[i].subrds[j] = id
			members[id].isChild = true
		}
	}

	root := findRoot()
	fmt.Printf("%d\n", solve(root, 0))
}

// Returns the ID of the root node.
func findRoot() int {
	for i := range members {
		if !members[i].isChild {
			return i
		}
	}
	panic("Root not found.")
}

// Find the min cost to bring subtree with root 'id'.
func solve(id int, req int) int {
	if dp[req][id] != -1 {
		return dp[req][id]
	}

	// Bribing this node is optional
	if req == 0 {
		// Bribe this node
		opt1 := solve(id, 1)

		// Bribe one of this nodes children
		opt2 := INF
		notReq := 0 // min cost of child trees
		for _, c := range members[id].subrds {
			notReq += solve(c, 0)
		}
		for _, c := range members[id].subrds {
			t := notReq - solve(c, 0) + solve(c, 1)
			if t < opt2 {
				opt2 = t
			}
		}

		// Take the min of the either options
		if opt1 < opt2 {
			dp[0][id] = opt1
		} else {
			dp[0][id] = opt2
		}
		return dp[0][id]
	} else { //
		cost := members[id].cost
		for _, c := range members[id].subrds {
			for _, gc := range members[c].subrds {
				cost += solve(gc, 0)
			}
		}
		dp[1][id] = cost
		return dp[1][id]
	}

}

// Returns the next token and converts it to an int
func nextInt() int {
	scanner.Scan()
	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return i
}
