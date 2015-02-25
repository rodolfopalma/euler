package main

import (
	"fmt"
	"github.com/rpalmaotero/euler/problems"
	"time"
)

func main() {
	n := 10
	var sum time.Duration

	for i := 0; i < n; i++ {
		start := time.Now()

		/*
			problems.Problem_1()
			problems.Problem_2()
			problems.Problem_3()
			problems.Problem_4()
			problems.Problem_5()
			problems.Problem_6()
			problems.Problem_7()
			problems.Problem_8()
			problems.Problem_9()
			problems.Problem_10()
			problems.Problem_11()
			problems.Problem_12()w
		*/

		problems.Problem_13()

		sum += time.Since(start)
	}

	fmt.Println(n, " run(s) took an average of", sum/time.Duration(n))

}
