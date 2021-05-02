package easy

func removeElement(nums []int, val int) int {
	length := 0
	x := 0
	if len(nums) == 0 {
		return x
	}

	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] == val {
			x = i
			for j := len(nums) - 1; j > i; j-- {
				if nums[j] != val {
					nums[x], nums[j] = nums[j], nums[x]
					length++
					break
				}
				continue
			}
		} else {
			length++
		}
	}
	return length
}
