package easy

func mySqrt(x int) int {
	for i := 0; i <= x/2+1; i++ {
		if i*i == x {
			return i
		}
		if i*i > x {
			return i - 1
		}
	}
	return 1
}
