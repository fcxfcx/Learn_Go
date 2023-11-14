package top_100_liked

import "math"

func SortList(head *ListNode) *ListNode {
	return SortOne(head, nil)
}

func SortOne(head *ListNode, tail *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	// 二分法，寻找中点
	fast, slow := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != tail {
			fast = fast.Next
		}
	}
	mid := slow
	mid.Next = nil
	return CombineTwoList(SortOne(head, mid), SortOne(mid, tail))
}

func CombineTwoList(h1, h2 *ListNode) *ListNode {
	// 合并两个排序列表
	dummy := &ListNode{}
	temp, temp1, temp2 := dummy, h1, h2
	for temp1 != nil && temp2 != nil {
		if temp1.Val > temp2.Val {
			temp.Next = temp2
			temp2 = temp2.Next
		} else {
			temp.Next = temp1
			temp1 = temp1.Next
		}
		temp = temp.Next
	}
	// 链接剩余部分
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummy.Next
}

// 合并k个升序链表
func MergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	} else if length == 1 {
		return lists[0]
	}
	left := MergeKLists(lists[:length/2])
	right := MergeKLists(lists[length/2:])
	return CombineTwoList(left, right)
}

// LRU
type LRUCache struct {
	Capacity  int
	DummyHead *LRUNode
	DummyTail *LRUNode
	Hash      map[int]*LRUNode
}

type LRUNode struct {
	Next *LRUNode
	Pre  *LRUNode
	Val  int
	Key  int
}

func Constructor(capacity int) LRUCache {
	dummyHead, dummyTail := &LRUNode{}, &LRUNode{}
	dummyHead.Next = dummyTail
	dummyTail.Pre = dummyHead
	return LRUCache{
		Capacity:  capacity,
		DummyHead: dummyHead,
		DummyTail: dummyTail,
		Hash:      make(map[int]*LRUNode),
	}
}

func (lc *LRUCache) DeleteHead() {
	head := lc.DummyHead.Next
	lc.DummyHead.Next = head.Next
	lc.DummyHead.Next.Pre = lc.DummyHead
	head.Next = nil
	head.Pre = nil
	delete(lc.Hash, head.Key)
}

func (lc *LRUCache) MoveToTail(node *LRUNode) {
	tempPre, tempNext := node.Pre, node.Next
	tempPre.Next = tempNext
	tempNext.Pre = tempPre
	tail := lc.DummyTail.Pre
	tail.Next = node
	node.Pre = tail
	node.Next = lc.DummyTail
	lc.DummyTail.Pre = node
}

func (lc *LRUCache) AddToTail(node *LRUNode) {
	tail := lc.DummyTail.Pre
	tail.Next = node
	node.Pre = tail
	node.Next = lc.DummyTail
	lc.DummyTail.Pre = node
}

func (lc *LRUCache) Get(key int) int {
	if node, ok := lc.Hash[key]; ok {
		// 被索引后，放置到队尾
		lc.MoveToTail(node)
		return node.Val
	} else {
		return -1
	}
}

func (lc *LRUCache) Put(key int, value int) {
	if node, ok := lc.Hash[key]; ok {
		node.Val = value
		lc.MoveToTail(node)
	} else {
		if len(lc.Hash) == lc.Capacity {
			// 超出容量，删除队头
			lc.DeleteHead()
		}
		// 新节点加入队尾
		newNode := &LRUNode{
			Val: value,
			Key: key,
		}
		lc.Hash[key] = newNode
		lc.AddToTail(newNode)
	}
}

// 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的中序遍历
func InorderTraversal(root *TreeNode) []int {
	// 使用非递归的方法，用栈模拟递归
	if root == nil {
		return []int{}
	}
	result, stack := []int{}, []*TreeNode{}
	stack = append(stack, root)
	temp := root.Left
	for len(stack) != 0 || temp != nil {
		if temp != nil {
			stack = append(stack, temp)
			temp = temp.Left
		} else {
			// 如果左侧到底，则弹出栈顶元素
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, top.Val)
			temp = top.Right
		}
	}
	return result
}

// 二叉树的最大深度
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := max(MaxDepth(root.Left), MaxDepth(root.Right))
	return depth + 1
}

// 翻转二叉树
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertedLeft := InvertTree(root.Left)
	invertedRight := InvertTree(root.Right)
	root.Right = invertedLeft
	root.Left = invertedRight
	return root
}

// 对称二叉树
func IsSymmetric(root *TreeNode) bool {
	return CheckSymmetric(root, root)
}

