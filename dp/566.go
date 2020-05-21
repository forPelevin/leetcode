package dp

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/97/dynamic-programming/566/
//Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
//
//Example:
//
//Input: [-2,1,-3,4,-1,2,1,-5,4],
//Output: 6
//Explanation: [4,-1,2,1] has the largest sum = 6.
//Follow up:
//
//If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
func maxSubArray(nums []int) int {
	currMax := nums[0]
	globalMax := nums[0]
	for i := 1; i < len(nums); i++ {
		currMax = max(currMax + nums[i], nums[i])
		globalMax = max(globalMax, currMax)
	}
	return globalMax
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
