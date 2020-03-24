package tree

/*
返回一个二叉树的镜像

1.框架，二叉树肯定有左右递归，结合第二点后再思考怎样递归，
也就是考虑使用前中后序哪个遍历方法
2.最重要的：考虑当我们遍历到一个节点时希望该节点做什么，
因为考虑当前节点也就是考虑所有节点（共性），其他节点的的作用是一样的。
3.考虑终止条件，再根据第二点看看该返回什么
*/
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	l := mirrorTree(root.Left) //获取左子树
	r := mirrorTree(root.Right) //获取右子树
	root.Left = r
	root.Right = l
	return root //返回调换顺序后的节点
}
