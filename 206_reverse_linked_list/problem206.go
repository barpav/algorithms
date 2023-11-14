package problem206

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time O(N), space O(1)
func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next, head.Next = head.Next, prev
		prev, head = head, next
	}
	return prev
}
