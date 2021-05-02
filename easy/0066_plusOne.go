package easy

func plusOne(digits []int) []int {
	// array length
	l := len(digits) - 1

	// simple case with no carry
	if digits[l] < 9 {
		digits[l]++
		return digits
	}

	// single 9
	if l == 0 && digits[l] == 9 {
		return []int{1, 0}
	}

	// multiple digits
	carry := true
	for i := l; i >= 0; i-- {
		if digits[i] == 9 && i > 0 && carry == true {
			digits[i] = 0
			digits[i-1]++
			carry = false
			continue
		}
		if digits[i] == 10 && i > 0 {
			digits[i] = 0
			digits[i-1]++
			continue
		}
		if i == 0 && digits[i] < 10 {
			return digits
		}
		if i == 0 && digits[i] == 10 {
			tmp := make([]int, l+2)
			tmp[0], tmp[1] = 1, 0
			for j := 2; j <= l+1; j++ {
				tmp[j] = digits[j-1]
			}
			return tmp
		}
	}
	return digits
}
