package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

// [Metadata]
// Title: Catalan Square
// URL: https://open.kattis.com/problems/catalansquare
// Categories: bigint
// Difficulty: 3.6

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	c := new(big.Int).SetInt64(1)
	for i := int64(1); i <= int64(n+1); i++ {
		c.Mul(c, new(big.Int).SetInt64(4*i*i - 2*i))
		c.Div(c, new(big.Int).SetInt64(i*i+i))
	}
	fmt.Println(c.String())
}
