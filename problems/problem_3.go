package problems

import (
	"fmt"
	"math"
)

func Problem_3() {
	n := 600851475143
	m := int(math.Sqrt(float64(n)))

	for i := m; i > 2; i-- {
		if !(n%2 == 0) && n%i == 0 && isPrime(i) {
			fmt.Println(i)
		}
	}
}
