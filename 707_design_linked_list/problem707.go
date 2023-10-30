package problem707

// Singly linked list
type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{val: -1}
}

// time O(N), space O(1)
func (l *MyLinkedList) Get(index int) int {
	node := l
	for i := 0; i < index; i++ {
		if node.next == nil {
			return -1
		}
		node = node.next
	}
	return node.val
}

// time O(1), space O(1)
func (l *MyLinkedList) AddAtHead(val int) {
	if l.val == -1 {
		l.val = val
		return
	}

	l.val, l.next = val, &MyLinkedList{val: l.val, next: l.next}
}

// time O(N), space O(1)
func (l *MyLinkedList) AddAtTail(val int) {
	node := l
	for node.next != nil {
		node = node.next
	}

	if node == l && l.val == -1 {
		l.val = val
		return
	}

	node.next = &MyLinkedList{val: val}
}

// time O(N), space O(1)
func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		l.AddAtHead(val)
		return
	}

	var prev *MyLinkedList
	node := l
	for i := 0; i < index; i++ {
		if node.next == nil && index > i+1 {
			return
		}
		prev, node = node, node.next
	}

	prev.next = &MyLinkedList{val: val, next: node}
}

// time O(N), space O(1)
func (l *MyLinkedList) DeleteAtIndex(index int) {
	var prev *MyLinkedList
	node := l
	for i := 0; i < index; i++ {
		if node.next == nil {
			return
		}
		prev, node = node, node.next
	}

	if prev != nil {
		prev.next = node.next
		return
	}

	// deleting head node
	if l.next != nil {
		*l = *l.next
	} else {
		l.val = -1
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
