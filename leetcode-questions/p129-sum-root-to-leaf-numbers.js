/**
 * LeetCode : 129. Sum Root to Leaf Numbers
 * https://leetcode.com/problems/sum-root-to-leaf-numbers/
 *
 * submission : faster than 64%
 */

class TreeNode {
    constructor(val, left, right) {
        this.val = (val === undefined ? 0 : val);
        this.left = (left === undefined ? null : left);
        this.right = (right === undefined ? null : right);
    }
}

/**
 * @param {TreeNode} root
 * @return {number}
 */
var sumNumbers = function (root) {
    let sum = 0

    function traverse(node, base) {
        if (node === null) {
            return
        }
        base = base * 10 + node.val
        if (node.left === null && node.right === null) {
            sum += base
        }
        traverse(node.left, base)
        traverse(node.right, base)
    }

    traverse(root, 0)

    return sum
};