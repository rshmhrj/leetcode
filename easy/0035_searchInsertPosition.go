package easy

func searchInsert(nums []int, target int) int {
	p := 0
	if nums[len(nums)-1] < target {
		return len(nums)
	}

	for i, v := range nums {
		if v == target {
			return i
		}
		if v < target {
			p = i + 1
			continue
		}
		if v > target {
			break
		}
	}

	return p
}
