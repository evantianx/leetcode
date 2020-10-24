// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
//
// Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
//
//  
// Example 1:
//
//
// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
//
//
// Example 2:
//
//
// Input: nums = [1]
// Output: 1
//
//
// Example 3:
//
//
// Input: nums = [0]
// Output: 0
//
//
// Example 4:
//
//
// Input: nums = [-1]
// Output: -1
//
//
// Example 5:
//
//
// Input: nums = [-2147483647]
// Output: -2147483647
//
//
//  
// Constraints:
//
//
// 	1 <= nums.length <= 2 * 104
// 	-231 <= nums[i] <= 231 - 1
//
//


func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	
	return res
}

