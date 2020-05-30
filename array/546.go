package array

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/546/
//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//Example:
//
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].
func twoSum(nums []int, target int) []int {
	indexBySecondNum := make(map[int]int,0)
	for i,v := range nums {
		secondNumInd, exist := indexBySecondNum[v]
		if exist {
			return []int{secondNumInd, i}
		}
		indexBySecondNum[target-v] = i
	}
	return nil
}
