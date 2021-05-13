package easy

func addBinary(a string, b string) string {
	if a == "0" {
		return b
	}
	if b == "0" {
		return a
	}

	c := ""

	bin := map[byte]int{48: 0, 49: 1}

	carry := false

	i := len(a) - 1
	j := len(b) - 1

	for i >= 0 && j >= 0 {
		if bin[uint8(a[i])]+bin[uint8(b[j])] == 0 {
			if carry {
				c = "1" + c
			} else {
				c = "0" + c
			}
			carry = false
		}
		if bin[uint8(a[i])]+bin[uint8(b[j])] == 1 {
			if carry {
				c = "0" + c
			} else {
				c = "1" + c
			}
		}
		if bin[uint8(a[i])]+bin[uint8(b[j])] == 2 {
			if carry {
				c = "1" + c
			} else {
				c = "0" + c
			}
			carry = true
		}
		i--
		j--
	}

	for ; i >= 0; i-- {
		if bin[uint8(a[i])] == 0 {
			if carry {
				c = "1" + c
				carry = false
			} else {
				c = "0" + c
			}
		}
		if bin[uint8(a[i])] == 1 {
			if carry {
				c = "0" + c
			} else {
				c = "1" + c
			}
		}
	}
	for ; j >= 0; j-- {
		if bin[uint8(b[j])] == 0 {
			if carry {
				c = "1" + c
				carry = false
			} else {
				c = "0" + c
			}
		}
		if bin[uint8(b[j])] == 1 {
			if carry {
				c = "0" + c
			} else {
				c = "1" + c
			}
		}
	}

	if carry {
		c = "1" + c
	}

	return c
}
