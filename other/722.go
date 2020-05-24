package other

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/722/
//Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.
//
//Example 1:
//
//Input: [3,0,1]
//Output: 2
//Example 2:
//
//Input: [9,6,4,2,3,5,7,0,1]
//Output: 8
//Note:
//Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?
func missingNumber(nums []int) int {
	expectedSum := 0
	actualSum := 0
	for i := 0; i <= len(nums); i++ {
		expectedSum += i
		if i < len(nums) {
			actualSum += nums[i]
		}
	}
	return expectedSum - actualSum
}
