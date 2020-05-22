package design

import (
	"math/rand"
	"time"
)

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/98/design/670/
//Shuffle a set of numbers without duplicates.
//
//Example:
//
//// Init an array with set 1, 2, and 3.
//int[] nums = {1,2,3};
//Solution solution = new Solution(nums);
//
//// Shuffle the array [1,2,3] and return its result. Any permutation of [1,2,3] must equally likely to be returned.
//solution.shuffle();
//
//// Resets the array back to its original configuration [1,2,3].
//solution.reset();
//
//// Returns the random shuffling of array [1,2,3].
//solution.shuffle();
type Solution struct {
	original []int
}

func Constructor(nums []int) Solution {
	return Solution{
		original: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.original
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, len(this.original))
	copy(nums, this.original)
	l := len(nums)
	for i := range nums {
		randInd := rand.Intn(l)
		nums[randInd], nums[i] = nums[i], nums[randInd]
	}
	return nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
