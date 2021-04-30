package easy

func twoSum(nums []int, target int) []int {
	var result []int
	for indexA := range nums {
		if len(result) == 0 {
			for indexB := range nums {
				if len(result) == 0 {
					if indexA != indexB {
						if nums[indexA]+nums[indexB] == target {
							return []int{indexA, indexB}
						}
					}
				}
			}
		}
	}
	return result
}
