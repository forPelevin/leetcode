package other

import (
	"math"
	"math/bits"
)

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/762/
//The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
//
//Given two integers x and y, calculate the Hamming distance.
//
//Note:
//0 ≤ x, y < 231.
//
//Example:
//
//Input: x = 1, y = 4
//
//Output: 2
//
//Explanation:
//1   (0 0 0 1)
//4   (0 1 0 0)
//↑   ↑
//
//The above arrows point to positions where the corresponding bits are different.
func hammingDistance(x int, y int) int {
	if x < 0 || y < 0 || x > math.MaxInt32 || y > math.MaxInt32 {
		return 0
	}
	res := 0
	mask := 1
	for i := 0; i < 32; i++ {
		if (x & mask) != (y & mask) {
			res++
		}
		mask <<= 1
	}
	return res
}

func cheaterHammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
