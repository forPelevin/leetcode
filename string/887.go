package string

import "bytes"

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/887/
//Write a function to find the longest common prefix string amongst an array of strings.
//
//If there is no common prefix, return an empty string "".
//
//Example 1:
//
//Input: ["flower","flow","flight"]
//Output: "fl"
//Example 2:
//
//Input: ["dog","racecar","car"]
//Output: ""
//Explanation: There is no common prefix among the input strings.
//Note:
//
//All given inputs are in lowercase letters a-z.

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	sLen := len(strs[0])
	buf := bytes.NewBuffer(make([]byte, 0, sLen))
Loop:
	for i := 0; i < sLen; i++ {
		b := strs[0][i]
		for j := 1; j < len(strs); j++ {
			currS := strs[j]
			if i >= len(currS) {
				break Loop
			}

			if currS[i] != b {
				break Loop
			}
			if j == len(strs)-1 {
				buf.WriteByte(b)
			}
		}
	}

	return buf.String()
}
