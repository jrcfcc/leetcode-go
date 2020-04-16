package tree

/*
给定一个二叉树，原地将它展开为链表。

例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6

利用后续遍历,反向处理
*/
func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	if root.Left != nil {
		var left = root.Left
		//找到左子树的右子节点
		for left.Right != nil {
			left = left.Right
		}
		var right = root.Right
		//将root的右节点挂到左子树的右子节点上
		left.Right = right
		//然后将root的左子树移到右子树
		root.Right = root.Left
		//root的左子树置为空
		root.Left = nil
	}
}
