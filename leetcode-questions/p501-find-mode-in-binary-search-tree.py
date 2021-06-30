# LeetCode : Find Mode in Binary Search Tree
# https://leetcode.com/problems/find-mode-in-binary-search-tree/
#
# submission : faster than 88%

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def findMode(root):
    res = []
    count = 0
    maximum = 0
    prev = TreeNode(int(float('-inf')))

    def inorder(node):
        if node:
            nonlocal res, count, prev, maximum
            inorder(node.left)
            if node.val == prev.val:
                count += 1
            else:
                count = 1
            prev = node
            if count == maximum:
                res.append(node.val)
            elif count > maximum:
                maximum = count
                res = [node.val]
            inorder(node.right)

    inorder(root)
    return res

