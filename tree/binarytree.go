package tree

type TreeNode struct {
  Val int
  Left  *TreeNode
  Right *TreeNode
}


/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
解题思路:递归查找
因为lowestCommonAncestor(root, p, q)的功能是找出以root为根节点的两个节点p和q的最近公共祖先，
所以递归体分三种情况讨论：
如果p和q分别是root的左右节点，那么root就是我们要找的最近公共祖先
如果p和q都是root的左节点，那么返回lowestCommonAncestor(root.left,p,q)
如果p和q都是root的右节点，那么返回lowestCommonAncestor(root.right,p,q)
边界条件讨论：

如果root是null，则说明我们已经找到最底了，返回null表示没找到
如果root与p相等或者与q相等，则返回root
如果左子树没找到，递归函数返回null，证明p和q同在root的右侧，那么最终的公共祖先就是右子树找到的结点
如果右子树没找到，递归函数返回null，证明p和q同在root的左侧，那么最终的公共祖先就是左子树找到的结点

*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	//p,q都在左边
	var leftNode = lowestCommonAncestor(root.Left,p,q)
	//p,q都在右边
	var rightNode = lowestCommonAncestor(root.Right,p,q)
	//左子树没找到,说明p,q都在右子树
	if leftNode == nil {
		return rightNode
	}
	//右子树没找到,说明p,q都在左子树
	if rightNode == nil {
		return leftNode
	}
	//p,q分别在左右两边,那么最近的公共祖先就是root
	return root
}
