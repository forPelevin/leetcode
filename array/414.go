package array

import "math"

//https://leetcode.com/problems/third-maximum-number/
//414. Third Maximum Number
//Given a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).
//
//Example 1:
//Input: [3, 2, 1]
//
//Output: 1
//
//Explanation: The third maximum is 1.
//Example 2:
//Input: [1, 2]
//
//Output: 2
//
//Explanation: The third maximum does not exist, so the maximum (2) is returned instead.
//Example 3:
//Input: [2, 2, 3, 1]
//
//Output: 1
//
//Explanation: Note that the third maximum here means the third maximum distinct number.
//Both numbers with value 2 are both considered as second maximum.
func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	firstMax, secondMax, thirdMax := int(math.Inf(-1)), int(math.Inf(-1)), int(math.Inf(-1))
	for _,n := range nums {
		if n >= firstMax {
			if n == firstMax {
				continue
			}

			thirdMax = secondMax
			secondMax = firstMax
			firstMax = n
			continue
		}
		if n >= secondMax {
			if n == secondMax {
				continue
			}

			thirdMax = secondMax
			secondMax = n
			continue
		}
		if n >= thirdMax {
			thirdMax = n
		}
	}

	if thirdMax != int(math.Inf(-1)) {
		return thirdMax
	}

	return firstMax
}
