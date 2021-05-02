package easy

func strStr(haystack string, needle string) int {
	output := -1

	// handle empty strings
	if needle == "" {
		return 0
	}

	// handle needle bigger than haystack
	if haystack == "" || len(needle) > len(haystack) {
		return -1
	}

	for i := range haystack {
		if i+len(needle) < len(haystack) {
			if needle == haystack[i:i+len(needle)] {
				output = i
				break
			}
		} else if needle == haystack[i:] {
			output = i
			break
		}
	}

	return output
}
