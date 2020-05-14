package string

import (
	"unicode"
)

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/883/
//Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
//
//Note: For the purpose of this problem, we define empty string as valid palindrome.
//
//Example 1:
//
//Input: "A man, a plan, a canal: Panama"
//Output: true
//Example 2:
//
//Input: "race a car"
//Output: false

func isPalindrome(s string) bool {
	r := []rune(s)
	i,j := 0,len(s)-1
	for i < j {
		lCh := r[i]
		if !unicode.IsLetter(lCh) && !unicode.IsNumber(lCh) {
			i++
			continue
		}
		rCh := r[j]
		if !unicode.IsLetter(rCh) && !unicode.IsNumber(rCh) {
			j--
			continue
		}
		lCh = unicode.ToLower(lCh)
		rCh = unicode.ToLower(rCh)
		if lCh != rCh {
			return false
		}
		i++
		j--
	}
	return true
}