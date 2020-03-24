package tree


/*
题目描述:
给定一个二叉树，返回它的 前序 遍历。
解题思路:
深度优先搜索(DFS)
在这个策略中，我们采用深度作为优先级，以便从跟开始一直到达某个确定的叶子，
然后再返回根到达另一个分支。
深度优先搜索策略又可以根据根节点、左孩子和右孩子的相对顺序被细分为前序遍历，
中序遍历和后序遍历。

广度优先搜索(BFS)
我们按照高度顺序一层一层的访问整棵树，高层次的节点将会比低层次的节点先被访问到。

从根节点开始，每次迭代弹出当前栈顶元素，并将其孩子节点压入栈中，先压右孩子再压左孩子。
在这个算法中，输出到最终结果的顺序按照 Top->Bottom 和 Left->Right，符合前序遍历的顺序。
*/
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack = make([]*TreeNode,1)
	//根节点入栈
	stack[0] = root
	var rtn = make([]int,0)
	for len(stack) != 0 {
		//取出栈尾元素
		e := stack[len(stack) - 1]
		//出栈
		stack = stack[0:len(stack) - 1]
		if e != nil {
			rtn = append(rtn,e.Val)
		}
		//右子节点先压栈
		if e.Right != nil {
			stack = append(stack,e.Right)
		}
		//左子节点后压栈
		if e.Left != nil {
			stack = append(stack,e.Left)
		}
	}
	return rtn
}

/*
二叉树的前序遍历,递归
*/
func preorder(root *TreeNode) []int {
	var rtn = make([]int,0)
	preorderrecursive(root,rtn)
	return rtn
}

func preorderrecursive(root *TreeNode,res []int)  {
	if root != nil {
		res = append(res,root.Val)
		if root.Left != nil {
			preorderrecursive(root.Left,res)
		}
		if root.Right != nil {
			preorderrecursive(root.Right,res)
		}
	}
}
