package array

//https://leetcode.com/problems/queens-that-can-attack-the-king/
//1222. Queens That Can Attack the King
//On an 8x8 chessboard, there can be multiple Black Queens and one White King.
//
//Given an array of integer coordinates queens that represents the positions of the Black Queens, and a pair of coordinates king that represent the position of the White King, return the coordinates of all the queens (in any order) that can attack the King.
//
//
//
//Example 1:
//
//
//
//Input: queens = [[0,1],[1,0],[4,0],[0,4],[3,3],[2,4]], king = [0,0]
//Output: [[0,1],[1,0],[3,3]]
//Explanation:
//The queen at [0,1] can attack the king cause they're in the same row.
//The queen at [1,0] can attack the king cause they're in the same column.
//The queen at [3,3] can attack the king cause they're in the same diagnal.
//The queen at [0,4] can't attack the king cause it's blocked by the queen at [0,1].
//The queen at [4,0] can't attack the king cause it's blocked by the queen at [1,0].
//The queen at [2,4] can't attack the king cause it's not in the same row/column/diagnal as the king.
//Example 2:
//
//
//
//Input: queens = [[0,0],[1,1],[2,2],[3,4],[3,5],[4,4],[4,5]], king = [3,3]
//Output: [[2,2],[3,4],[4,4]]
//Example 3:
//
//
//
//Input: queens = [[5,6],[7,7],[2,1],[0,7],[1,6],[5,1],[3,7],[0,3],[4,0],[1,2],[6,3],[5,0],[0,4],[2,2],[1,1],[6,4],[5,4],[0,0],[2,6],[4,5],[5,2],[1,4],[7,5],[2,3],[0,5],[4,2],[1,0],[2,7],[0,1],[4,6],[6,1],[0,6],[4,3],[1,7]], king = [3,4]
//Output: [[2,3],[1,4],[1,6],[3,7],[4,3],[5,4],[4,5]]
//
//
//Constraints:
//
//1 <= queens.length <= 63
//queens[0].length == 2
//0 <= queens[i][j] < 8
//king.length == 2
//0 <= king[0], king[1] < 8
//At most one piece is allowed in a cell.

const (
	n = 8
	m = 8
)

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	if len(queens) == 0 || len(king) == 0 {
		return nil
	}

	attackQueens := make([][]int, 0, 8)
	lQ := fetchQueenFromDirection(king, queens, func(i,j int) (int, int) {
		return i - 1, j
	})
	if lQ != nil {
		attackQueens = append(attackQueens, lQ)
	}

	rQ := fetchQueenFromDirection(king, queens, func(i,j int) (int, int) {
		return i + 1, j
	})
	if rQ != nil {
		attackQueens = append(attackQueens, rQ)
	}

	tQ := fetchQueenFromDirection(king, queens, func(i,j int) (int, int) {
		return i, j + 1
	})
	if tQ != nil {
		attackQueens = append(attackQueens, tQ)
	}

	bQ := fetchQueenFromDirection(king, queens, func(i,j int) (int, int) {
		return i, j - 1
	})
	if bQ != nil {
		attackQueens = append(attackQueens, bQ)
	}

	rDiagDown := fetchQueenFromDirection(king, queens, func(i,j int) (int, int) {
		return i + 1, j + 1
	})
	if rDiagDown != nil {
		attackQueens = append(attackQueens, rDiagDown)
	}

	rDiagUp := fetchQueenFromDirection(king, queens, func(i,j int) (int, int) {
		return i - 1, j + 1
	})
	if rDiagUp != nil {
		attackQueens = append(attackQueens, rDiagUp)
	}

	lDiagUp := fetchQueenFromDirection(king, queens, func(i,j int) (int, int) {
		return i - 1, j - 1
	})
	if lDiagUp != nil {
		attackQueens = append(attackQueens, lDiagUp)
	}

	lDiagDown := fetchQueenFromDirection(king, queens, func(i,j int) (int, int) {
		return i + 1, j - 1
	})
	if lDiagDown != nil {
		attackQueens = append(attackQueens, lDiagDown)
	}

	return attackQueens
}

type directionFunc = func(i,j int) (int, int)

func fetchQueenFromDirection(k []int, queens [][]int, dirFn directionFunc) []int {
	i, j := dirFn(k[0],k[1])
	for abs(i) < n && abs(j) < m {
		for qI, q := range queens {
			if q[0] == i && q[1] == j {
				queens = append(queens[:qI], queens[qI+1:]...)
				return q
			}
		}
		i,j = dirFn(i,j)
	}

	return nil
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}