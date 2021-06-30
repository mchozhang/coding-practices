# LeetCode : 99. Recover Binary Search Tree
# https://leetcode.com/problems/recover-binary-search-tree/

# submission : faster than 57%

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def recoverTree(root: TreeNode) -> None:
    first, second, prev = None, None, TreeNode(float('-inf'))

    def inorder(node):
        nonlocal first, second, prev
        if not node:
            return
        inorder(node.left)
        if prev.val > node.val:
            first = first or prev
            second = node
        prev = node
        inorder(node.right)

    inorder(root)
    first.val, second.val = second.val, first.val

