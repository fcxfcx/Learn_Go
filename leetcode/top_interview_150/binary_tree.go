package top_interview_150

import (
	"math"
)

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

// 二叉树的最近公共祖先
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		// 左右都不为空，则只有一种可能即p和q分居左右两侧
		return root
	}
	if left != nil {
		return left
	}
	return right
}

// 二叉树的右视图
func RightSideView(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	cur_stack := make([]*TreeNode, 0)
	next_stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	cur_stack = append(cur_stack, root)
	for len(cur_stack) != 0 {
		result = append(result, cur_stack[len(cur_stack)-1].Val)
		for _, node := range cur_stack {
			if node.Left != nil {
				next_stack = append(next_stack, node.Left)
			}
			if node.Right != nil {
				next_stack = append(next_stack, node.Right)
			}
		}
		cur_stack = next_stack
		next_stack = make([]*TreeNode, 0)
	}
	return result
}

// 二叉树的层平均值
func AverageOfLevels(root *TreeNode) []float64 {
	result := make([]float64, 0)
	cur_stack := make([]*TreeNode, 0)
	next_stack := make([]*TreeNode, 0)
	cur_stack = append(cur_stack, root)
	temp, tempCount := 0, 0
	for len(cur_stack) != 0 {
		for _, node := range cur_stack {
			temp += node.Val
			tempCount += 1
			if node.Left != nil {
				next_stack = append(next_stack, node.Left)
			}
			if node.Right != nil {
				next_stack = append(next_stack, node.Right)
			}
		}
		result = append(result, float64(temp)/float64(tempCount))
		temp = 0
		tempCount = 0
		cur_stack = next_stack
		next_stack = make([]*TreeNode, 0)
	}
	return result
}

// 二叉树的层序遍历
func LevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		temp := make([]int, 0)
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[i]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, temp)
		queue = queue[n:]
	}
	return result
}

// 二叉树的锯齿形层序遍历
func ZigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	leftToRight := true
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		n := len(queue)
		temp := make([]int, 0)
		for i := 0; i < n; i++ {
			node := queue[i]
			if leftToRight {
				temp = append(temp, node.Val)
			} else {
				temp = append(temp, queue[n-1-i].Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, temp)
		queue = queue[n:]
		leftToRight = !leftToRight
	}
	return result
}

// 二叉搜索树的最小绝对差
func GetMinimumDifference(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

// 二叉搜索树第k小的元素
func KthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		stack, root = stack[:len(stack)-1], stack[len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

// 验证二叉搜索树
func IsValidBST(root *TreeNode) bool {
	var check func(root *TreeNode, min int, max int) bool
	check = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}
		if root.Val >= min && root.Val <= max {
			if !check(root.Left, min, root.Val-1) || !check(root.Right, root.Val+1, max) {
				return false
			}
			return true
		}
		return false
	}
	return check(root, math.MinInt64, math.MaxInt64)
}
