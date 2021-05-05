#  117. Populating Next Right Pointers in Each Node II
#  https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
#
#  submission : faster than 70%

class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next


def connect(root: 'Node') -> 'Node':
    levelStart = root
    while levelStart:
        cur = levelStart
        nextLevelStart = None
        while cur:
            # find start point of next level
            if not nextLevelStart:
                nextLevelStart = cur.left if cur.left else cur.right

            # find the node point to the 
            rightPointer = cur.right 
            if cur.left:
                if cur.right:
                    cur.left.next = cur.right
                else:
                    rightPointer = cur.left
                
            if rightPointer:
                curNext = cur.next
                right = None
                while curNext:
                    if curNext.left:
                        right = curNext.left
                        break
                    elif curNext.right:
                        right = curNext.right
                        break
                    curNext = curNext.next
                rightPointer.next = right
            cur = cur.next
        levelStart = nextLevelStart
    return root
        