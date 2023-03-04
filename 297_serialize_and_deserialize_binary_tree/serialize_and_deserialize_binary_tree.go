package serialize_and_deserialize_binary_tree

import (
	"fmt"
	"strings"
)

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "!"
	}

	var b strings.Builder
	fmt.Fprintf(&b, "%d.", root.Val)

	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		switch {
		case queue[0] == nil && queue[1] == nil:
			b.WriteByte('_')
		case queue[0] == nil:
			fmt.Fprintf(&b, "!%d.", queue[1].Val)
			queue = append(queue, queue[1].Left, queue[1].Right)
		case queue[1] == nil:
			fmt.Fprintf(&b, "%d.!", queue[0].Val)
			queue = append(queue, queue[0].Left, queue[0].Right)
		default:
			fmt.Fprintf(&b, "%d.%d.", queue[0].Val, queue[1].Val)
			queue = append(queue, queue[0].Left, queue[0].Right, queue[1].Left, queue[1].Right)
		}
		queue = queue[2:]
	}

	return b.String()
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "!" {
		return nil
	}

	root := &TreeNode{}
	queue, node := []*TreeNode{root}, 1
	dLen, val, valIsPositive := len(data), 0, true

	for i := 0; i < dLen; i++ {
		switch {
		case data[i]-'0' < 10:
			val = val*10 + int(data[i]-'0')
		case data[i] == '-':
			valIsPositive = false
		case data[i] == '.':
			if !valIsPositive {
				val = -val
			}
			switch node {
			case 1:
				queue[0].Val, node = val, 2
			case 2:
				queue[0].Left, node = &TreeNode{Val: val}, 3
				queue = append(queue, queue[0].Left)
			case 3:
				queue[0].Right, node = &TreeNode{Val: val}, 2
				queue = append(queue, queue[0].Right)
				queue = queue[1:]
			}
			val, valIsPositive = 0, true
		case data[i] == '!':
			if node == 2 {
				node = 3
			} else {
				queue, node = queue[1:], 2
			}
		case data[i] == '_':
			if len(queue) > 0 {
				queue, node = queue[1:], 2
			} else {
				break
			}
		}
	}

	return root
}