func CheckSymmetric(p, q *TreeNode) bool {
	// 检查左右两边是否对称
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && CheckSymmetric(p.Left, q.Right) && CheckSymmetric(p.Right, q.Left)
}

// 二叉树的直径
func DiameterOfBinaryTree(root *TreeNode) int {
	result := 0
	var maxPathNode func(root *TreeNode) int
	maxPathNode = func(root *TreeNode) int {
		if root == nil {
			return -1
		}
		leftMax := maxPathNode(root.Left)
		rightMax := maxPathNode(root.Right)
		if result < leftMax+rightMax+1 {
			result = leftMax + rightMax + 1
		}
		if leftMax > rightMax {
			return leftMax + 1
		} else {
			return rightMax + 1
		}
	}
	maxPathNode(root)
	return result
}

// 二叉树层序遍历
func LevelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	curLevel, nextLevel := []*TreeNode{}, []*TreeNode{}
	curLevel = append(curLevel, root)
	for len(curLevel) != 0 {
		temp := []int{}
		for len(curLevel) != 0 {
			temp = append(temp, curLevel[0].Val)
			if curLevel[0].Left != nil {
				nextLevel = append(nextLevel, curLevel[0].Left)
			}
			if curLevel[0].Right != nil {
				nextLevel = append(nextLevel, curLevel[0].Right)
			}
			curLevel = curLevel[1:]
		}
		res = append(res, temp)
		curLevel = nextLevel
		nextLevel = []*TreeNode{}
	}
	return
}

// 验证二叉搜索树
func IsValidBST(root *TreeNode) bool {
	var check func(root *TreeNode, max int, min int) bool
	check = func(root *TreeNode, max int, min int) bool {
		if root == nil {
			return true
		}
		if root.Val <= max && root.Val >= min {
			if !check(root.Left, root.Val-1, min) || !check(root.Right, max, root.Val+1) {
				return false
			}
			return true
		}
		return false
	}
	return check(root, math.MaxInt64, math.MinInt64)
}

// 二叉搜索树中第k小的元素
func KthSmallest(root *TreeNode, k int) int {
	order := InorderTraversal(root)
	return order[k-1]
}

// 二叉树的右视图
func RightSideView(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	// 层序遍历
	for len(queue) != 0 {
		// 当前层
		n := len(queue)
		temp := &TreeNode{}
		for i := 0; i < n; i++ {
			temp = queue[i]
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		queue = queue[n:]
		res = append(res, temp.Val)
	}
	return
}

// 二叉树展开为链表
func Flatten(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			next := cur.Left
			findRight := next
			for findRight.Right != nil {
				findRight = findRight.Right
			}
			findRight.Right = cur.Right
			cur.Left, cur.Right = nil, next
		}
		cur = cur.Right
	}
}

// 从前序与中序遍历序列构造二叉树
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	} else if len(preorder) == 1 && len(inorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}
	}
	rootVal := preorder[0]
	left := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			left = i
			break
		}
	}
	leftNode := BuildTree(preorder[1:left+1], inorder[:left])
	rightNode := BuildTree(preorder[left+1:], inorder[left+1:])
	root := &TreeNode{
		Val:   rootVal,
		Left:  leftNode,
		Right: rightNode,
	}
	return root
}

// 路径总和Ⅲ
func PathSum(root *TreeNode, targetSum int) int {
	total := 0
	if root == nil {
		return total
	}
	hash := map[int]int{}
	hash[0] = 1
	var findPrefix func(node *TreeNode, prefix int)
	findPrefix = func(node *TreeNode, prefix int) {
		if val, ok := hash[prefix-targetSum]; ok {
			total += val
		}
		hash[prefix]++
		if node.Left != nil {
			findPrefix(node.Left, prefix+node.Left.Val)
		}
		if node.Right != nil {
			findPrefix(node.Right, prefix+node.Right.Val)
		}
		hash[prefix]--
	}
	findPrefix(root, root.Val)
	return total
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
		return root
	}
	if left != nil {
		return left
	}
	return right
}

// 二叉树最大路径和
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxPath := math.MinInt64
	var maxPathOfNode func(node *TreeNode) int
	maxPathOfNode = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMax := maxPathOfNode(node.Left)
		rightMax := maxPathOfNode(node.Right)
		maxPath = max(maxPath, leftMax+rightMax+node.Val)
		outputMax := max(leftMax, rightMax) + node.Val
		return max(outputMax, 0)
	}
	maxPathOfNode(root)
	return maxPath
}
