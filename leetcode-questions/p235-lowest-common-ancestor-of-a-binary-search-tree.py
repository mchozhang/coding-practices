# LeetCode : 235. Lowest Common Ancestor of a Binary Search Tree
# https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

# submission : faster than 53%

def lowestCommonAncestor(root: 'TreeNode', p: 'TreeNode', q: 'TreeNode'):
    if not root and p.val == root.val or q.val == root.val:
        return root
    if p.val > q.val:
        p, q = q, p

    if p.val < root.val < q.val:
        return root
    elif q.val < root.val:
        return lowestCommonAncestor(root.left, p, q)
    else:
        return lowestCommonAncestor(root.right, p, q)
