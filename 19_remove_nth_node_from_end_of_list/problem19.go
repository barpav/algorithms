package problem19

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time O(N), space O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	wrap := &ListNode{Next: head}
	slow, fast := wrap, wrap

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}

	slow.Next = slow.Next.Next

	return wrap.Next
}
