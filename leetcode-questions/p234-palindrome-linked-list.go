/**
 * LeetCode : Palindromic Substrings
 * https://leetcode.com/problems/palindromic-substrings/
 *
 * submission : faster than 53%
 */
package main

func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	var pre *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}
	if fast != nil {
		slow = slow.Next
	}
	for pre != nil {
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}
	return true
}