package array

//https://leetcode.com/problems/count-square-submatrices-with-all-ones/
//1277. Count Square Submatrices with All Ones
//Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.
//
//
//
//Example 1:
//
//Input: matrix =
//[
//[0,1,1,1],
//[1,1,1,1],
//[0,1,1,1]
//]
//Output: 15
//Explanation:
//There are 10 squares of side 1.
//There are 4 squares of side 2.
//There is  1 square of side 3.
//Total number of squares = 10 + 4 + 1 = 15.
//Example 2:
//
//Input: matrix =
//[
//[1,0,1],
//[1,1,0],
//[1,1,0]
//]
//Output: 7
//Explanation:
//There are 6 squares of side 1.
//There is 1 square of side 2.
//Total number of squares = 6 + 1 = 7.
//
//
//Constraints:
//
//1 <= arr.length <= 300
//1 <= arr[0].length <= 300
//0 <= arr[i][j] <= 1

func countSquares(matrix [][]int) int {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return 0
	}

	n := len(matrix)
	m := len(matrix[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		dp[i][0] = matrix[i][0]
	}

	for j := 0; j < m; j++ {
		dp[0][j] = matrix[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				continue
			}

			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
		}
	}

	subMatCount := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			subMatCount += dp[i][j]
		}
	}

	return subMatCount
}

func min(nums ...int) int {
	min := nums[0]
	for _, n := range nums {
		if n > min {
			continue
		}

		min = n
	}
	return min
}
