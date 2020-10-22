// Given a string s, find the length of the longest substring without repeating characters.
//
//  
// Example 1:
//
//
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
//
//
// Example 2:
//
//
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
//
//
// Example 3:
//
//
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
//
//
// Example 4:
//
//
// Input: s = ""
// Output: 0
//
//
//  
// Constraints:
//
//
// 	0 <= s.length <= 5 * 104
// 	s consists of English letters, digits, symbols and spaces.
//
//


func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right + 1 < len(s) && freq[s[right + 1] - 'a'] == 0 {
			freq[s[right + 1] - 'a']++
			right++
		} else {
			freq[s[left] - 'a']--
			left++
		}

		if right - left + 1 > result {
			result = right - left + 1
		}
	}

	return result
	
}

