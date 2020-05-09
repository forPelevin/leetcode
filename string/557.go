package string

import (
	"bytes"
)

//https://leetcode.com/problems/reverse-words-in-a-string-iii/
//557. Reverse Words in a String III
//
//Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.
//
//Example 1:
//Input: "Let's take LeetCode contest"
//Output: "s'teL ekat edoCteeL tsetnoc"
//Note: In the string, each word is separated by single space and there will not be any extra space in the string.

func reverseWords(s string) string {
	if len(s) <= 1 {
		return s
	}

	buf := bytes.NewBuffer(make([]byte, 0, len(s)))
	from, to := 0,0
	for i, ch := range s {
		if ch != ' ' {
			to = i + 1
			if i != len(s) - 1 {
				continue
			}
		}

		word := []byte(s[from:to])
		if len(word) < 1 {
			continue
		}

		if buf.Len() != 0 {
			buf.WriteString(" ")
		}
		for j := len(word) - 1; j >= 0; j-- {
			buf.WriteByte(word[j])
		}
		from = i + 1
	}
	return buf.String()
}

