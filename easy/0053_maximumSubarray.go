package easy

func maxSubArray(nums []int) int {
	// set max to first value
	max := nums[0]

	// short-circuit for small arrays
	if len(nums) == 1 {
		return max
	}

	// set first value as sum
	sum := max

	// if negative, start at 0
	if sum < 0 {
		sum = 0
	}

	// first value already taken into account, so iterate from second value
	// add value to sum and if sum is greater than max, use it.
	// if sum is negative, reset to 0
	for _, v := range nums[1:] {
		sum += v
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return max
}
