package tree


/*
题目:判断一个二叉树是否是对称二叉树

解题思路:
如果一个树的左子树与右子树镜像对称，那么这个树是对称的。
因此，该问题可以转化为：两个树在什么情况下互为镜像？
如果同时满足下面的条件，两个树互为镜像：

它们的两个根结点具有相同的值。
每个树的右子树都与另一个树的左子树镜像对称。

*/
func isSymmetric(root *TreeNode) bool {
	return isMirror(root,root)
}

//递归判断树是否对称
func isMirror(root1,root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val &&
		isMirror(root1.Left,root2.Right) &&
		isMirror(root1.Right,root2.Left)
}

//迭代判断树是否对称
/*
除了递归的方法外，我们也可以利用队列进行迭代。队列中每两个连续的结点应该是相等的，
而且它们的子树互为镜像。最初，队列中包含的是 root 以及 root。
该算法的工作原理类似于 BFS，但存在一些关键差异。每次提取两个结点并比较它们的值。
然后，将两个结点的左右子结点按相反的顺序插入队列中。当队列为空时，
或者我们检测到树不对称（即从队列中取出两个不相等的连续结点）时，该算法结束。
*/
func isMirror2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var quene = make([]*TreeNode,0)
	quene = append(quene,root)
	quene = append(quene,root)
	for len(quene) > 0 {
		root1 := quene[0]
		root2 := quene[1]
		quene = quene[2:]
		if root1 == nil && root2 == nil {
			continue
		}
		if root1 == nil || root2 == nil {
			return false
		}
		if root1.Val != root2.Val {
			return false
		}
		quene = append(quene,root1.Left)
		quene = append(quene,root2.Right)

		quene = append(quene,root1.Right)
		quene = append(quene,root2.Left)
	}
	return true
}
