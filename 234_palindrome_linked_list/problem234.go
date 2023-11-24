package problem234

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time O(n), space O(1)
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	fast = nil // tail
	for slow != nil {
		slow, slow.Next, fast = slow.Next, fast, slow
	}

	for head != nil && fast != nil {
		if head.Val != fast.Val {
			return false
		}
		head, fast = head.Next, fast.Next
	}

	return true
}
