// The count-and-say sequence is a sequence of digit strings defined by the recursive formula:
//
//
// 	countAndSay(1) = "1"
// 	countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1), which is then converted into a different digit string.
//
//
// To determine how you "say" an integer, split it into the minimal number of substrings so that each substring is all the same character, then for each substring, say the number of characters, then say the character. To convert the saying into a digit string, replace the counts with a number and concatenate every saying.
//
// For example, the saying and conversion for digit string "3322251":
//
//
// "3322251" -> "33"     +  "222"      +  "5"    +  "1"
//           -> two 3's  +  three 2's  +  one 5  +  one 1
//           -> 2   3    +  3     2    +  1   5  +  1   1
//           -> "23"     +  "32"       +  "15"   +  "11"
//           -> "23321511"
//
//
// Given an integer n, return the nth term of the count-and-say sequence.
//
//  
// Example 1:
//
//
// Input: n = 1
// Output: "1"
// Explanation: This is the base case.
//
//
// Example 2:
//
//
// Input: n = 4
// Output: "1211"
// Explanation:
// countAndSay(1) = "1"
// countAndSay(2) = say "1" = one 1 = "11"
// countAndSay(3) = say "11" = two 1's = "21"
// countAndSay(4) = say "21" = one 2 + one 1 = "12" + "11" = "1211"
//
//
//  
// Constraints:
//
//
// 	1 <= n <= 30
//
//


func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prev := countAndSay(n - 1)
	var say string
	count := 0
	for i := range prev {
		count++
		if i == len(prev) - 1 || prev[i] != prev[i+1] {
			say += strconv.Itoa(count) + prev[i:i+1]
			count = 0
		}
	}

	return say
}

