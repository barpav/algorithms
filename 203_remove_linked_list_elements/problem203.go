package problem203

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time O(N), space O(1)
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil || val == 0 {
		return head
	}

	var root, prev *ListNode
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val == val {
			if prev != nil {
				prev.Next = nil
			}
			continue
		}

		if prev != nil {
			prev.Next = cur
		}

		prev = cur

		if root == nil {
			root = cur
		}
	}

	return root
}
