package tree

/*
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

递归解法:
题目分析：
前序遍历特点： 节点按照 [ 根节点 | 左子树 | 右子树 ] 排序，
以题目示例为例：[ 3 | 9 | 20 15 7 ]

中序遍历特点： 节点按照 [ 左子树 | 根节点 | 右子树 ] 排序，
以题目示例为例：[ 9 | 3 | 15 20 7 ]
根据题目描述输入的前序遍历和中序遍历的结果中都不含重复的数字，其表明树中每个节点值都是唯一的。

根据以上特点，可以按顺序完成以下工作：

前序遍历的首个元素即为根节点 root 的值；
在中序遍历中搜索根节点 root 的索引 ，可将中序遍历划分为 [ 左子树 | 根节点 | 右子树 ] 。
根据中序遍历中的左（右）子树的节点数量，可将前序遍历划分为 [ 根节点 | 左子树 | 右子树 ] 。
自此可确定 三个节点的关系 ：1.树的根节点、2.左子树根节点、3.右子树根节点（即前序遍历中左（右）子树的首个元素）。

子树特点： 子树的前序和中序遍历仍符合以上特点，以题目示例的右子树为例：前序遍历：[20 | 15 | 7]，中序遍历 [ 15 | 20 | 7 ] 。

根据子树特点，我们可以通过同样的方法对左（右）子树进行划分，每轮可确认三个节点的关系 。
此递推性质让我们联想到用 递归方法 处理。
递归解析：
递推参数： 前序遍历中根节点的索引pre_root、中序遍历左边界in_left、中序遍历右边界in_right。
终止条件： 当 in_left > in_right ，子树中序遍历为空，说明已经越过叶子节点，此时返回 nullnull 。
递推工作：
建立根节点root： 值为前序遍历中索引为pre_root的节点值。
搜索根节点root在中序遍历的索引i： 为了提升搜索效率，本题解使用哈希表 dic 预存储中序遍历的值与索引的映射关系，
每次搜索的时间复杂度为 O(1)。
构建根节点root的左子树和右子树： 通过调用 recur() 方法开启下一层递归。
左子树： 根节点索引为 pre_root + 1 ，中序遍历的左右边界分别为 in_left 和 i - 1。
右子树： 根节点索引为 i - in_left + pre_root + 1（即：根节点索引 + 左子树长度 + 1），
中序遍历的左右边界分别为 i + 1 和 in_right。

*/
var p []int
var dict map[int]int
func buildTree(preorder []int, inorder []int) *TreeNode {
	var l = len(inorder)
	if l == 0 {
		return nil
	}
	p = preorder
	dict = make(map[int]int)
	for i:=0;i<l;i++{
		dict[inorder[i]] = i
	}
	return recur(0,0,l-1)
}

//参数分别为前序遍历根节点索引,中序遍历左边界与中序遍历右边界
func recur(pre_root,in_left,in_right int) *TreeNode{
	//中序遍历的左边界大于右边界,说明子树已经为空了,不用递归
	if in_left > in_right {
		return nil
	}
	root := &TreeNode{Val:  p[pre_root] ,}
	//根节点在中序遍历中的索引
	i := dict[p[pre_root]]
	//树的根节点索引为 pre_root ，则左子树的根节点为 pre_root + 1 （这是先序遍历把树分成 [ 根节点 | 左子树 | 右子树 ] 的特性）
	//左子树的左边界还是in_left,左子树的右边界就是中序遍历的根节点-1
	root.Left = recur(pre_root + 1,in_left,i - 1)
	//右子树的前序遍历根节点就是 pre_root + 1 + 左子树长度，
	//左子树长度是中序遍历根节点索引i - 左子树的中序遍历左边界in_left ，因此为 i - in_left + pre_root + 1
	//右子树的中序遍历左边界就是根节点索引+1,右子树的右边界就还是in_right
	root.Right = recur(i - in_left + pre_root + 1,i + 1,in_right)
	return root
}




var post []int
var dictPost map[int]int
func buildTreePost(inorder []int, postorder []int) *TreeNode {
	post = postorder
	dictPost = make(map[int]int)
	for k,v := range inorder {
		dictPost[v] = k
	}
	return build(len(postorder)-1,0,len(inorder)-1)
}


func build(postorder_root int, inorder_left,inorder_right int) *TreeNode {
	if inorder_left > inorder_right {
		return nil
	}
	var root = &TreeNode{Val:post[postorder_root]}
	//根节点在中序遍历中的索引
	var i = dictPost[post[postorder_root]]
	//左子树的后续遍历根节点是根节点索引postorder_root-右子树长度-1
	//右子树长度就是中序遍历右边界-中序遍历根索引inorder_right-i
	root.Left = build(postorder_root-1-inorder_right+i,inorder_left,i-1)
	root.Right = build(postorder_root-1,i+1,inorder_right)
	return root
}
