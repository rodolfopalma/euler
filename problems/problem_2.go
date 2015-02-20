package problems

import (
	"fmt"
)

func Problem_2() {
	n := 4000000

	sum := 0

	for i, j := 1, 2; j < n; i, j = j, i+j {
		if j%2 == 0 {
			sum += j
		}
	}

	fmt.Println(sum)
}
