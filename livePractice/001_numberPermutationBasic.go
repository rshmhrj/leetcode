package livePractice

import (
	"math"
)

/**
Create a program makeNumberBasic(z) which, when given an input of a number (z) 5 digits, returns the number of all possible permutations of digits (1 through 9 inclusive) that when added will equal z.

For example, if z is 3, your program will find that four permutations of digits add up to that value (3, 2+1, 1+1+1, and 1+2), and thus return 4.
You may limit makeNumberBasic(z) to the use of five digits
Repeat use of a digit is acceptable: e.g. 1+1+1 would be a valid addition of digits equalling 3.
Use of a single digit is acceptable as a permutation: e.g. 3 is itself a valid permutation of digits that add up to 3.
makeNumberBasic(z) is looking for permutations 15, not combinations 5: 1+5 and 5+1 would count as two unique possible ways to add to 6, not one. Read more about permutations vs. combinations here 15.
If no permutations of the digits 1 through 9 add up to the number z, your function should return 0.

makeNumberBasic(2) = 2  // 1,1; 2
makeNumberBasic(3) = 4  // 1,1,1; 1,2; 2,1; 3
makeNumberBasic(5) = 16 // 1,1,1,1,1; 1,2,1,1; 1,1,2,1; 1,1,1,2; 2,1,1,1; 5; 1,1,3; 1,3,1; 3,1,1; 1,2,2; 2,1,2; 2,2,1; 1,4; 4,1; 2,3; 3,2;
makeNumberBasic(7) =    //[[7],[1,6],[2,5],[3,4],[1,1,1,1,1,1,1],[2,2,2,1]]

// 1 unique #: count += 1
n = n! / (n-r)! = 1! / (1-1)!
// 2 unique #: count += 2
n = n! / (n-r)! = 2! / (2-1)!
// 3 #s, 2 unique: count += 3
n = n! / (n-r)! = 3! / (3-1)! = 3!/2! = 3
// 3 #s, 3 unique 1,2,3; 1,3,2; 2,1,3; 2,3,1; 3,1,2; 3,2,1;
n = n! / (n-r)! = 3! / (3-2)! = 3!/1! = 6

// store 1-9 values in array
// [1,1,0,0,0]
// i = 0 to 5 (represents first index in array)
// j = 0 to 5 (represents last index in array)
// start with index 0 and increment 1-9 to see if sum is x
// if sum is x, then apply formula
// start with indices 0 and 1 and increment 1-9 to see if sum is x
// if sum is x, then apply formula
// ... repeat for indicies 0-5
**/

func makeNumberBasic(x int) int {
	output := 0
	for c := 0; c <= 99999; c++ {
		if sum(c) == x && !(stripZeros(c) < c) {
			output += 1
			continue
		}
	}
	return output
}

func sum(x int) int {
	res := 0
	for x != 0 {
		res += x % 10
		x /= 10
	}
	return res
}

func stripZeros(x int) int {
	i := 0
	res := 0
	for x != 0 {
		if r := x % 10; r > 0 {
			res += r * int(math.Pow(10, float64(i)))
			i++
		}
		x /= 10
	}
	return res
}
