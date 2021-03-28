package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Right = left
	root.Left = right

	return root
}
