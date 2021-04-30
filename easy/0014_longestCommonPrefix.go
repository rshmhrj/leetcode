package easy

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	strs = strs[1:]
	for _, v := range strs {
		for ; strings.Index(v, prefix) != 0; prefix = prefix[:len(prefix)-1] {
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}
