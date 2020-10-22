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


/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
  const leftChars = []
  for (let i = 0; i < s.length; i++) {
    const c = s.charAt(i)
    if (c === '(' || c === '[' || c === '{') {
      leftChars.push(c)
    } else {
      const popC = leftChars.pop()
      if (c === ')' && popC !== '(' 
          || c === ']' && popC !=='['
          || c === '}' && popC !=='{') {
        return false
      }
    }
  }
  return leftChars.length === 0
};
