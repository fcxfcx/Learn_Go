package leetcode_master

import (
	"math"
	"strconv"
)

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
func IsBalanced(root *TreeNode) bool {
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

// No.257 二叉树的所有路径
func BinaryTreePaths(root *TreeNode) []string {
	res := []string{}
	var travel func(node *TreeNode, temp string)
	travel = func(node *TreeNode, temp string) {
		if node.Left == nil && node.Right == nil {
			s := temp + strconv.Itoa(node.Val)
			res = append(res, s)
			return
		}
		temp = temp + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			travel(node.Left, temp)
		}
		if root.Right != nil {
			travel(node.Right, temp)
		}
	}
	travel(root, "")
	return res
}

// No.404 左叶子之和
func SumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftVal := SumOfLeftLeaves(root.Left)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftVal = root.Left.Val
	}
	rightVal := SumOfLeftLeaves(root.Right)
	return leftVal + rightVal
}

// No.513 找树左下角的值
func FindBottomLeftValue(root *TreeNode) int {
	q := []*TreeNode{}
	q = append(q, root)
	var res int
	for len(q) != 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[i]
			if i == 0 {
				res = node.Val
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[n:]
	}
	return res
}

// 117. 路径之和
func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	left := HasPathSum(root.Left, targetSum-root.Val)
	right := HasPathSum(root.Right, targetSum-root.Val)
	return left || right
}

// 106. 从中序和后序遍历序列构造二叉树
func BuildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	midVal, midIndex := postorder[n-1], 0
	for i := 0; i < n; i++ {
		if inorder[i] == midVal {
			midIndex = i
		}
	}
	left := BuildTree(inorder[0:midIndex], postorder[0:midIndex])
	right := BuildTree(inorder[midIndex+1:], postorder[midIndex:n-1])
	return &TreeNode{
		Val:   midVal,
		Left:  left,
		Right: right,
	}
}

// No.654 最大二叉树
func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	maxIndex := 0
	for i := 0; i < n; i++ {
		if nums[i] >= nums[maxIndex] {
			maxIndex = i
		}
	}
	leftTree := ConstructMaximumBinaryTree(nums[0:maxIndex])
	var rightTree *TreeNode
	if maxIndex != n {
		rightTree = ConstructMaximumBinaryTree(nums[maxIndex+1:])
	}
	return &TreeNode{
		Val:   nums[maxIndex],
		Left:  leftTree,
		Right: rightTree,
	}
}

// No.617 合并二叉树
func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = MergeTrees(root1.Left, root2.Left)
	root1.Right = MergeTrees(root1.Right, root2.Right)
	return root1
}

// No.700 二叉搜索树的搜索
func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val < val {
		return SearchBST(root.Right, val)
	} else {
		return SearchBST(root.Left, val)
	}
}

// No.98 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var valid func(node *TreeNode, min int, max int) bool
	valid = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val >= max || node.Val <= min {
			return false
		}
		left := valid(node.Left, min, node.Val)
		right := valid(node.Right, root.Val, max)
		return left && right
	}
	return valid(root, math.MinInt, math.MaxInt)
}
