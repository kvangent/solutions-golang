package main

// [Metadata]
// Title: Otpor
// URL: https://open.kattis.com/problems/otpor
// Categories: ad hoc, stack
// Difficulty: 2.8

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func main() {
	// use bufio.Scanner for improved scanning performance
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Parse the resistor count, values, and circuit
	rCt, _ := strconv.Atoi(next(scanner))
	resistors := make([]float64, rCt)
	for i := range resistors {
		tmp, _ := strconv.ParseFloat(next(scanner), 64)
		resistors[i] = float64(tmp)
	}
	scanner.Scan()
	circuit := scanner.Text()

	// Create a stack for processing
	stack := make([]interface{}, 0, len(circuit))
	for i, r := range circuit {
		switch r {
		case '(', '-', '|':
			stack = append(stack, r)
		case 'R': // Convert numeric rune to resistor value
			rIdx := int(circuit[i+1] - '1')
			stack = append(stack, resistors[rIdx])
		case ')':
			sOp := stack[len(stack)-2].(rune) // series operator
			var total float64
			for {
				// Pop value and operand from the stack
				val, op := stack[len(stack)-1].(float64), stack[len(stack)-2].(rune)
				stack = stack[:len(stack)-2]
				// Track series or parallel value
				if sOp == '-' {
					total += val
				} else { // sOp == '|'
					total += 1.0 / val
				}
				if op == '(' { // quit when we hit the start
					break
				}
			}
			// Add value to stack
			if sOp == '-' {
				stack = append(stack, total)
			} else { // sOp == '|'
				stack = append(stack, 1.0/total)
			}
		}
	}
	fmt.Printf("%.5f", stack[0].(float64))
}

func next(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
