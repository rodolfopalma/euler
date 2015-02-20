package problems

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func Problem_4() {

	palis := make([]int, 1)

	for i := 100; i < 999; i++ {
		for j := 999; j > 99; j-- {
			prod := j * i
			str := strconv.Itoa(prod)
			l := len(str)
			pali := false

			if l%2 == 0 {
				if str[:(l/2)] == reverseString(str[(l/2):]) {
					pali = true
				}
			} else {
				if str[:int(math.Floor(float64(l/2)))] == reverseString(str[1+int(math.Ceil(float64(l/2))):]) {
					pali = true
				}
			}
			if pali {
				val, _ := strconv.Atoi(str)
				palis = append(palis, val)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(palis)))

	fmt.Println(palis)

}
