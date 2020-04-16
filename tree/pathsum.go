package tree

/*
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。
从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
*/
var resPathSum [][]int
func pathSum(root *TreeNode, sum int) [][]int {
	resPathSum = make([][]int,0)
	var res = make([]int,0)
	haspathsum(root,sum,res)
	return resPathSum
}

func haspathsum(root *TreeNode, sum int,res []int) {
	if root == nil {
		return
	}
	res = append(res,root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		resPathSum = append(resPathSum,res)
		return
	}
	var lres = make([]int,len(res))
	copy(lres,res)
	haspathsum(root.Left,sum - root.Val,lres)
	haspathsum(root.Right,sum - root.Val,lres)
}


func pathSum2(root *TreeNode, sum int) [][]int {
	resPathSum = make([][]int,0)
	if root == nil {
		return resPathSum
	}
	var res = make([]int,0)
	backtrack(root,sum,res)
	return resPathSum
}
func backtrack(root *TreeNode, sum int,res []int) {
	res = append(res,root.Val)
	//结束回溯
	if root.Left == nil && root.Right == nil && root.Val == sum {
		var rtn = make([]int,len(res))
		copy(rtn,res)
		resPathSum = append(resPathSum,rtn)
		res = res[:len(res)-1]
		return
	}
	//选择左子树
	if root.Left != nil {
		var lres = make([]int,len(res))
		copy(lres,res)
		backtrack(root.Left,sum - root.Val,lres)
	}
	//选择右子树
	if root.Right != nil {
		var rres = make([]int,len(res))
		copy(rres,res)
		backtrack(root.Right,sum - root.Val,rres)
	}
	res = res[:len(res)-1]
}

