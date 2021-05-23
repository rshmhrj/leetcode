package easy

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[string]int)
	tMap := make(map[string]int)

	for _, v := range s {
		sMap[fmt.Sprintf("%v", v)]++
	}

	for _, v := range t {
		tMap[fmt.Sprintf("%v", v)]++
	}

	for i, v := range sMap {
		if tMap[i] != v {
			return false
		}
	}

	return true
}
