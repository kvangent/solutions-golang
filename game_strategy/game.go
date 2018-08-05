package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

// [Metadata]
// Title: Game Strategy
// URL: https://open.kattis.com/problems/game
// Categories: minmax, dp
// Difficulty: 4.7

var positions [][]bitset // pos -> sets -> destinations
var results [][]int      // goal -> start
const inf = 26           // max is 25 edges, so 26 is equivalently infinite

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	n, _ := strconv.Atoi(next(scanner))
	positions = make([][]bitset, n)
	results = make([][]int, n)
	// Parse input
	for i := range positions {
		// Each position has several sets
		m, _ := strconv.Atoi(next(scanner))
		positions[i] = make([]bitset, m)
		for j := range positions[i] {
			// Each set has several potential destinations
			set := next(scanner)
			for _, r := range set {
				positions[i][j].add(int(r - 'a'))
			}
		}
	}
	// See which nodes can be forces to goal node
	for goal := range results {
		results[goal] = backtrack(goal)
	}
	// Print results
	for start := 0; start < n; start++ {
		for goal := 0; goal < n; goal++ {
			fmt.Printf("%d ", results[goal][start])
		}
		fmt.Println()
	}
}

func backtrack(goal int) []int {
	var visited bitset // track the nodes that if we reach, we win
	dist := make([]int, len(results)) // rounds to each node
	for i := range dist { dist[i] = -1 }
	// The goal node is always a win
	visited.add(goal)
	dist[goal] = 0
	for {
		var next bitset
		for i, choices := range positions {
			// only evaluate nodes we haven't reached yet
			if !visited.contains(i) {
				best := inf
				for _, set := range choices {
					// if a choice exists that only leads to nodes that force to the goal
					// we can add that node to our set of current winners
					if visited.containsAll(set) {
						next.add(i)
						v := maxFromSet(set, dist) + 1 // bob always picks the highest destination
						if v < best { best = v } // alice always picks the set with the lowest best
					}
				}
				if best < inf { dist[i] = best }
			}
		}
		if next == 0 { break } // quit if no progress
		visited.addAll(next)
	}
	return dist
}

type bitset int
func (b *bitset) add(idx int) { *b |= 1<<uint(idx) }
func (b *bitset) addAll(o bitset) { *b |= o }
func (b bitset) contains(idx int) bool { return b | 1<<uint(idx) == b }
func (b bitset) containsAll(o bitset) bool { return b | o == b }

// Returns the biggest value from an array where in index is in the bitset
func maxFromSet(b bitset, vals []int) int {
	max := -1
	for i, dist := range vals {
		if b.contains(i) && dist > max {
			max = dist
		}
	}
	return max
}

func next(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
