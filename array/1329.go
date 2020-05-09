package array

import "sort"

//https://leetcode.com/problems/sort-the-matrix-diagonally/
//Given a m * n matrix mat of integers, sort it diagonally in ascending order from the top-left to the bottom-right then return the sorted array.
//
//
//
//Example 1:
//
//
//Input: mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
//Output: [[1,1,1,1],[1,2,2,2],[1,2,3,3]]
//
//
//Constraints:
//
//m == mat.length
//n == mat[i].length
//1 <= m, n <= 100
//1 <= mat[i][j] <= 100

func diagonalSort(mat [][]int) [][]int {
	if len(mat) <= 1 {
		return mat
	}

	n := len(mat)
	m := len(mat[0])

	tmp := make([]int, 0, n)
	for i := 0; i < m; i++ {
		col := i
		row := 0
		tmp = tmp[:0]

		for row < n && col < m {
			tmp = append(tmp, mat[row][col])
			row++
			col++
		}

		sort.Ints(tmp)

		col = i
		row = 0
		k := 0
		for row < n && col < m {
			mat[row][col] = tmp[k]
			row++
			col++
			k++
		}
	}

	for j := 1; j < n; j++ {
		col := 0
		row := j
		tmp = tmp[:0]

		for row < n && col < m {
			tmp = append(tmp, mat[row][col])
			row++
			col++
		}

		sort.Ints(tmp)

		col = 0
		row = j
		k := 0
		for row < n && col < m {
			mat[row][col] = tmp[k]
			row++
			col++
			k++
		}
	}

	return mat
}