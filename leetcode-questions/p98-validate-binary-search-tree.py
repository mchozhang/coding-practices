# LeetCode : Validate Binary Search Tree
# https://leetcode.com/problems/validate-binary-search-tree/
# Given a binary tree, determine if it is a valid binary search tree (BST).
#
# Example:
# Input: [5,1,4,null,null,3,6]
# Output: false
#
# submission : faster than 84.54%


class TreeNode:
    def __init__(self, x, left=None, right=None):
        self.val = x
        self.left = left
        self.right = right


def isValidBST(root: TreeNode) -> bool:
    def recursiveValidate(root, minVal, maxVal):
        if not root:
            return True
        if root.val <= minVal or root.val >= maxVal:
            return False

        return recursiveValidate(root.left, minVal, root.val) and recursiveValidate(root.right, root.val, maxVal)

    def recursiveValidate2(root):
        if not root:
            return True, None, None

        result, minimum, maximum = True, root.val, root.val
        if root.right:
            result, minRight, maxRight = recursiveValidate2(root.right)
            if not result or minRight <= root.val:
                return False, None, None
            else:
                maximum = maxRight

        if root.left:
            result, minLeft, maxLeft = recursiveValidate2(root.left)
            if not result or maxLeft >= root.val:
                return False, None, None
            else:
                minimum = minLeft

        return result, minimum, maximum

    return recursiveValidate(root, -9999999999, 9999999999)


if __name__ == "__main__":
    node8 = TreeNode(4)
    node7 = TreeNode(1)
    node5 = TreeNode(12)
    node6 = TreeNode(18)

    node4 = TreeNode(6)
    node3 = TreeNode(3, node7, node8)
    node2 = TreeNode(15, node5, node6)
    node1 = TreeNode(5, node3, node4)
    node0 = TreeNode(10, node1, node2)

    print(isValidBST(node0))
