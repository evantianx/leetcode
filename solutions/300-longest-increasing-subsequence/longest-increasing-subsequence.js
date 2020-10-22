// Given an unsorted array of integers, find the length of longest increasing subsequence.
//
// Example:
//
//
// Input: [10,9,2,5,3,7,101,18]
// Output: 4 
// Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4. 
//
// Note: 
//
//
// 	There may be more than one LIS combination, it is only necessary for you to return the length.
// 	Your algorithm should run in O(n2) complexity.
//
//
// Follow up: Could you improve it to O(n log n) time complexity?
//


var lengthOfLIS = function(nums) {
   const dp = []
   
   for (const num of nums) {
       if (!dp.length || num > dp[dp.length - 1]) {
           dp.push(num)           
           continue
       }
       
       const index = binarySearch(dp, num)
       dp[index] = num
   }
    
   return dp.length
};

const binarySearch = (arr, target) => {
    let low = 0
    let high = arr.length - 1
    
    while (low < high) {
        const mid = Math.floor((high - low) / 2) + low
        
        if (arr[mid] >= target) {
            high = mid
        } else {
            low = mid + 1
        }
    }
    
    return low
}
