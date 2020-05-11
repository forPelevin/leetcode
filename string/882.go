package string

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/882/
//Given two strings s and t , write a function to determine if t is an anagram of s.
//
//Example 1:
//
//Input: s = "anagram", t = "nagaram"
//Output: true
//Example 2:
//
//Input: s = "rat", t = "car"
//Output: false
//Note:
//You may assume the string contains only lowercase alphabets.
//
//Follow up:
//What if the inputs contain unicode characters? How would you adapt your solution to such case?

// Solution for lowercase alphabets
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := make([]int32, 26)
	for _, ch := range s {
		counter[ch-'a']++
	}
	for _, ch := range t {
		counter[ch-'a']--
	}
	for _, c := range counter {
		if c != 0 {
			return false
		}
	}
	return true
}

// Solution for unicode
func isUnicodeAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := make(map[int32]int, len(s))
	for _, ch := range s {
		if _, ok := counter[ch]; !ok {
			counter[ch] = 1
			continue
		}
		counter[ch]++
	}
	for _, ch := range t {
		if _, ok := counter[ch]; !ok {
			counter[ch] = -1
			continue
		}
		counter[ch]--
	}
	for _, c := range counter {
		if c != 0 {
			return false
		}
	}
	return true
}

