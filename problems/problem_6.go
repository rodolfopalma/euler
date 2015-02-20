package problems

import (
	"fmt"
)

func Problem_6() {
	sqr_sum := 0
	sum := 0

	for i := 1; i <= 100; i++ {
		sqr_sum += i * i
		sum += i
	}

	fmt.Println(sum*sum - sqr_sum)
}
