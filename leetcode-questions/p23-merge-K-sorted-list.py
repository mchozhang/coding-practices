#  LeetCode : Merge K sorted List
#  https://leetcode.com/problems/merge-k-sorted-lists/
#  Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
#
#  Example:
#  Input: [1->4->5, 1->3->4, 2->6]
#  Output: 1->1->2->3->4->4->5->6
#
#  submission : faster than 78.87%

import heapq
from queue import PriorityQueue
from typing import List


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

    def __lt__(self, other):
        return self.val < other.val


def mergeKLists(lists: List[ListNode]) -> ListNode:
    """
    based on heap, fastest
    """
    heap = []
    dummy = ListNode(0)
    mergingNode = dummy
    for node in lists:
        if node:
            heapq.heappush(heap, node)

    while len(heap) > 0:
        node = heapq.heappop(heap)
        mergingNode.next = node
        mergingNode = mergingNode.next
        if node.next:
            heapq.heappush(heap, node.next)
    return dummy.next


def mergeKLists2(lists: List[ListNode]) -> ListNode:
    """
    based on priority queue
    """
    queue = PriorityQueue()
    dummy = ListNode(0)
    mergingNode = dummy

    for node in lists:
        if node:
            queue.put(node)

    while not queue.empty():
        node = queue.get()
        mergingNode.next = node
        mergingNode = mergingNode.next
        if node.next:
            queue.put(node.next)

    return dummy.next


def mergeKLists3(lists: List[ListNode]) -> ListNode:
    """
    based on merge sort
    """
    mergeHead = None
    for node in lists:
        mergeHead = merge(mergeHead, node)

    return mergeHead


def merge(node1: ListNode, node2: ListNode) -> ListNode:
    """
    merge 2 sorted linked lists
    """
    dummy = ListNode(0)
    current = dummy
    while node1 and node2:
        if node1.val < node2.val:
            current.next = node1
            node1 = node1.next
        else:
            current.next = node2
            node2 = node2.next
        current = current.next

    if node1:
        current.next = node1

    if node2:
        current.next = node2

    return dummy.next


if __name__ == "__main__":
    l1 = ListNode(1)
    l1.next = ListNode(4)
    l1.next.next = ListNode(5)

    l2 = ListNode(1)
    l2.next = ListNode(3)
    l2.next.next = ListNode(4)

    l3 = ListNode(2)
    l3.next = ListNode(6)

    node = mergeKLists([l1, l2, l3])
    while node:
        print(node.val)
        node = node.next
