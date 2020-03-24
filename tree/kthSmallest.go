package tree

/*
给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。
中序遍历,然后返回便利数组的第k-1个元素
*/
func kthSmallest(root *TreeNode, k int) int {
	var res = make([]int,0)
	var stack = make([]*TreeNode,0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack,root.Left)
			root = root.Left
		}
		root = stack[len(stack)-1]
		res = append(res,root.Val)
		root = root.Right
		stack = stack[:len(stack)-1]
	}
	return res[k-1]
}


