package easy

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}

	nums := make([]int, n)
	nums[0], nums[1] = 1, 1
	for i := 2; i <= n/2; i++ {
		if nums[i] == 1 {
			continue
		}
		nums[0] += removeMults(i, nums)
	}
	return n - nums[0] - 1
}

func removeMults(x int, nums []int) int {
	var removed int
	for i := x + x; i < len(nums); i += x {
		if nums[i] == 0 {
			nums[i] = 1
			removed += 1
		}
	}
	return removed
}
