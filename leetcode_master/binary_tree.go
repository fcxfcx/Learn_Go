package leetcode_master

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
