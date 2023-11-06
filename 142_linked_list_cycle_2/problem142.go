package problem142

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time O(N), space O(1)
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow, fast = slow.Next, fast.Next.Next
	}

	slow, fast = head, fast.Next
	for fast != slow {
		slow, fast = slow.Next, fast.Next
	}

	return slow
}
