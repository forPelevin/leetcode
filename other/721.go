package other

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/721/
//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//An input string is valid if:
//
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
//Note that an empty string is also considered valid.
//
//Example 1:
//
//Input: "()"
//Output: true
//Example 2:
//
//Input: "()[]{}"
//Output: true
//Example 3:
//
//Input: "(]"
//Output: false
//Example 4:
//
//Input: "([)]"
//Output: false
//Example 5:
//
//Input: "{[]}"
//Output: true
var closingToOpeningBrackets = map[byte]byte{
	'}': '{',
	']': '[',
	')': '(',
}

func isValid(s string) bool {
	if s == "" {
		return true
	}

	stack := make([]byte, 0, len(s)/2)
	for i := 0; i < len(s); i++ {
		// If this is a closed bracket, check that there is an opened bracket in stack
		if _, ok := closingToOpeningBrackets[s[i]]; ok {
			if len(stack) < 1 {
				return false
			}

			if stack[len(stack)-1] != closingToOpeningBrackets[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}

		// If this is an opened bracket, put it in the stack
		stack = append(stack, s[i])
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
