package med

import (
	"math"
)

func divide(dividend int, divisor int) int {
	var res int
	var resultNegative bool
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		resultNegative = true
	}

	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	for dividend-divisor >= 0 {
		dividend -= divisor
		res += 1
	}

	if resultNegative {
		res = -res
	}

	if res > math.MaxInt32 {
		res = math.MaxInt32
	}

	if res < math.MinInt32 {
		res = math.MinInt32
	}

	return res
}
