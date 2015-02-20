package problems

import (
	"fmt"
)

func Problem_9() {
	// Absolutley brute force
	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			for k := 1; k < 1000; k++ {

				if i*i+j*j == k*k {
					if i+j+k == 1000 {
						fmt.Println("Found a triple:[", i, j, k, "]")
						fmt.Println(i * j * k)
					}
				}

			}
		}
	}
}
