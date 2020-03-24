package tree

/*
返回二叉树的中序遍历结果
前序: 根左右
中序: 左根右
后序: 左右根
*/
func inorderTraversal(root *TreeNode) []int {
	var res = make([]int,0)
	helper(root,res)
	return res
}

//递归
func helper(root *TreeNode,res []int) {
	if root != nil {
		if root.Left != nil {
			helper(root.Left,res)
		}
		res = append(res,root.Val)
		if root.Right != nil {
			helper(root.Right,res)
		}
	}
}

//迭代,利用辅助栈
func iteration(root *TreeNode ) []int {
	var res,stack = make([]int,0),make([]*TreeNode,0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			//先压栈
			stack = append(stack,root)
			root = root.Left
		}
		//取栈尾元素,先取到的是左子节点,然后是右节点
		root = stack[len(stack)-1]
		//加到遍历顺序中
		res = append(res,root.Val)
		//取右节点
		root = root.Right
		stack = stack[:len(stack)-1]
	}
	return res
}


/*
迭代遍历二叉树时,使用辅助栈来完成时
前序遍历 根左右
后续遍历 左右根
都比较好完成，因为我们遍历时，都是先取根节点，然后再处理左右节点
前序遍历时,我们先把根节点入栈,然后循环的取节点,压栈。我们先取栈尾元素,
直接把该节点的值加到遍历顺序中,然后再把右节点和左节点压到栈中,这样
我们取栈尾元素的时候就是先取到的左子节点
最终的遍历顺序就是根左右

后续遍历时，也是先把根节点入栈,然后循环的取节点,压栈。我们还是先取栈尾元素
直接把该节点的值加入到遍历顺序中,然后再把左节点和右节点压栈,这样我们取栈尾元素
的时候，先取到的就是右子节点,最终的遍历顺序就是根右左
然后把这个遍历顺序反转就是后续遍历

中序遍历就不能这样处理,中序遍历时左根右
那么只能先遍历到左节点的叶子节点,这个过程中,循环压左节点入栈,到底后
再取出栈尾元素,这时候的栈尾元素就是左叶子节点,
*/



