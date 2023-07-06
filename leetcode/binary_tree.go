package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的最大深度
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := max(MaxDepth(root.Left), MaxDepth(root.Right))
	return depth + 1
}

// 相同的树
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil {
		return p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
	}
	return false
}

// 翻转二叉树
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := InvertTree(root.Left)
	right := InvertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

// 对称二叉树
func IsSymmetric(root *TreeNode) bool {
	return checkSymmetric(root, root)
}

func checkSymmetric(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && checkSymmetric(p.Left, q.Right) && checkSymmetric(p.Right, q.Left)
}
