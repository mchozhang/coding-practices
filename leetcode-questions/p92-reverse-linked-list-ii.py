#!/usr/bin/env python
# -*- coding: utf-8 -*-

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def reverseBetween(head: ListNode, left: int, right: int) -> ListNode:
    i = 1
    current = head
    prev = None
    while i < left:
        prev = current
        current = current.next
        i += 1

    preReversed = prev
    reversedTail = current
    prev = None
    while i <= right:
        next = current.next
        current.next = prev
        prev = current
        current = next
        i += 1

    if preReversed:
        preReversed.next = prev

    reversedTail.next = current
    if left > 1:
        return head
    return prev
