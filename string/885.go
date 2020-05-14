package string

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/885/
//Implement strStr().
//
//Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
//Example 1:
//
//Input: haystack = "hello", needle = "ll"
//Output: 2
//Example 2:
//
//Input: haystack = "aaaaa", needle = "bba"
//Output: -1
//Clarification:
//
//What should we return when needle is an empty string? This is a great question to ask during an interview.
//
//For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

func strStr(haystack string, needle string) int {
	nLen := len(needle)
	if nLen == 0 {
		return 0
	}
	hLen := len(haystack)
	if hLen < nLen {
		return -1
	}

	for i := 0; i < hLen; i++ {
		rBound := i + nLen
		if rBound > hLen {
			return -1
		}
		if needle == haystack[i:rBound] {
			return i
		}
	}
	return -1
}