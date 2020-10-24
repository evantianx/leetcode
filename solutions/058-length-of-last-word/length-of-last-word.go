// Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word (last word means the last appearing word if we loop from left to right) in the string.
//
// If the last word does not exist, return 0.
//
// Note: A word is defined as a maximal substring consisting of non-space characters only.
//
// Example:
//
//
// Input: "Hello World"
// Output: 5
//
//
//  
//


func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := 0

	for i := len(s) - 1; i > -1; i-- {
		if s[i:i+1] != " " {
			count++
		} else if count > 0 {
			return count
		}
	}

	return count
}

