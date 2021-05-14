package med

import (
	"math"
)

func myAtoi(s string) int {
	// convert to rune array
	r := []rune(s)
	// assume positive
	sign := 1

	var res float64
	var start, end int
	var numStarted, signSet bool

	for i, v := range r {
		if isNum(r[i]) {
			// if no nums before this, set flag and update start
			if !numStarted {
				numStarted = true
				start = i
			}
			// update end
			end = i
			continue
		}
		if !isNum(r[i]) {
			if numStarted {
				break
			} else {
				// ignore whitespace
				if v == ' ' {
					continue
				}
				// set sign
				if v == '-' && sign == 1 {
					if signSet {
						return 0
					}
					sign *= -1
					signSet = true
					continue
				}
				// update flag for plus
				if v == '+' {
					if signSet {
						return 0
					}
					if i+1 < len(r)-1 {
						if !isNum(r[i+1]) {
							return 0
						}
					}
					signSet = true
					continue
				}
				// otherwise quick return
				return 0
			}
		}
	}
	// temp array to hold actual numbers
	var t []rune
	if end+1 > len(r) {
		t = r[start:]
	} else {
		t = r[start : end+1]
	}

	// edge checking
	if len(t) == 0 || !numStarted {
		return 0
	}

	// populate result with values from temp array
	j := 0
	for i := len(t) - 1; i >= 0; i-- {
		res += float64(t[i]-'0') * math.Pow(10, float64(j))
		j++
	}

	// update sign of result
	res *= float64(sign)

	// clamp to bounds
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	if res < math.MinInt32 {
		return math.MinInt32
	}

	return int(res)
}

func isNum(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}
