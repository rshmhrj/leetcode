package med

func subsets(nums []int) [][]int {
	var res [][]int

	// base cases
	if len(nums) == 0 {
		res = append(res, []int{})
		return res
	}

	if len(nums) == 1 {
		res = append(res, []int{})
		res = append(res, nums)
		return res
	}

	// split first num from rest of nums and get all subsets
	curr := nums[0]
	remainder := subsets(nums[1:])
	length := len(remainder)

	// create clone of remaining nums, and insert the current number into each subset
	var clone = make([][]int, length)
	for i := range remainder {
		clone[i] = make([]int, len(remainder[i]))
		clone[i] = append(clone[i], curr)
		copy(clone[i], remainder[i])
	}

	// append clone and subsets and return
	res = append(res, clone...)
	res = append(res, remainder...)
	return res
}
