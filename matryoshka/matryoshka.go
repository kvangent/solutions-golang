package main

import (
	"bufio"
	"os"
	"strconv"
	"math"
	"fmt"
	"strings"
)

// [Metadata]
// Title: Matryoshka
// URL: https://open.kattis.com/problems/matryoshka
// Categories: dp
// Difficulty: 4.2

var dolls []int
var memSolve []int
var memGroup [][]int
var minRange [][]int
var cumFreqs [][]int

const INF = math.MaxInt32 / 3
const UNSOLVED = -1

func main(){
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	// Scan input
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// Parse dolls
	dolls = make([]int, n)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	for i := range dolls {
		dolls[i], _ = strconv.Atoi(input[i])
	}

	// Prepare memoization tables (-1 is unsolved)
	memSolve = make([]int, n+1)
	memGroup = make([][]int, n+1)
	for i := 0; i < len(memSolve); i++ {
		memSolve[i], memGroup[i] = UNSOLVED, make([]int, n+1)
		for j := range memGroup {
			memGroup[i][j] = UNSOLVED
		}
	}

	// Calculate minimum for each range [start, end)
	minRange = make([][]int, n+1)
	for i := range minRange {
		minRange[i] = make([]int, n+1)
		min := INF
		for j := i; j < len(dolls); j++ {
			if dolls[j] < min {
				min = dolls[j]
			}
			minRange[i][j+1] = min
		}
	}

	// Calculate number of dolls from [start, i) below a certain size
	cumFreqs = make([][]int, n+1)
	for i := range cumFreqs {
		cumFreqs[i] = make([]int, 501)
	}
	for i := 0; i < len(cumFreqs); i++ {
		for j := i; j < len(dolls); j++{
			cumFreqs[i][dolls[j]] += 1
		}
	}
	for i := 0; i < len(cumFreqs); i++ {
		for j := 1; j < len(cumFreqs[i]); j++ {
			cumFreqs[i][j] += cumFreqs[i][j-1]
		}
	}

	// Solve the problem!
	solution := solve(0)
	if solution < INF {
		fmt.Println(solution)
	} else {
		fmt.Println("impossible")
	}
}

// Returns the solution for the dolls for [i to n).
func solve(i int) int {
	if i == len(dolls) { return 0 }
	if memSolve[i] != UNSOLVED { return memSolve[i] }
	min := INF
	for j := i+1; j <= len(dolls); j++ {
		if isCompleteSet(i, j) {
			result := group(i, j) + solve(j)
			if result < min {
				min = result
			}
		}
	}
	memSolve[i] = min
	return min
}

// Returns true if [start, end) forms a complete set.
func isCompleteSet(start, end int) bool {
	seen := make(map[int]bool, end-start)
	max := 0
	for i := start; i < end; i++ {
		size := dolls[i]
		if _, inSet := seen[size]; inSet {
			return false // duplicate doll size
		}
		seen[size] = true
		if size > max {
			max = size
		}
	}
	// (1, 2, 3, ..., MAX) if valid set
	return len(seen) == max
}


// Returns number of moves to turn [start, end) into a set.
func group(start, end int) int {
	if end < start + 2 {
		return 0 // less than one doll
	}
	if memGroup[start][end] != UNSOLVED {
		return memGroup[start][end]
	}
	min := INF
	for mid := start+1; mid < end; mid++ {
		result := group(start, mid) + group(mid, end) + combine(start, mid, end)
		if result < min {
			min = result
		}
	}
	memGroup[start][end] = min
	return min
}

// Returns cost to combine group [start, mid) with [mid, end).
func combine(start, mid, end int) int { // O(n)
	leftMin := minRange[start][mid]
	rightMin := minRange[mid][end]
	cost := end - start // total number of dolls
	if leftMin < rightMin { //if left has smaller group
		// don't open ones in left that are smaller than rightMin
		cost -= cumFreqs[start][rightMin-1] - cumFreqs[mid][rightMin-1]
	} else {
		// don't open ones in right that are smaller than leftMin
		cost -= cumFreqs[mid][leftMin-1] - cumFreqs[end][leftMin-1]
	}
	return cost
}
