package problems

import (
	"fmt"
)

// Longest Collatz sequence
func Problem_14() {
	max := 0

	for i := 0; i < 1000000; i++ {
		if i%2 != 0 {
			current := len(recursiveCollatz([]int{i}))

			if current > max {
				max = current
				fmt.Println("Current max is", i, "with", current, "terms.")
			}
		}
	}
}

func recursiveCollatz(seq []int) []int {
	last := seq[len(seq)-1]

	if last <= 1 {
		return seq
	} else {
		var next int

		if last%2 == 0 {
			// Even
			next = last / 2
		} else {
			// Odd
			next = 3*last + 1
		}

		return recursiveCollatz(append(seq, next))
	}

}
