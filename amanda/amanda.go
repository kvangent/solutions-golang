package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// [Metadata]
// Title: Amanda Lounges
// URL: https://open.kattis.com/problems/amanda
// Categories: graph, coloring
// Difficulty: 5.6

type Airport struct {
	isSet bool
	hasLounge bool
	nbrs  map[int]bool
}

var airports []Airport

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// airports, routes
	n, m := nextInt(scanner), nextInt(scanner)
	airports = make([]Airport, n)
	for i := range airports {
		airports[i].nbrs = make(map[int]bool)
	}
	edges := make([][]int, 0)
	for i := 0; i < m; i++ {
		a, b := nextInt(scanner)-1, nextInt(scanner)-1
		req := nextInt(scanner)
		if req == 1 {
			airports[a].nbrs[b] = true
			airports[b].nbrs[a] = true
		} else { // Save edge to setLounge later
			edges = append(edges, []int{a, b, req})
		}
	}
	lounges, valid := setRequired(edges)
	for i := 0; valid && i < n; i++ {
		if !airports[i].isSet {
			// Color the lounges, and count the smallest group
			c1, c2, v := setLounge(&airports[i], true)
			if c1 < c2 {
				lounges += c1
			} else {
				lounges += c2
			}
			valid = v && valid
		}
	}
	if valid {
		fmt.Println(lounges)
	} else {
		fmt.Println("impossible")
	}
}

func setRequired(edges [][]int) (lounges int, valid bool) {
	for _, edge := range edges {
		for i := 0; i <= 1; i++ {
			ct, _, valid := setLounge(&airports[edge[i]], edge[2] == 2)
			if !valid {
				return 0, false
			}
			lounges += ct
		}
	}
	return lounges,true
}

func setLounge(airport *Airport, state bool) (a int, b int, valid bool) {
	if airport.isSet {
		return 0,0, airport.hasLounge == state
	}
	airport.isSet, airport.hasLounge = true, state
	if state {
		a++
	} else {
		b++
	}
	// Set neighbors to opposite value
	for nbr := range airport.nbrs {
		aCt, bCt, valid := setLounge(&airports[nbr], !state)
		if !valid {
			return 0, 0, false
		}
		a += aCt
		b += bCt
	}
	return a, b, true
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
