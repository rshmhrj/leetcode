package easy

import (
	"math"
)

func removeDuplicates(nums []int) int {
	// quick return for non-duplicate arrays
	if len(nums) <= 1 {
		return len(nums)
	}
	deDupedIndex := 0
	seen := math.MinInt32
	for i, v := range nums {
		// if first element, increment count and move on
		if i == 0 {
			seen = v
			continue
		}

		// if greater than what we saw before, increment index, swap places, and update seen value
		if v > seen {
			deDupedIndex++
			nums[deDupedIndex], nums[i] = nums[i], nums[deDupedIndex]
			seen = v
			continue
		}
	}
	return deDupedIndex + 1
}
