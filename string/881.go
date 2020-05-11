package string

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/881/
//Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.
//
//Examples:
//
//s = "leetcode"
//return 0.
//
//s = "loveleetcode",
//return 2.
//Note: You may assume the string contain only lowercase letters.

const defaultInd = -1

func firstUniqChar(s string) int {
	counter := make([]int, 26)
	for _, ch := range s {
		counter[ch-'a']++
	}
	for i, ch := range s {
		if counter[ch-'a'] == 1 {
			return i
		}
	}
	return defaultInd
}
