package easy

func findKthPositive(arr []int, k int) int {
	var input IntSlice
	for _, value := range arr {
		input = append(input, value)
	}
	var seenNums IntSlice
	var missingNums IntSlice
	var i int = 1

	var j int

	for {
		if j < len(arr) {
			for _, value := range arr {
				if i == value || missingNums.Has(i) || seenNums.Has(i) || input.Has(i) {
					break
				}
				if i != value && !missingNums.Has(i) && !seenNums.Has(i) && !input.Has(i) {
					missingNums = append(missingNums, i)
				}
			}
		} else {
			missingNums = append(missingNums, i)
		}

		if len(missingNums) == k+3 {
			//fmt.Printf("arr: %v\t k: %v\t :: missing integers are %v\t :: kth integer is \t%v \n", arr, k, missingNums, missingNums[k-1])
			return missingNums[k-1]
		}
		if j < len(arr) && arr[j] <= i {
			seenNums = append(seenNums, arr[j])
			j++
		}
		i++

	}
}

type IntSlice []int

func (list IntSlice) Has(a int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
