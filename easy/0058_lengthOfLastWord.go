package easy

func lengthOfLastWord(s string) int {
	l := 0
	var space byte = 32

	if len(s) == 1 && s[0] == space {
		return 0
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == space && l > 0 {
			return l
		}
		if s[i] == space && l == 0 {
			continue
		}
		l++
	}

	return l
}
