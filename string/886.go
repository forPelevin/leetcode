package string

import "bytes"

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/886/
//The count-and-say sequence is the sequence of integers with the first five terms as following:
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//1 is read off as "one 1" or 11.
//11 is read off as "two 1s" or 21.
//21 is read off as "one 2, then one 1" or 1211.
//
//Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence. You can do so recursively, in other words from the previous member read off the digits, counting the number of digits in groups of the same digit.
//
//Note: Each term of the sequence of integers will be represented as a string.
//
//
//
//Example 1:
//
//Input: 1
//Output: "1"
//Explanation: This is the base case.
//Example 2:
//
//Input: 4
//Output: "1211"
//Explanation: For n = 3 the term was "21" in which we have two groups "2" and "1", "2" can be read as "12

func countAndSay(n int) string {
	buf := bytes.NewBufferString("1")
	for n > 1 {
		subBuf := bytes.NewBuffer(make([]byte, 0, 100))
		bs := buf.Bytes()
		for i := 0; i < len(bs); i++ {
			count := 1
			for i+1 < len(bs) && bs[i] == bs[i+1] {
				i++
				count++
			}
			subBuf.WriteByte(byte(count) + '0')
			subBuf.WriteByte(bs[i])
		}
		buf = subBuf
		n--
	}
	return buf.String()
}
