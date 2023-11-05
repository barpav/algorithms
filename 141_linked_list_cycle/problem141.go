package problem141

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time O(N), space O(1)
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow, fast = slow.Next, fast.Next.Next
	}

	return true
}
