// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
//
// 	Open brackets must be closed by the same type of brackets.
// 	Open brackets must be closed in the correct order.
//
//
//  
// Example 1:
//
//
// Input: s = "()"
// Output: true
//
//
// Example 2:
//
//
// Input: s = "()[]{}"
// Output: true
//
//
// Example 3:
//
//
// Input: s = "(]"
// Output: false
//
//
// Example 4:
//
//
// Input: s = "([)]"
// Output: false
//
//
// Example 5:
//
//
// Input: s = "{[]}"
// Output: true
//
//
//  
// Constraints:
//
//
// 	1 <= s.length <= 104
// 	s consists of parentheses only '()[]{}'.
//
//


func isValid(s string) bool {
    if len(s) == 0 {
    	return true
    }

    stack := make([]rune, 0)

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else if ((v == ')') && len(stack) > 0 && stack[len(stack) - 1] == '(') ||
			((v == ']') && len(stack) > 0 && stack[len(stack) - 1] == '[') ||
			((v == '}') && len(stack) > 0 && stack[len(stack) - 1] == '{') {
			stack = stack[:len(stack) - 1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

