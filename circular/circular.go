package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// [Metadata]
// Title: Circular DNA
// URL: https://open.kattis.com/problems/circular
// Categories: ad hoc
// Difficulty: 2.5

type gene struct {
	isStart bool
	i       int
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Parse the input data.
	n, err := strconv.Atoi(next(scanner))
	check(err)
	genes := make([]gene, n)
	for i := range genes {
		t := []rune(next(scanner))
		genes[i].isStart = t[0] == 's'
		genes[i].i, err = strconv.Atoi(string(t[1:]))
		check(err)
	}

	// Iterate through to find where each gene can be cut.
	ct, min := make(map[int]int, n), make(map[int]int, n)
	totS, totE := make(map[int]int, n), make(map[int]int, n)
	for _, g := range genes {
		// Track the starts and stops to represent nesting.
		if g.isStart {
			ct[g.i] += 1
			totS[g.i] += 1
		} else {
			ct[g.i] -= 1
			totE[g.i] += 1
		}
		// Track the minimum of each gene to valid cuts.
		if ct[g.i] < min[g.i] {
			min[g.i] = ct[g.i]
		}
	}
	// Find the number of genes that can currently be cut.
	valid := 0
	for i := range ct {
		if totS[i] == totE[i] && ct[i] == min[i] {
			valid += 1
		}
	}
	// Iterate through the genes and track max valid cuts.
	maxI, max := 0, valid
	for i, g := range genes {
		if valid > max {
			maxI = i
			max = valid
		}
		if totS[g.i] == totE[g.i] && ct[g.i] == min[g.i] {
			valid -= 1
		}
		if g.isStart {
			ct[g.i] += 1
		} else {
			ct[g.i] -= 1
		}
		if totS[g.i] == totE[g.i] && ct[g.i] == min[g.i] {
			valid += 1
		}


	}
	fmt.Printf("%d %d", maxI+1, max)
}

func check(err error){
	if err != nil {
		panic(err)
	}
}

func next(scanner *bufio.Scanner) string{
	scanner.Scan()
	check(scanner.Err())
	return scanner.Text()
}