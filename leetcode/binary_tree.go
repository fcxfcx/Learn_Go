package leetcode

import "math"

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

// 从前序与中序遍历序列构造二叉树
func BuildTreeOne(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	left := BuildTreeOne(preorder[1:(i+1)], inorder[:i])
	right := BuildTreeOne(preorder[(i+1):], inorder[(i+1):])
	root.Left = left
	root.Right = right
	return root
}

// 从中序与后序遍历序列构造二叉树
func BuildTreeTwo(inorder []int, postorder []int) *TreeNode {
	post_length := len(postorder)
	if post_length == 0 {
		return nil
	}
	root := &TreeNode{
		Val: postorder[post_length-1],
	}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	root.Right = BuildTreeTwo(inorder[(i+1):], postorder[i:(post_length-1)])
	root.Left = BuildTreeTwo(inorder[:i], postorder[:i])
	return root
}

// 二叉树展开为链表
func Flatten(root *TreeNode) {
	var pre func(node *TreeNode) []*TreeNode
	pre = func(node *TreeNode) []*TreeNode {
		result := make([]*TreeNode, 0)
		if node != nil {
			result = append(result, node)
			result = append(result, pre(node.Left)...)
			result = append(result, pre(node.Right)...)
		}
		return result
	}
	list := pre(root)
	for i := 1; i < len(list); i++ {
		p, c := list[i-1], list[i]
		p.Left, p.Right = nil, c
	}
}

// 路径总合
func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return HasPathSum(root.Left, targetSum-root.Val) || HasPathSum(root.Right, targetSum-root.Val)
}

// 求根节点到叶节点数字之和
func SumNumbers(root *TreeNode) int {
	var dfs func(root *TreeNode, preSum int) int
	dfs = func(root *TreeNode, preSum int) int {
		if root == nil {
			return 0
		}
		sum := preSum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return sum
		}
		return dfs(root.Left, sum) + dfs(root.Right, sum)
	}
	return dfs(root, 0)
}

// 二叉树中的最大路径和
func MaxPathSum(root *TreeNode) int {
	max_Sum := math.MinInt
	var scan func(root *TreeNode) int
	scan = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := scan(root.Left)
		right := scan(root.Right)
		max_Sum = max(max_Sum, root.Val+left+right)
		output_max := root.Val + max(left, right)
		return max(output_max, 0)
	}
	scan(root)
	return max_Sum
}

// 二叉树中序遍历迭代器
type BSTIterator struct {
	stack []*TreeNode
	cur   *TreeNode
}

func BSTConstructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		cur: root,
	}
}

func (it *BSTIterator) Next() int {
	for node := it.cur; node != nil; node = node.Left {
		it.stack = append(it.stack, node)
	}
	it.cur, it.stack = it.stack[len(it.stack)-1], it.stack[:len(it.stack)-1]
	val := it.cur.Val
	it.cur = it.cur.Right
	return val
}

func (it *BSTIterator) HasNext() bool {
	return it.cur != nil || len(it.stack) != 0
}
