package easy

import (
	"math"
)

// titleToNumber converts an excel column name to an integer by calculating as a base 26 number
func titleToNumber(columnTitle string) int {
	c := []rune(columnTitle)
	var res int
	l := len(c) - 1

	// iterate right to left
	for i := l; i >= 0; i-- {
		// convert to base26: 'A' becomes 1 and 'Z' becomes 26
		c[i] -= '@'

		// take power of current position in base26 * current value
		res += int(math.Pow(float64(26), float64(l-i))) * int(c[i])
	}
	return res
}
