package offer

type TreeNode struct {
	Val int
	Left  *TreeNode
	Right *TreeNode
}

/*
重建二叉树

输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

思路:
前序遍历的首个元素即为根节点 root 的值；
在中序遍历中搜索根节点 root 的索引 ，可将中序遍历划分为 [ 左子树 | 根节点 | 右子树 ] 。
根据中序遍历中的左（右）子树的节点数量，可将前序遍历划分为 [ 根节点 | 左子树 | 右子树 ]

*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	var root = preorder[0]
	var l,leftNum,rightNum = len(inorder),0,0
	for i:=0;i<l;i++{
		if inorder[i] == root {
			leftNum = i + 1
			rightNum = l - i - 1
			break
		}
	}

}