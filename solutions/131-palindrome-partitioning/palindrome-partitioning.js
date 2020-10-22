// Given a string s, partition s such that every substring of the partition is a palindrome.
//
// Return all possible palindrome partitioning of s.
//
// Example:
//
//
// Input: "aab"
// Output:
// [
//   ["aa","b"],
//   ["a","a","b"]
// ]
//
//


/**
 * @param {string} s
 * @return {string[][]}
 */
var partition = function(s) {
  const validDecompositions = []
  const decompInProgress = []
   
  decomposeString(s, 0, decompInProgress, validDecompositions)
  
  return validDecompositions
};

const decomposeString = (s, buildPointer, decompInProgress, validDecompositions) => {
  if (buildPointer === s.length) {
    validDecompositions.push([...decompInProgress])
    return
  } else {
    for (let i = buildPointer; i < s.length; i++) {    
      const subString = s.substring(buildPointer, i+1)
      if (isPalindrome(subString)) {
        decompInProgress.push(subString)
        decomposeString(s, i+1, decompInProgress, validDecompositions)
        decompInProgress.pop()
      }
    }
  }
}

const isPalindrome = str => str === str.split('').reverse().join('')
