package other

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/601/
//Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
//
//
//In Pascal's triangle, each number is the sum of the two numbers directly above it.
//
//Example:
//
//Input: 5
//Output:
//[
//[1],
//[1,1],
//[1,2,1],
//[1,3,3,1],
//[1,4,6,4,1]
//]
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	first := []int{1}
	if numRows == 1 {
		return [][]int{first}
	}
	res := make([][]int, numRows)
	res[0] = first
	res[1] = []int{1, 1}
	if numRows == 2 {
		return res
	}

	for i := 2; i < numRows; i++ {
		prevLine := res[i-1]
		res[i] = make([]int, len(prevLine)+1)
		res[i][0] = 1
		for j := 1; j < len(prevLine); j++ {
			res[i][j] = prevLine[j] + prevLine[j-1]
		}
		res[i][len(prevLine)] = 1
	}
	return res
}
