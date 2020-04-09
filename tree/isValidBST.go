package tree

/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
隐含特性,树的左子树的节点全部小于右子树
*/
func isValidBST(root *TreeNode) bool {
	return valid(root,nil,nil)
}

//把根节点的逻辑处理正确,剩下的交给递归框架
//二叉搜索树的
func valid(root,min,max *TreeNode ) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val{
		return false
	}
	if max != nil && root.Val >= max.Val{
		return false
	}
	return valid(root.Left,min,root) && valid(root.Right,root,max)
}
