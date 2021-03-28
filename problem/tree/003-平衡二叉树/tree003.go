package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalance(root *TreeNode) bool {

	if nil == root {
		return true
	}

	if !isBalance(root.Left) || !isBalance(root.Right) {
		return false
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	if abs(l-r) > 1 {
		return false
	}
	return true
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Right), maxDepth(root.Left)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
