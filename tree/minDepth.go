package tree

/*
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
*/
var resMinDepth int
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	resMinDepth = int(^uint(0) >> 1)
	depth(root,0)
	return resMinDepth
}

func depth(root *TreeNode,curr int) {
	//节点为空,直接返会上一层的高度
	if root == nil {
		return
	}
	curr++
	//当前节点是叶子节点,返回当前节点高度
	if root.Left == nil && root.Right == nil {
		if curr < resMinDepth {
			resMinDepth = curr
		}
		return
	}
	if root.Left != nil {
		depth(root.Left,curr)
	}
	if root.Right != nil {
		depth(root.Right,curr)
	}
}