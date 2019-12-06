# 21. Merge Two Sorted Lists
# https://leetcode.com/problems/merge-two-sorted-lists/
# Merge two sorted linked lists and return it as a new list.
# The new list should be made by splicing together the nodes of the first two lists.
#
# Example:
# Input: 1->2->4, 1->3->4
# Output: 1->1->2->3->4->4
#
# submission: faster than 98.57%


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


def mergeTwoLists(l1: ListNode, l2: ListNode) -> ListNode:
    if not l1:
        return l2

    if not l2:
        return l1

    if l1.val < l2.val:
        head = l1
        current_1 = l1.next
        current_2 = l2
    else:
        head = l2
        current_1 = l1
        current_2 = l2.next

    current_merge = head

    while current_1 or current_2:
        if current_1 and current_2:
            if current_1.val < current_2.val:
                current_merge.next = current_1
                current_1 = current_1.next
            else:
                current_merge.next = current_2
                current_2 = current_2.next

            current_merge = current_merge.next

        elif current_1:
            current_merge.next = current_1
            break
        else:
            current_merge.next = current_2
            break

    return head


def recursive_merge_two_list(l1: ListNode, l2: ListNode) -> ListNode:
    if not l1:
        return l2

    if not l2:
        return l1

    if l1.val < l2.val:
        l1.next = recursive_merge_two_list(l1.next, l2)
        return l1
    else:
        l2.next = recursive_merge_two_list(l1, l2.next)
        return l2


if __name__ == "__main__":
    l1 = ListNode(1)
    l1.next = ListNode(1)
    l1.next.next = ListNode(3)

    l2 = ListNode(1)
    l2.next = ListNode(2)
    l2.next.next = ListNode(5)
    l2.next.next.next = ListNode(6)

    # merge = Solution().mergeTwoLists(l1, l2)
    merge = recursive_merge_two_list(l1, l2)

    while merge:
        print(merge.val)
        merge = merge.next
