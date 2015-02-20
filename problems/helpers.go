package problems

import (
	"math"
)

func isPrime(n int) bool {

	if n <= 1 {
		return false
	}

	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func reverseString(s string) string {
	if len(s) < 2 {
		return s
	}

	return s[len(s)-1:len(s)] + reverseString(s[0:len(s)-1])
}

func numberOfDivisors(i int) int {
	// Every number has at least 2 divisors; i.e. 1 and itself
	n := 1

	for k := 2; float64(k) < math.Sqrt(float64(i)); k++ {
		if i%k == 0 {
			n++
		}
	}

	return n * 2
}
