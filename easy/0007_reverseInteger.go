package easy

import (
	"math"
)

func reverse(x int) int {
	//original := x
	var answer int
	digits := countDigits(x)
	for x != 0 {
		answer += (x % 10) * int(math.Pow(10, float64(digits-1)))
		x /= 10
		digits--
	}

	if answer >= int(math.Pow(2, 31)) || answer < int(math.Pow(-2, 31)) {
		answer = 0
	}

	//fmt.Printf("Input: %v\t :: reversed digits: \t%v\n", original, answer)
	return answer
}

func countDigits(x int) int {
	count := 0
	for x != 0 {
		x /= 10
		count += 1
	}
	return count
}
