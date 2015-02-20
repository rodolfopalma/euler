package problems

import (
	"fmt"
)

func Problem_5() {
	i := 20

	for {
		valid := true

		for j := 2; j <= 20; j++ {
			if i%j != 0 {
				valid = false
				break
			}
		}

		if valid {
			fmt.Println(i)
			break
		}

		i++
	}
}
