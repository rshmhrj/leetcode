package easy

import (
	"strings"
)

func buddyStrings(a string, b string) bool {
	// check lengths
	if len(a) != len(b) {
		return false
	}
	if len(a) == 1 {
		return false
	}

	// check if same string already
	if a == b {
		alphabet := "abcdefghijklmnopqrstuvwxyz"
		// if any repeating letter
		for i := range alphabet {
			if strings.Count(a, string(alphabet[i])) > 1 {
				return true
			}
		}
	}

	var misplacedLetters []int
	// find misplaced letter(s)
	for i := range b {
		if a[i] != b[i] {
			misplacedLetters = append(misplacedLetters, i)
		}
	}

	// validate only 2 misplacedLetters
	if len(misplacedLetters) != 2 {
		return false
	}

	// validate swap results in same string
	newA := []rune(a)
	//goland:noinspection ALL
	newA[misplacedLetters[0]], newA[misplacedLetters[1]] = newA[misplacedLetters[1]], newA[misplacedLetters[0]]

	if string(newA) == b {
		return true
	}

	// return if swap results in same string
	return false
}
