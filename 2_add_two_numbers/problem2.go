package problem2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	root := result

	for {
		if l1 != nil {
			result.Val, l1 = result.Val+l1.Val, l1.Next
		}

		if l2 != nil {
			result.Val, l2 = result.Val+l2.Val, l2.Next
		}

		if result.Val > 9 {
			result.Val, result.Next = result.Val-10, &ListNode{Val: 1}
		}

		if l1 == nil && l2 == nil {
			break
		}

		if result.Next == nil {
			result.Next = &ListNode{}
		}

		result = result.Next
	}

	return root
}
