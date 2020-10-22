// Given a string s containing only digits, return all possible valid IP addresses that can be obtained from s. You can return them in any order.
//
// A valid IP address consists of exactly four integers, each integer is between 0 and 255, separated by single dots and cannot have leading zeros. For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses and "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses. 
//
//  
// Example 1:
// Input: s = "25525511135"
// Output: ["255.255.11.135","255.255.111.35"]
// Example 2:
// Input: s = "0000"
// Output: ["0.0.0.0"]
// Example 3:
// Input: s = "1111"
// Output: ["1.1.1.1"]
// Example 4:
// Input: s = "010010"
// Output: ["0.10.0.10","0.100.1.0"]
// Example 5:
// Input: s = "101023"
// Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
//
//  
// Constraints:
//
//
// 	0 <= s.length <= 3000
// 	s consists of digits only.
//
//


/**
 * @param {string} s
 * @return {string[]}
 */
var restoreIpAddresses = function(s) {
  const allIpAddresses = []
  const path = []
  snapShotIp(allIpAddresses, s, 0, path, 0)
  return allIpAddresses
};

const snapShotIp = (allAddresses, s, builderIndex, path, segment) => {
  if (segment === 4 && builderIndex === s.length) {
    allAddresses.push(`${path[0]}.${path[1]}.${path[2]}.${path[3]}`)
  } else if (segment === 4 || builderIndex === s.length) {
    return
  }
  
  for (let len = 1; len <= 3 && builderIndex + len <= s.length; len++) {
    let snapshot = s.substring(builderIndex, builderIndex + len)
    let value = parseInt(snapshot, 10)
    
    if (value > 255 || len >= 2 && s.charAt(builderIndex) === '0') {
      break
    }
    
    path[segment] = value
    snapShotIp(allAddresses, s, builderIndex + len, path, segment + 1)
    path[segment] = -1
  }
}
