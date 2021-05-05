#  116. Populating Next Right Pointers in Each Node
#  https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
#
#  submission : faster than 60%

class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next


def connect(root: 'Node') -> 'Node':
    levelStart = root
    while levelStart:
        current = levelStart
        while current:
            if current.left:
                current.left.next = current.right
            if current.next and current.right:
                current.right.next = current.next.left
        levelStart = levelStart.left
    return root


