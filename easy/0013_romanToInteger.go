package easy

func romanToInt(s string) int {

	if len(s) <= 2 {
		switch s {
		case "I":
			return 1
		case "II":
			return 2
		case "III":
			return 3
		case "IV":
			return 4
		case "V":
			return 5
		case "VI":
			return 6
		case "VII":
			return 7
		case "VIII":
			return 8
		case "IX":
			return 9
		case "X":
			return 10
		case "XL":
			return 40
		case "L":
			return 50
		case "XC":
			return 90
		case "C":
			return 100
		case "CD":
			return 400
		case "D":
			return 500
		case "CM":
			return 900
		case "M":
			return 1000
		//case "MCD":
		//  return 1400
		default:
			if len(s) == 0 {
				return 0
			}
			return romanToInt(string(s[0])) +
				romanToInt(string(s[1:]))
		}
	} else {
		for i := range s {
			if i < len(s)-2 {
				if string(s[i:i+2]) == "IV" || string(s[i:i+2]) == "IX" || string(s[i:i+2]) == "XL" || string(s[i:i+2]) == "XC" || string(s[i:i+2]) == "CD" || string(s[i:i+2]) == "CM" {
					return romanToInt(string(s[:i])) +
						romanToInt(string(s[i:i+2])) +
						romanToInt(string(s[i+2:]))
				}
			}
		}
	}

	return romanToInt(string(s[0:len(s)/2])) + romanToInt(string(s[len(s)/2:]))
}
