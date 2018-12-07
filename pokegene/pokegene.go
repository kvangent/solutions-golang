package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// [Metadata]
// Title: Pokegene
// URL: https://open.kattis.com/problems/pokegene
// Categories: strings
// Difficulty: 6.7

const MOD = 1000000007

type Monster struct {
	id   int
	gene string
	hash []uint
}

var mon []Monster

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 512*1024), 1024*1024)
	scanner.Split(bufio.ScanWords)

	n := nextInt(scanner) // monsters in db
	q := nextInt(scanner) // trainers

	// Read in monsters
	mon = make([]Monster, n)
	for i := range mon {
		scanner.Scan()
		mon[i].id = i
		mon[i].gene = scanner.Text()
		// Calc a hash for each prefix
		mon[i].hash = make([]uint, len(mon[i].gene)+1)
		for j, r := range mon[i].gene {
			mon[i].hash[j+1] = (mon[i].hash[j]*29 + uint(r)) % MOD
		}
	}
	// Sort alphabetically
	sort.Slice(mon, func(i, j int) bool {
		return mon[i].gene < mon[j].gene
	})
	mapping := make([]int, n)
	for i, m := range mon {
		mapping[m.id] = i
	}

	// Process trainers
	for i := 0; i < q; i++ {
		k, l := nextInt(scanner), nextInt(scanner)
		owned := make([]int, k)
		for i := range owned {
			owned[i] = mapping[nextInt(scanner)-1]
		}
		sort.Ints(owned)
		ct := count(owned, l)
		fmt.Printf("%d\n", ct)
	}
}

func count(owned []int, l int) int {
	// Check each window of size l
	tot := 0
	for i := 0; i+l <= len(owned); i++ {
		//fmt.Printf("    i=%v\n", i)
		start, end := i, i+l-1
		bound := 0
		// Check the front boundary for any existing LCP
		if start > 0 {
			if b := match(owned[start-1], owned[start]); b > bound {
				bound = b
			}
		}
		// Check the bottom boundary for any future LCP
		if end+1 < len(owned) {
			if b := match(owned[end], owned[end+1]); b > bound {
				bound = b
			}
		}
		// Check for the LCP specific to this window
		lcp := match(owned[start], owned[end]) - bound
		//fmt.Printf("    lcp=%v\n", lcp)
		if lcp > 0 {
			tot += lcp
		}
	}
	return tot
}

// Match returns the size of the longest common prefix between two monsters,
// using a binary search over the two gene hashes
func match(s, e int) int {
	x, y := mon[s].hash, mon[e].hash
	lo, hi := 0, len(x)-1
	if len(x) > len(y) {
		hi = len(y)-1
	}
	for lo != hi {
		mid := (lo + hi + 1) / 2
		if x[mid] == y[mid] {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
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
