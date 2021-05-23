package easy

func majorityElement(nums []int) int {
	elems := make(map[int]int)

	for _, v := range nums {
		elems[v]++
	}

	var majorityElementValue int
	var majorityElementCount int

	for i, v := range elems {
		if v > len(nums)/2 && v > majorityElementCount {
			majorityElementValue = i
			majorityElementCount = v
		}
	}

	return majorityElementValue
}
