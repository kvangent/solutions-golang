package main

// [Metadata]
// Title: Dominant Strings
// URL: https://open.kattis.com/problems/dominant
// Categories: strings, maps
// Difficulty: 9.0

import (
	"bufio"
	"os"
	"sort"
)

type multiset map[rune]int

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	writer := bufio.NewWriter(os.Stdout)
	words := make(map[string]multiset, 0) // Track dominate words
	for scanner.Scan() {
		w := scanner.Text()
		wSet := newMultiset(w)
		wDominate := true
		for o, oSet := range words {
			if len(w) == len(o) {
				continue // words of the same length can't be supersets of each other
			}
			if wSet.dominates(oSet) {
				delete(words, o) // o is no longer dominate
			} else if oSet.dominates(wSet) {
				wDominate = false
				break
			}
		}
		if wDominate {
			words[w] = wSet
		}
	}

	sorted := make([]string, 0, len(words))
	for w := range words {
		sorted = append(sorted, w)
	}
	sort.Strings(sorted)
	for _, w := range sorted {
		writer.WriteString(w)
		writer.WriteRune('\n')
	}
	writer.Flush()
}

// Returns a multiset created from a word
func newMultiset(word string) multiset {
	set := make(multiset, len(word))
	for _, r := range word {
		set[r] += 1
	}
	return set
}

// Returns true if x is a proper superset of y
func (x multiset) dominates(y multiset) bool {
	// If x has less runes, x is not a superset
	if len(x) < len(y) {
		return false
	}
	// if y has a higher rune count, x is not a superset
	for r, ct := range y {
		if ct > x[r] {
			return false
		}
	}
	return true
}
