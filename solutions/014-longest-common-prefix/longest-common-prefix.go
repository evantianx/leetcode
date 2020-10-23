// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//
//  
// Example 1:
//
//
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
//
//
// Example 2:
//
//
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
//
//
//  
// Constraints:
//
//
// 	0 <= strs.length <= 200
// 	0 <= strs[i].length <= 200
// 	strs[i] consists of only lower-case English letters.
//
//


func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
    	return ""
    }

    if len(strs) == 1 {
    	return strs[0]
    }

	substringMap := make(map[string]int)

	var longestSubstring string
	var longestLength int
	// populate the map
	for _, str := range strs {
		for j := 1; j <= len(str); j++ {
			substring := str[:j]
			substringMap[substring] += 1

			if substringMap[substring] == len(strs) && len(substring) > longestLength {
				longestLength = len(substring)
				longestSubstring = substring
			}
		}
	}

	return longestSubstring
}

