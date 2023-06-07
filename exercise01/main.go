package main

import (
	"bufio"
	"fmt"
	"os"
)

func calculateAmountToPay(s []int) int {
	amount := 0
	// Санёк, вставь свой код тут
	return amount
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var n int
		fmt.Fscan(in, &n)
		s := make([]int, n, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &s[j])
		}
		fmt.Fprintln(out, calculateAmountToPay(s))
	}
}
