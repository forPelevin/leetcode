package sortingsearching

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/96/sorting-and-searching/774/
//You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.
//
//Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.
//
//You are given an API bool isBadVersion(version) which will return whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.
//
//Example:
//
//Given n = 5, and version = 4 is the first bad version.
//
//call isBadVersion(3) -> false
//call isBadVersion(5) -> true
//call isBadVersion(4) -> true
//
//Then 4 is the first bad version.
func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	from := 1
	to := n
	mid := 0
	for from <= to {
		mid = (from + to) / 2
		if isBadVersion(mid) {
			if mid == 1 {
				return mid
			}

			if !isBadVersion(mid-1) {
				return mid
			}
			to = mid - 1
			continue
		}
		from = mid + 1
	}
	return mid
}