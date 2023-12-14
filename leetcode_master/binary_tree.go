package leetcode_master

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// No.102 二叉树的层序遍历
func LevelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) > 0 {
		n := len(q)
		temp := []int{}
		for i := 0; i < n; i++ {
			node := q[i]
			temp = append(temp, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, temp)
		q = q[n:]
	}
	return
}

// No.107 二叉树的层序遍历Ⅱ
func LevelOrderBottom(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) > 0 {
		n := len(q)
		temp := []int{}
		for i := 0; i < n; i++ {
			node := q[i]
			temp = append(temp, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = preInsertSlice(temp, res)
		q = q[n:]
	}
	return
}

func preInsertSlice(i []int, s [][]int) [][]int {
	res := append([][]int{i}, s...)
	return res
}

// No.111 二叉树的最小深度
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) != 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[i]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			if node.Left == nil && node.Right == nil {
				return depth
			}
		}
		depth++
		q = q[n:]
	}
	return depth
}

// No.226 翻转二叉树
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = InvertTree(root.Left)
	root.Right = InvertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

// No.101 对称二叉树
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return checkSymmetric(left.Left, right.Right) && checkSymmetric(left.Right, right.Left)
}

// No.222 完全二叉树的节点个数
func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth, rightDepth := 0, 0
	left, right := root.Left, root.Right
	for left != nil {
		left = left.Left
		leftDepth++
	}
	for right != nil {
		right = right.Right
		rightDepth++
	}
	if leftDepth == rightDepth {
		return (2 << leftDepth) - 1
	}
	return CountNodes(left) + CountNodes(right) + 1
}

// No.110 平衡二叉树
func isBalanced(root *TreeNode) bool {
	return nodeDepth(root) != -1
}

func nodeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth, rightDepth := nodeDepth(root.Left), nodeDepth(root.Right)
	if leftDepth == -1 || rightDepth == -1 {
		return -1
	}
	if leftDepth-rightDepth > 1 || rightDepth-leftDepth > 1 {
		return -1
	}
	return max(leftDepth, rightDepth) + 1
}
