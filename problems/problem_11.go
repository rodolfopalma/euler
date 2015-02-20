package problems

import (
	"fmt"
	"sort"
)

func Problem_11() {
	grid := [][]int{
		{8, 02, 22, 97, 38, 15, 0, 40, 0, 75, 4, 5, 07, 78, 52, 12, 50, 77, 91, 8},
		{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 04, 56, 62, 00},
		{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 03, 49, 13, 36, 65},
		{52, 70, 95, 23, 04, 60, 11, 42, 69, 24, 68, 56, 01, 32, 56, 71, 37, 02, 36, 91},
		{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80},
		{24, 47, 32, 60, 99, 03, 45, 02, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50},
		{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70},
		{67, 26, 20, 68, 02, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21},
		{24, 55, 58, 05, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72},
		{21, 36, 23, 9, 75, 00, 76, 44, 20, 45, 35, 14, 00, 61, 33, 97, 34, 31, 33, 95},
		{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 03, 80, 04, 62, 16, 14, 9, 53, 56, 92},
		{16, 39, 05, 42, 96, 35, 31, 47, 55, 58, 88, 24, 00, 17, 54, 24, 36, 29, 85, 57},
		{86, 56, 00, 48, 35, 71, 89, 07, 05, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58},
		{19, 80, 81, 68, 05, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 04, 89, 55, 40},
		{04, 52, 8, 83, 97, 35, 99, 16, 07, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66},
		{88, 36, 68, 87, 57, 62, 20, 72, 03, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69},
		{04, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36},
		{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 04, 36, 16},
		{20, 73, 35, 29, 78, 31, 90, 01, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 05, 54},
		{01, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 01, 89, 19, 67, 48}}

	dx, dy := len(grid[0]), len(grid)
	n := 4
	values := make([]int, 0)

	for i, j := n-1, n-1; i <= dx-n && j <= dy-n; {

		// Top
		top := 1
		left := 1
		right := 1
		diagtopleft := 1
		diagtopright := 1

		for k, y, z := i, 0, 0; k > i-n; k-- {
			top = top * grid[k][j]

			// Left and Diag Top Left
			for l := j; l > j-n; l-- {
				if y < n {
					left = left * grid[i][l]
				}

				if y%(n+1) == 0 {
					diagtopleft = diagtopleft * grid[k][l]
				}

				y++
			}

			// Right and Diag Top Right
			for l := j; l < j+n; l++ {
				if z < n {
					right = right * grid[i][l]
				}

				if z%(n+1) == 0 {
					diagtopright = diagtopright * grid[k][l]
				}

				z++
			}

		}

		bottom := 1
		diagbottomleft := 1
		diagbottomright := 1

		// Bottom

		for k, y, z := i, 0, 0; k < i+n; k++ {
			bottom = bottom * grid[k][j]

			// Diag Bottom Left
			for l := j; l > j-n; l-- {
				if y%(n+1) == 0 {
					diagbottomleft = diagbottomleft * grid[k][l]
				}

				y++
			}

			// Diag Bottom Right
			for l := j; l < j+n; l++ {
				if z%(n+1) == 0 {
					diagbottomright = diagbottomright * grid[k][l]
				}

				z++
			}
		}

		// Reseting when doing edges
		if i == dx-n {
			i = n - 1
			j++
		} else {
			i++
		}

		values = append(values, top, right, bottom, left, diagbottomleft, diagbottomright, diagtopleft, diagtopright)
	}

	sort.Ints(values)
	fmt.Println(values[len(values)-1:])
	fmt.Println(numberOfDivisors(21))
}
