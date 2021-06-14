package med

func permute(nums []int) [][]int {
	var res [][]int

	// base case
	if len(nums) == 1 {
		res = [][]int{[]int{nums[0]}}
		return res
	}

	// basic insertion
	if len(nums) == 2 {
		res = [][]int{[]int{nums[0], nums[1]}, []int{nums[1], nums[0]}}
		return res
	}
	r := permute(nums[1:])

	for _, v := range r {
		for j := 0; j <= len(v); j++ {
			tmp := insert(nums[0], v, j)
			res = append(res, tmp)
		}
	}

	//for i := 0; i < len(nums)-2; i += 2 {
	//  res = append(res, insert(i+3, tmp[i])...)
	//  res = append(res, insert(i+3, tmp[i+1])...)
	//}

	return res
}

func insert(n int, nums []int, pos int) []int {
	var res []int
	res = append(res, nums[:pos]...)
	res = append(res, n)
	res = append(res, nums[pos:]...)
	return res
}

//func insert(n int, nums []int, pos int) []int {
//  var res []int
//  l := len(nums)
//  for i := 0; i <= l; i++ {
//    if i == 0 {
//      tmp := []int{n}
//      tmp = append(tmp, nums...)
//      res = append(res, tmp)
//      continue
//    }
//    if i == l {
//      tmp := nums
//      tmp = append(tmp, n)
//      res = append(res, tmp)
//      continue
//    }
//    b := make([]int, i-1)
//    b = append(b, nums[:i]...)
//    a := make([]int, len(nums)-i-1)
//    a = append(a, nums[i:]...)
//    tmp := b
//    tmp = append(tmp, n)
//    tmp = append(tmp, a...)
//    res = append(res, tmp)
//  }
//  return res
//}
