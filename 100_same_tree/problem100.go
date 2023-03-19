package problem100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	switch {
	case p == q:
		return true
	case p != nil && q != nil && p.Val == q.Val:
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	default:
		return false
	}
}
