package easy

import (
	"math"
)

func thirdMax(nums []int) int {
	// init variables
	lowerBound := math.MinInt32 - 1
	max, secondMax, thirdMax := lowerBound, lowerBound, lowerBound

	// find max
	for _, v := range nums {
		// skip assignment if it is a duplicate
		if v == max || v == secondMax || v == thirdMax {
			continue
		}
		// assign all three if max is found
		if v > max && v != lowerBound {
			thirdMax, secondMax, max = secondMax, max, v
			continue
		}
		if v > secondMax && v != lowerBound {
			thirdMax, secondMax = secondMax, v
			continue
		}
		if v > thirdMax {
			thirdMax = v
		}
	}

	// if no third max, return max
	if thirdMax == lowerBound || thirdMax == secondMax {
		return max
	}
	return thirdMax
}
