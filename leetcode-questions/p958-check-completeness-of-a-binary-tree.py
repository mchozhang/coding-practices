#  958. Check Completeness of a Binary Tree
#  https://leetcode.com/problems/check-completeness-of-a-binary-tree/
#
#  submission : faster than 58%

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def isCompleteTree(root: TreeNode) -> bool:
    if root is None:
        return True
    queue = [root]
    foundLast = False
    while queue:
        current = queue.pop(0)
        if current is None:
            foundLast = True
        else:
            if foundLast:
                return False
            queue.append(current.left)
            queue.append(current.right)
    return True
