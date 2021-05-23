package easy

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	var res []string

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			res = append(res, "FizzBuzz")
			continue
		}
		if i%5 == 0 {
			res = append(res, "Buzz")
			continue
		}
		if i%3 == 0 {
			res = append(res, "Fizz")
			continue
		}
		res = append(res, strconv.Itoa(i))
	}

	return res
}
