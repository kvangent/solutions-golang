package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// [Metadata]
// Title: Baloni
// URL: https://open.kattis.com/problems/baloni
// Categories: ad hoc
// Difficulty: 4.1

func main(){
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	// Scan input
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// Parse balloons height
	balloons := make([]int, n)
	for i := range balloons {
		scanner.Scan()
		balloons[i], _ = strconv.Atoi(scanner.Text())
	}

	count := 0
	darts := make(map[int]int)
	for _, h := range balloons {
		if darts[h] > 0 { // Dart drops
			darts[h]--
		} else { // Throw new dart
			count++
		}
		darts[h-1]++
	}

	fmt.Print(strconv.Itoa(count))
}
