package tree

/*
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。
*/
var max int
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max = 0
	dfsDepth(root,0)
	return max
}

func dfsDepth(root *TreeNode,depth int) {
	if root == nil {
		return
	}
	depth++
	if depth > max {
		max = depth
	}
	if root.Left != nil {
		dfsDepth(root.Left,depth)
	}
	if root.Right != nil {
		dfsDepth(root.Right,depth)
	}
}
