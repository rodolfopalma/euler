// PROBLEM 12
// Highly divisible triangular number

package problems

import (
	"fmt"
)

func Problem_12() {
	target := 500
	sum := 0
	current := 0

	for i := 1; current < target; i++ {
		sum += i
		current = numberOfDivisors(sum)
	}

	fmt.Println(sum, current)

}
