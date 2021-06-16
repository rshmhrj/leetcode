package med

func longestPalindrome(s string) string {
	l := len(s)
	// quick return if length is 1
	if l == 1 {
		return s
	}

	// global max
	var longest string

	// two finger traversal across string, checking for palindromes of increasing length
	// track local max starting from each index
	for i := 0; i < l; i++ {
		// only consider checking for substrings longer than current longest
		for j := i + len(longest) + 1; j <= l; j++ {
			var w string
			if j == l {
				w = s[i:]
			} else {
				w = s[i:j]
			}
			if isPalindrome(w) && len(w) > len(longest) {
				longest = w
			}
		}
	}

	// return max
	return longest
}

func isPalindrome(s string) bool {
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
