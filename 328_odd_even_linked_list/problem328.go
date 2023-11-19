package problem328

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	lastOdd, firstEven, lastEven := head, head.Next, head.Next
	if lastEven == nil {
		return head
	}

	odd := lastEven.Next
	if odd == nil {
		return head
	}

	even := odd.Next

	for {
		lastOdd.Next = odd
		odd.Next = firstEven
		lastEven.Next = even

		if even == nil {
			break
		}

		lastOdd, lastEven = odd, even

		odd = even.Next
		if odd == nil {
			break
		}

		even = odd.Next
	}

	return head
}
