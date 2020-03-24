package tree

/*
返回二叉树的后序遍历

使用广度优先搜索
*/
func postorderTraversal(root *TreeNode) []int {
	var stack = make([]*TreeNode,0)
	var res = make([]int,0)
	//按广度遍历一遍
	stack = append(stack,root)
	for len(stack) > 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if e != nil {
			res = append(res,e.Val)
			if e.Left != nil {
				stack = append(stack,e.Left)
			}
			if e.Right != nil {
				stack = append(stack,e.Right)
			}
		}
	}
	var l = len(res)
	var mid = (l-1) >> 1
	for i:=0;i<=mid;i++{
		res[i],res[l-i-1] = res[l-i-1],res[i]
	}
	return res
}

func postorder(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		postorder(root.Left)
	}
	if root.Right != nil {
		postorder(root.Right)
	}
	res = append(res,root.Val)
	return
}
