package easy

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	reverse := ""

	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}

	//fmt.Printf("Reversed string: %v\n", reverse)

	if s == reverse {
		return true
	}

	return false
}
