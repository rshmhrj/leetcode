package med

func lengthOfLongestSubstring(s string) int {
	var longestSubstring string
	// for each letter
	if len(s) == 1 {
		longestSubstring = s
	}
	for i := range s {
		endThisLoop := false

		// start at next letter and continue to end of string
		for e, c := range s[i+1:] {
			if endThisLoop {
				break
			}
			// compare that letter to all letters before
			//currentSubstring := s[i : i+e+1]
			currentSubstring := StringSlice(string(s[i : i+e+1]))
			if isRepeat := currentSubstring.Has(string(c)); isRepeat {
				// if it was a repeat then take substring from beginning of this iteration to letter before this repeated letter
				// check to see if it is longer
				if len(currentSubstring) > len(longestSubstring) {
					longestSubstring = currentSubstring.ToString()
				}
				endThisLoop = true
				break
			}
			// if letter is the last letter, and there were no other repeats, then take the whole substring
			if i+e+1 == len(s)-1 {
				// check to see if it is longer
				if len(currentSubstring)+1 > len(longestSubstring) {
					longestSubstring = currentSubstring.ToString() + string(s[len(s)-1])
				}
				endThisLoop = true
				break
			}
		}
	}
	//fmt.Printf("Input: \t %s   \t :: \t\t substring : \t %s \t :: \t length : \t %d\n",
	//  s, longestSubstring, len(longestSubstring))
	return len(longestSubstring)
}

type StringSlice string

func (s StringSlice) Has(c string) bool {
	for _, l := range s {
		if string(c) == string(l) {
			return true
		}
	}
	return false
}
func (s StringSlice) ToString() string {
	var output string
	for _, l := range s {
		output += string(l)
	}
	return output
}
