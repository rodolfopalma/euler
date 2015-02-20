package problems

import (
	"fmt"
)

func Problem_10() {
	sum := 0

	for i := 0; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
