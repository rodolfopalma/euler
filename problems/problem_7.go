package problems

import (
	"fmt"
)

func Problem_7() {
	target := 10001

	for i, n := 1, 0; n < target; i++ {
		if isPrime(i) {
			n++
			fmt.Println(i, n)
		}
	}
}
