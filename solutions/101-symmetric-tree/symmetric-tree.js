// Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
//
// For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
//
//
//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3
//
//
//  
//
// But the following [1,2,2,null,3,null,3] is not:
//
//
//     1
//    / \
//   2   2
//    \   \
//    3    3
//
//
//  
//
// Follow up: Solve it both recursively and iteratively.
//


/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {boolean}
 */
var isSymmetric = function(root) {
  return root === null || check(root.left, root.right)
};

const check = (leftSub, rightSub) => {
  if (leftSub === null && rightSub === null) {
    return true
  } else if (leftSub !== null && rightSub !== null) {
    return leftSub.val === rightSub.val && check(leftSub.left, rightSub.right) && check(leftSub.right, rightSub.left)
  }
  return false
}
