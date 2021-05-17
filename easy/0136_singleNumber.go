package easy

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	stack := initStack(len(nums))

	i := 0
	for _, v := range nums {
		for k := 0; k <= i; k++ {
			if v == stack[k] && i != 0 {
				stack[k] = stack[i-1]
				i--
				break
			}
			if k == i {
				stack[i] = v
				i++
				break
			}
		}

	}
	return stack[0]
}

func initStack(len int) []int {
	stack := make([]int, len/2+2)
	for i := range stack {
		stack[i] = -31000
	}
	return stack
}
