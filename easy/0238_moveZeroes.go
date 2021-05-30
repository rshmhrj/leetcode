package easy

func moveZeroes(nums []int) {
	l := len(nums)
	// short circuit for small arrays
	if l <= 1 {
		return
	}

	// for each item
	// if 0, swap with next non-0
	for i, v := range nums[:l-1] {
		if v == 0 && i != l-1 {
			for j := i + 1; j < l; j++ {
				if nums[j] == 0 {
					continue
				}
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
}
