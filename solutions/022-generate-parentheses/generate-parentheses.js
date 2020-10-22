// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
//  
// Example 1:
// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]
// Example 2:
// Input: n = 1
// Output: ["()"]
//
//  
// Constraints:
//
//
// 	1 <= n <= 8
//
//


/**
 * @param {number} n
 * @return {string[]}
 */
var generateParenthesis = function(n, str = '', left = 0, right = 0) {
  let arr = []
  if (str.length === n * 2) return str
  if (left < n) arr = arr.concat(generateParenthesis(n, str + '(', left + 1, right))
  if (right < n && left > right) arr = arr.concat(generateParenthesis(n, str + ')', left, right + 1))
  
  return arr
};
